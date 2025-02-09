package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Contact struct {
	Name        string `json:"name"`
	PhoneNumber int    `json:"phoneNumber"`
	Email       string `json:"email"`
}

func New(n, e string, p int) Contact {
	return Contact{Name: n, PhoneNumber: p, Email: e}
}

func main() {
	// collect contacts
	var contacts []Contact

	//- input name, phone number, and email
	var name, email string
	var phoneNumber int

	count := 3
	for i := 0; i < count; i++ {
		fmt.Printf("Name:\n")
		fmt.Scan(&name)

		fmt.Printf("Phone number: \n")
		fmt.Scan(&phoneNumber)

		fmt.Printf("Email: \n")
		fmt.Scan(&email)

		// add new contact and store to a slice
		contact := New(name, email, phoneNumber)
		contacts = append(contacts, contact)
	}

	// store to json
	err := SaveToFile("contacts.json", contacts)
	if err != nil {
		fmt.Println(err.Error())
	}

	// TODO: see all contacts from json file
	fmt.Println(contacts)
}

func SaveToFile[T any](filename string, data T) error {
	json, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return os.WriteFile(filename, json, 0644)
}
