package main

import (
	"Annuaire/annuaire"
	"flag"
	"fmt"
)

func main() {
	var action string
	var name, surname, tel string

	flag.StringVar(&action, "action", "list", "Lister les personnes dans l'annuaire")
	flag.StringVar(&name, "name", "", "Name of the person to add")
	flag.StringVar(&surname, "surname", "", "Surname of the person to add")
	flag.StringVar(&tel, "tel", "", "Tel number of the person to add")
	flag.Parse()

	switch action {
	case "list":
		fmt.Println("Annuaire actuel :")
		annuaire.PrintAnnuaire()
	case "add":
		annuaire.AddPerson(name, surname, tel)
	case "remove":
		annuaire.RemovePerson(name, surname)
	case "search":
		annuaire.SearchPerson(name)
	case "update":
		annuaire.UpdatePerson(name, surname, tel)
	default:
		fmt.Printf("Unknown action: %s\n", action)
		return
	}

}
