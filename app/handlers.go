package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/bankingApp/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct {
	Name    string `json:"name,omitempty" xml:"name"`
	City    string `json:"city,omitempty" xml:"city"`
	ZipCode string `json:"zipCode,omitempty" xml:"zipCode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	//customers := []Customer{
	//	{"Fran", "Cordoba", "1690"},
	//	{"Gabo", "Buenos Aires", "1419"},
	//}

	customers, _ := ch.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)

	} else {

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}

}
