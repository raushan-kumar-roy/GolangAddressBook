package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func editContact(addressBook map[string]*Contact) {
	fmt.Print("Enter the name of the contact to edit: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	contact, ok := addressBook[name]
	if !ok {
		fmt.Println("Contact not found")
		return
	}

	fmt.Println("1. First Name")
	fmt.Println("2. Last Name")
	fmt.Println("3. Address")
	fmt.Println("4. City")
	fmt.Println("5. State")
	fmt.Println("6. Zip Code")
	fmt.Println("7. Phone Number")
	fmt.Println("8. Email")
	fmt.Print("Enter your choice: ")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Print("Enter new first name: ")
		fmt.Scanln(&contact.FirstName)

	case 2:
		fmt.Print("Enter new last name: ")
		fmt.Scanln(&contact.LastName)
	case 3:
		fmt.Print("Enter new address: ")
		fmt.Scanln(&contact.Address)

	case 4:
		fmt.Print("Enter new city: ")
		fmt.Scanln(&contact.City)

	case 5:
		fmt.Print("Enter new state: ")
		fmt.Scanln(&contact.State)

	case 6:
		fmt.Print("Enter new zip code: ")
		fmt.Scanln(&contact.Zip)

	case 7:
		fmt.Print("Enter new phone number: ")
		fmt.Scanln(&contact.Phone)

	case 8:
		fmt.Print("Enter new email: ")
		fmt.Scanln(&contact.Email)

	default:
		fmt.Println("Invalid input, please try again")
		return
	}
	err := writeAddressBookToFile(addressBook)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("Contact edited successfully")
	}
}
