package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Contact struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var contacts []Contact

func main() {
	router := mux.NewRouter()

	contacts = append(contacts, Contact{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	contacts = append(contacts, Contact{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	contacts = append(contacts, Contact{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	router.HandleFunc("/contacts", GetContacts).Methods("GET")
	router.HandleFunc("/contacts/{id}", GetContact).Methods("GET")
	router.HandleFunc("/contacts/{id}", CreateContact).Methods("POST")
	router.HandleFunc("/contacts/{id}", DeleteContact).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8002", router))
}

func GetContacts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(contacts)
}

func GetContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, contact := range contacts {
		if contact.ID == params["id"] {
			json.NewEncoder(w).Encode(contact)
			return
		}
	}

	json.NewEncoder(w).Encode(&Contact{})
}

func CreateContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var contact Contact
	_ = json.NewDecoder(r.Body).Decode(&contact)
	contact.ID = params["id"]
	contacts = append(contacts, contact)
	json.NewEncoder(w).Encode(contacts)
}

func DeleteContact(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, contact := range contacts {
		if contact.ID == params["id"] {
			contacts = append(contacts[:index], contacts[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(contacts)
	}
}
