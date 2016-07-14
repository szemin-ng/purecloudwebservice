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

// AccountResponse contains account information sent back to PureCloud Web Services Data Dip Connector.
// It follows the format in https://developer.mypurecloud.com/api/webservice-datadip/service-contracts.html
type AccountResponse struct {
	Account Account `json:"Account"`
}

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

// AccountByAccountNumberRequest is the request sent from PureCloud Web Services Data Dip Connector to this app to retrieve
// account information using an account number to query
// It follows the format in https://developer.mypurecloud.com/api/webservice-datadip/service-contracts.html
type AccountByAccountNumberRequest struct {
	AccountNumber   string `json:"AccountNumber"`
	CustomAttribute string `json:"CustomAttribute,omitempty"`
}

func main() {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s\n", port)

	// Setup HTTP server
	var r *mux.Router
	r = mux.NewRouter()
	r.HandleFunc("/GetAccountByAccountNumber", getAccountByAccountNumber).Methods("POST")
	r.HandleFunc("/GetAccountByContactId", getAccountByContactID).Methods("POST")
	r.HandleFunc("/GetAccountByPhoneNumber", getAccountByPhoneNumber).Methods("POST")
	r.HandleFunc("/GetContactByPhoneNumber", getContactByPhoneNumber).Methods("POST")
	r.HandleFunc("/GetMostRecentOpenCaseByContactId", getMostRecentOpenCaseByContactID).Methods("POST")

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

// getAccountByAccountNumber handles HTTP POSTs to /GetAccountByAccountNumber. It reads the request sent and returns
// a response
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

	/* Just an example, hardcode response here.  Sending the following response
		{
		    "Account": {
	    	    "Id": "123",
	        	"Name": "Ng Sze Min",
		        "Number": "123",
	    	    "EmailAddresses": {
	        	    "EmailAddress": [
	            	    {
	                	    "EmailAddress": "szemin.ng@inin.com",
	                    	"EmailType": 1
		                }
	    	        ]
	        	},
		        "PhoneNumbers": {
	    	        "PhoneNumber": [
	        	        {
	            	        "Number": "+60327763333",
	                	    "PhoneType": 1
		                },
	    	            {
	        	            "Number": "+18002671364",
	            	        "PhoneType": 2
	                	}
		            ]
	    	    },
	        	"Addresses": {
	            	"Address": [
	                	{
	                    	"City": "Kuala Lumpur",
		                    "Country": "Malaysia",
	    	                "Line1": "Unit 9.1, Level 9, Menara Prestige",
	        	            "Line2": "No. 1, Jalan Pinang",
	            	        "PostalCode": "50450",
	                	    "State": "FT",
	                    	"Type": "MY"
		                },
	    	            {
	        	            "City": "Indianapolis",
	            	        "Country": "United States",
	                	    "Line1": "7601 Interactive Way",
	                    	"PostalCode": "46278",
		                    "State": "IN",
	    	                "Type": "US"
	        	        }
	            	]
		        },
	    	    "CustomAttribute": "Custom data here"
		    }
		}*/
	var resp = AccountResponse{
		Account: Account{
			ID:     "123",
			Name:   "Ng Sze Min",
			Number: "123",
			Addresses: &Addresses{
				Address: []Address{
					Address{City: "Kuala Lumpur", Country: "Malaysia", Line1: "Unit 9.1, Level 9, Menara Prestige", Line2: "No. 1, Jalan Pinang", PostalCode: "50450", State: "FT", Type: "MY"},
					Address{City: "Indianapolis", Country: "United States", Line1: "7601 Interactive Way", PostalCode: "46278", State: "IN", Type: "US"},
				},
			},
			PhoneNumbers: &PhoneNumbers{
				PhoneNumbers: []PhoneNumber{
					PhoneNumber{Number: "+60327763333", PhoneType: 1},
					PhoneNumber{Number: "+18002671364", PhoneType: 2},
				},
			},
			EmailAddresses: &EmailAddresses{
				EmailAddress: []EmailAddress{
					EmailAddress{EmailAddress: "szemin.ng@inin.com", EmailType: 1},
				},
			},
			CustomAttribute: "Custom data here",
		},
	}

	log.Println("Sending reply from /GetAccountByAccountNumber...")

	// Write reply
	var b []byte
	if b, err = json.Marshal(resp); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(b); err != nil {
		log.Printf("Failed to write: %s\n", err)
	}
}

func getAccountByContactID(w http.ResponseWriter, r *http.Request) {
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

func getMostRecentOpenCaseByContactID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintln(w, http.StatusNotImplemented)
}
