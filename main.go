package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	addressBook := make(map[string]*Contact)

	err := writeAddressBookToFile(addressBook)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("Address book saved to file successfully")
	}

	for {
		fmt.Println("Welcome to the Address Book")
		fmt.Println("1. Add a new contact")
		fmt.Println("2. Edit an existing contact")
		fmt.Println("3. Delete a contact")
		fmt.Println("4. View all contacts")
		fmt.Println("6. Search contacts by city or state ")
		fmt.Println("7. Search contacts by city or state through map")
		fmt.Println("8. count contacts by city or state ")
		fmt.Println("9. Quit")
		fmt.Print("Enter your choice: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			addContact(addressBook)
		case "2":
			editContact(addressBook)
		case "3":
			deleteContact(addressBook)
		case "4":
			viewContacts(addressBook)

		case "6":
			fmt.Print("Enter city or state to search: ")
			cityOrState := ""
			fmt.Scanln(&cityOrState)
			contacts := searchByCityOrState(addressBook, cityOrState, cityOrState)
			if len(contacts) == 0 {
				fmt.Println("No contacts found")
			} else {
				fmt.Println("Matching contacts:")
				for _, contact := range contacts {
					fmt.Println(contact.FirstName, contact.LastName)
				}
			}

		case "7":
			fmt.Print("Enter city or state to search: ")
			cityOrState := ""
			fmt.Scanln(&cityOrState)
			contactsByCityOrState := make(map[string][]*Contact)

			for _, contact := range addressBook {
				if _, ok := contactsByCityOrState[contact.City]; !ok {
					contactsByCityOrState[contact.City] = searchByCityOrState(addressBook, contact.City, "")
				}
				if _, ok := contactsByCityOrState[contact.State]; !ok {
					contactsByCityOrState[contact.State] = searchByCityOrState(addressBook, "", contact.State)
				}
			}
		case "8":
			fmt.Print("Enter city or state to search: ")
			cityOrState := ""
			fmt.Scanln(&cityOrState)
			contactsByCityOrState := filterContact(addressBook, func(c *Contact) bool {
				return c.City == cityOrState || c.State == cityOrState
			})
			if len(contactsByCityOrState) == 0 {
				fmt.Println("No contacts found")
			} else {
				fmt.Println("Matching contacts by city or state:")
				for cityOrState, count := range contactsByCityOrState {
					fmt.Printf("%s: %d contacts\n", cityOrState, count)
				}
			}
		case "9":
			fmt.Println("Thank you for using the Address Book")
			return
		default:
			fmt.Println("Invalid input, please try again")
		}

	}
}
