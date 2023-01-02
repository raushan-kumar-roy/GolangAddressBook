package main

import "fmt"

func viewContacts(addressBook map[string]*Contact) {
	fmt.Println("All contacts:")
	for name, contact := range addressBook {
		fmt.Println("Name:", name)
		fmt.Println("  Address:", contact.Address)
		fmt.Println("  City:", contact.City)
		fmt.Println("  State:", contact.State)
		fmt.Println("  Zip:", contact.Zip)
		fmt.Println("  Phone:", contact.Phone)
		fmt.Println("  Email:", contact.Email)
	}
}
