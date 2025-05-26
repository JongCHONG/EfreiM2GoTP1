package main

import (
	"Annuaire/annuaire"
	exportation "Annuaire/exportation"
	importation "Annuaire/importation"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
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
	case "export":
		var contacts []exportation.Person
		for person := range annuaire.Annuaire {
			contacts = append(contacts, exportation.Person{
				Name:    person.Name,
				Surname: person.Surname,
				Tel:     person.Tel,
			})
		}
		exportation.ExportAnnuaireJSON("annuaire.json", contacts)
	case "import":
		importation.ImportAnnuaireJSON("annuaire.json")
	case "serve":
		tmpl := template.Must(template.ParseFiles("index.html"))
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			var persons []annuaire.Person
			for person := range annuaire.Annuaire {
				persons = append(persons, person)
			}
			tmpl.Execute(w, persons)
		})
		log.Println("Serveur démarré sur http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	default:
		fmt.Printf("Unknown action: %s\n", action)
		return
	}

}
