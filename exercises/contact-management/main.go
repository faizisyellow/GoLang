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
	var contacts []Contact
	var name, email string
	var phoneNumber, choose int

	fmt.Println("====Contact Management====")
	fmt.Println("1. Add new contact")
	fmt.Println("2. List all contacts")
	fmt.Println("3. Search contact")
	fmt.Println("4. Quit")

	fmt.Print("Choose an option: ")
	fmt.Scan(&choose)

	switch choose {
	case 1:
		fmt.Print("Name: ")
		fmt.Scan(&name)

		fmt.Print("Phone number: ")
		fmt.Scan(&phoneNumber)

		fmt.Print("Email: ")
		fmt.Scan(&email)

		// Add new contact and store to a slice
		contact := New(name, email, phoneNumber)
		contacts = append(contacts, contact)

		// Store to JSON file
		err := SaveToFile("contacts.json", contacts)
		if err != nil {
			fmt.Println("Error saving contacts:", err)
		}

	case 2:
		// Read contacts from JSON file
		savedContacts, err := ReadDataFromFile("contacts.json")
		if err != nil {
			fmt.Println("Error reading contacts:", err)
			return
		}

		fmt.Println("Saved Contacts:")
		for _, contact := range savedContacts {
			fmt.Printf("Name: %s, Phone: %d, Email: %s\n",
				contact.Name, contact.PhoneNumber, contact.Email)
		}
	case 3:
		// search contact by name
		contacts, err := ReadDataFromFile("contacts.json")
		if err != nil {
			fmt.Println("Can not read file ", err)
			return
		}

		searchContact := SearchByName(contacts)
		for i, value := range searchContact {
			fmt.Printf("Found Contact %v: \nName: %v\nContact: %v\nEmail: %v\n", i+1, value.Name, value.PhoneNumber, value.Email)
		}
	case 4:
		os.Exit(1)
	}
}

func SaveToFile(filename string, data []Contact) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, jsonData, 0644)
}

func ReadDataFromFile(filename string) ([]Contact, error) {
	var contacts []Contact

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &contacts)
	if err != nil {
		return nil, err
	}

	return contacts, nil
}

func SearchByName(data []Contact) []Contact {
	var name string
	var foundContacts []Contact

	fmt.Print("Name: ")
	fmt.Scan(&name)

	for _, contact := range data {
		if contact.Name == name {
			foundContacts = append(foundContacts, contact)
		}
	}

	if len(foundContacts) == 0 {
		fmt.Printf("Contact not found")
	}
	return foundContacts
}
