package main

import (
	"bufio"
	"fmt"
	"os"
)

func deleteContact(addressBook map[string]*Contact) {
	fmt.Print("Enter the name of the contact to delete: ")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	delete(addressBook, name)

	err = writeAddressBookToFile(addressBook)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("Contact deleted successfully")
	}
}
