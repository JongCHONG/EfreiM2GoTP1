package main

import (
	"Annuaire/annuaire"
	"Annuaire/export"
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
	flag.StringVar(&action, "export", "", "Export the annuaire to a file (e.g., export.json)")
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
	// ...existing code...
	case "export":
		var contacts []export.Person
		for person := range annuaire.Annuaire {
			contacts = append(contacts, export.Person{
				Name:    person.Name,
				Surname: person.Surname,
				Tel:     person.Tel,
			})
		}
		export.ExportAnnuaireJSON("export.json", contacts)
		// ...existing code...
	default:
		fmt.Printf("Unknown action: %s\n", action)
		return
	}

}
