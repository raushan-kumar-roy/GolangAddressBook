package main

import "fmt"

func addContact(addressBook map[string]*Contact) error {
	newContact := &Contact{}

	fmt.Print("Enter first name: ")
	fmt.Scanln(&newContact.FirstName)
	fmt.Print("Enter last name: ")
	fmt.Scanln(&newContact.LastName)

	duplicateContacts := checkForDuplicateName(addressBook, newContact)
	if len(duplicateContacts) > 0 {
		return fmt.Errorf("There is already a contact with the name %s %s in the address book", newContact.FirstName, newContact.LastName)
	}

	fmt.Print("Enter address: ")
	fmt.Scanln(&newContact.Address)
	fmt.Print("Enter city: ")
	fmt.Scanln(&newContact.City)
	fmt.Print("Enter state: ")
	fmt.Scanln(&newContact.State)
	fmt.Print("Enter zip code: ")
	fmt.Scanln(&newContact.Zip)
	fmt.Print("Enter phone number: ")
	fmt.Scanln(&newContact.Phone)
	fmt.Print("Enter email: ")
	fmt.Scanln(&newContact.Email)

	addressBook[newContact.FirstName+" "+newContact.LastName] = newContact

	err := writeAddressBookToFile(addressBook)
	if err != nil {
		return fmt.Errorf("Error writing to file: %v", err)
	}
	fmt.Println("Contact added successfully")
	return nil
}
