package main

func searchContactsByCityOrState(contactsByCityOrState map[string][]*Contact, cityOrState string) []*Contact {
	contacts, ok := contactsByCityOrState[cityOrState]
	if !ok {
		return nil
	}
	return contacts
}
