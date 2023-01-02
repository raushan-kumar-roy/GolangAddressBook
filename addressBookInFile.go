package main

import "os"

func writeAddressBookToFile(addressBook map[string]*Contact) error {
	file, err := os.Create("address_book.txt")
	if err != nil {
		return err
	}
	defer file.Close()

	for name, contact := range addressBook {
		_, err := file.WriteString(name + ":\n")
		if err != nil {
			return err
		}
		_, err = file.WriteString("  First Name: " + contact.FirstName + "\n")
		if err != nil {
			return err
		}
		_, err = file.WriteString("  Last Name: " + contact.LastName + "\n")
		if err != nil {
			return err
		}
		_, err = file.WriteString("  Address: " + contact.Address + "\n")
		if err != nil {
			return err
		}
		_, err = file.WriteString("  City: " + contact.City + "\n")
		if err != nil {
			return err
		}
		_, err = file.WriteString("  State: " + contact.State + "\n")
		if err != nil {
			return err
		}
		_, err = file.WriteString("  Zip: " + contact.Zip + "\n")
		if err != nil {
			return err
		}
		_, err = file.WriteString("  Phone: " + contact.Phone + "\n")
		if err != nil {
			return err
		}
		_, err = file.WriteString("  Email: " + contact.Email + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
