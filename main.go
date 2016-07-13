package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
)

type Account struct {
	ID              string          `json:"Id,omitempty"`
	Name            string          `json:"Name,omitempty"`
	Number          string          `json:"Number,omitempty"`
	EmailAddresses  *EmailAddresses `json:"EmailAddresses,omitempty"`
	PhoneNumbers    *PhoneNumbers   `json:"PhoneNumbers,omitempty"`
	Addresses       *Addresses      `json:"Addresses,omitempty"`
	CustomAttribute string          `json:"CustomAttribute,omitempty"`
}

type Addresses struct {
	Address []Address `json:"Address,omitempty"`
}

type Address struct {
	City       string `json:"City,omitempty"`
	Country    string `json:"Country,omitempty"`
	Line1      string `json:"Line1,omitempty"`
	Line2      string `json:"Line2,omitempty"`
	Line3      string `json:"Line3,omitempty"`
	PostalCode string `json:"PostalCode,omitempty"`
	State      string `json:"State,omitempty"`
	Type       string `json:"Type,omitempty"`
}

type EmailAddresses struct {
	EmailAddress []EmailAddress `json:"EmailAddress,omitempty"`
}

type EmailAddress struct {
	EmailAddress string  `json:"EmailAddress,omitempty"`
	EmailType    float32 `json:"EmailType,omitempty`
}

type PhoneNumbers struct {
	PhoneNumbers []PhoneNumber `json:"PhoneNumber,omitempty"`
}

type PhoneNumber struct {
	Number    string  `json:"Number,omitempty"`
	PhoneType float32 `json:"PhoneType,omitempty"`
}

type AccountByAccountNumberRequest struct {
	AccountNumber   string `json:"AccountNumber"`
	CustomAttribute string `json:"CustomAttribute,omitempty"`
}

type AccountByContactIDRequest struct {
	ContactID       string `json:"ContactId"`
	CustomAttribute string `json:"CustomAttribute,omitempty"`
}

func main() {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s\n", port)

	var r *mux.Router
	r = mux.NewRouter()
	r.HandleFunc("/GetAccountByAccountNumber", getAccountByAccountNumber).Methods("POST")
	r.HandleFunc("/GetAccountByContactId", getAccountByContactId).Methods("POST")
	r.HandleFunc("/GetAccountByPhoneNumber", getAccountByPhoneNumber).Methods("POST")
	r.HandleFunc("/GetContactByPhoneNumber", getContactByPhoneNumber).Methods("POST")
	r.HandleFunc("/GetMostRecentOpenCaseByContactId", getMostRecentOpenCaseByContactId).Methods("POST")

	// Start HTTP server
	var server *http.Server
	server = &http.Server{Addr: ":" + port, Handler: r}
	log.Println("Starting server...")
	go func() {
		server.ListenAndServe()
	}()

	// Wait for SIGINT or SIGKILL
	var interrupt = make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	<-interrupt
}

func getAccountByAccountNumber(w http.ResponseWriter, r *http.Request) {
	var err error

	log.Println("Processing /GetAccountByAccountNumber...")

	// Retrieve request body
	var req AccountByAccountNumberRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Failed to decode JSON request body: %s\n", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, err)
		return
	}

	/*	var resp = Account{
		ID:     "1234",
		Name:   "Test",
		Number: "abc123",
		Addresses: &Addresses{
			Address: []Address{
				Address{City: "KL", Country: "MY", Line1: "Jalan", Line2: "Jiran", Line3: "9", PostalCode: "56000", State: "FT", Type: "X"},
			},
		},
		PhoneNumbers: &PhoneNumbers{
			PhoneNumbers: []PhoneNumber{
				PhoneNumber{Number: "78901234", PhoneType: 9},
			},
		},
		EmailAddresses: &EmailAddresses{
			EmailAddress: []EmailAddress{
				EmailAddress{EmailAddress: "szemin.ng@inin.com", EmailType: 8},
			},
		},
		CustomAttribute: "Custom",
	}*/

	log.Println("Sending reply from /GetAccountByAccountNumber...")

	var a string = `{"Account.Id": "123",` +
		`	"Account.Name": "asd",` +
		`	"Account.Number": "123123",` +
		`	"Account.Addresses.Address.City": [ "City1", "City2" ],` +
		`	"Account.Addresses.Address.Country": ["Address1", "Address2"],` +
		`	"Account.Addresses.Address.Line1": ["Line1", "Line2"],` +
		`	"Account.Addresses.Address.Line2": [ "Line1","Line2"],` +
		`	"Account.Addresses.Address.Line3": ["Line1", "Line2"],` +
		`	"Account.Addresses.Address.PostalCode": ["Post1","Post2"],` +
		`	"Account.Addresses.Address.State": ["State1","State2"],` +
		`	"Account.Addresses.Address.Type": ["Type1", "Type2"],` +
		`	"Account.PhoneNumbers.PhoneNumber.Number": ["Num1", "Num2"],` +
		`	"Account.PhoneNumbers.PhoneNumber.PhoneType": ["Type1", "Type2"],` +
		`	"Account.EmailAddresses.EmailAddress.EmailAddress": ["Email1","Email2"],` +
		`	"Account.EmailAddresses.EmailAddress.EmailType": ["type1", "type2"],` +
		`	"Account.CustomAttribute": "custom"}`

	// Write reply
	/*	var b []byte
		if b, err = json.Marshal(resp); err != nil {
			panic(err)
		}*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write([]byte(a)); err != nil {
		log.Printf("Failed to write: %s\n", err)
	}
}

func getAccountByContactId(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintln(w, http.StatusNotImplemented)
}

func getAccountByPhoneNumber(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintln(w, http.StatusNotImplemented)
}

func getContactByPhoneNumber(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintln(w, http.StatusNotImplemented)
}

func getMostRecentOpenCaseByContactId(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintln(w, http.StatusNotImplemented)
}
