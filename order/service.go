package order

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"unicode"

	orderrpc "github.com/bigbluedisco/tech-challenge/backend/v1/order/rpc"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"

	"github.com/bigbluedisco/tech-challenge/backend/v1/store"
)

// Service holds RPC handlers for the order service. It implements the orderrpc.ServiceServer interface.
type service struct {
	orderrpc.UnimplementedServiceServer
	s store.OrderStore
}

func NewService(s store.OrderStore) *service {
	return &service{s: s}
}

// Fetch all existing orders in the system.
func (s *service) ListOrders(ctx context.Context, r *orderrpc.ListOrdersRequest) (*orderrpc.ListOrdersResponse, error) {
	return &orderrpc.ListOrdersResponse{Orders: s.s.Orders()}, nil
}

// Remove diacritics
func removeDiacritics(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, _ := transform.String(t, s)
	return result
}

// Quest to website (return addr, postcode,city )
func quest(path string) (string, string, string, error) {

	//Quest to https://api-adresse.data.gouv.fr/
	res, err := http.Get(path)
	if err != nil {
		return "", "", "", err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
		return "", "", "", err
	}

	var result ModelAdr
	if err := json.Unmarshal([]byte(body), &result); err != nil {
		return "", "", "", err
	}

	// If search returns no result, return error
	if len(result.Features) == 0 || err != nil {
		return "", "", "", errors.New("address not found")
	}

	// Else return the first result in the list
	postcode := result.Features[0].Properties.PostCode
	city := result.Features[0].Properties.City
	adrname := result.Features[0].Properties.Name
	return adrname, postcode, city, nil
}

// Verify Address
func verifyAddr(order *orderrpc.Order) error {

	// 1) Get Country Correct
	country := order.GetAddr().GetCountry()
	if !(strings.EqualFold(country, "france") || strings.EqualFold(country, "fr") || strings.EqualFold(country, "")) {
		return errors.New("send in France only")
	}
	order.Addr.Country = "France"

	// 2) Get Address correct
	city := removeDiacritics(order.Addr.City)
	postcode := order.Addr.PostalCode
	addr := removeDiacritics(order.Addr.Address)

	if city == "" || postcode == "" || addr == "" {
		return errors.New("address not complete")
	}

	query := fmt.Sprintf("https://api-adresse.data.gouv.fr/search/?" +
		"q=" + url.QueryEscape(addr) + "&" +
		"city=" + url.QueryEscape(city) + "||" +
		"postcode=" + url.QueryEscape(postcode))

	fmt.Println(query)

	addrq, codeq, cityq, err := quest(query)
	if err != nil {
		return err
	}

	order.Addr.Address = addrq
	order.Addr.City = cityq
	order.Addr.PostalCode = codeq

	return nil
}

// Create an order
func (s *service) CreateOrder(ctx context.Context, order *orderrpc.Order) (*orderrpc.CreateOrderResponse, error) {

	// 1) Check customer first and last name not empty
	if order.GetC().FirstName == "" || order.GetC().LastName == "" {
		return nil, errors.New("customer name not completed")
	}

	// 2) Verify all products exist in the store
	pdstor := store.NewProductStore()
	for i := 0; i < len(order.ProdQuant); i++ {
		if item, err := pdstor.Product(order.ProdQuant[i].Pid); item == nil || err != nil {
			return nil, errors.New("product (" + order.ProdQuant[i].Pid + ") not found")
		}
	}

	// 3) Verify the Shipping Address
	err := verifyAddr(order)
	if err != nil {
		return nil, err
	}

	// 4) Upsert Order to OrderStore
	s.s.SetOrder(order)

	return &orderrpc.CreateOrderResponse{}, nil
}
