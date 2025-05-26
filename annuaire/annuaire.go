package annuaire

import "fmt"

type Person struct {
	Name    string
	Surname string
	Tel     string
}

var Annuaire = map[Person]string{
	{Name: "John", Surname: "Doe", Tel: "123-456-7890"}: "",
}

func PrintAnnuaire() {
	for person := range Annuaire {
		fmt.Printf("Name: %s Surname: %s, Tel: %s\n", person.Name, person.Surname, person.Tel)
	}
}

func AddPerson(Name, Surname, Tel string) {
	person := Person{Name: Name, Surname: Surname, Tel: Tel}
	if _, exists := Annuaire[person]; exists {
		fmt.Printf("Person %s %s already exists in the annuaire.\n", Name, Surname)
		return
	}
	Annuaire[person] = ""
	fmt.Printf("Added person: %s %s, Tel: %s\n", Name, Surname, Tel)
}

func RemovePerson(Name, Surname string) {
	found := false
	for person := range Annuaire {
		if person.Name == Name && person.Surname == Surname {
			delete(Annuaire, person)
			fmt.Printf("Removed person: %s %s\n", Name, Surname)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Person %s %s not found in the annuaire.\n", Name, Surname)
	}
}

func SearchPerson(Name string) {
	found := false
	for person := range Annuaire {
		if person.Name == Name {
			fmt.Printf("Found person: %s %s, Tel: %s\n", person.Name, person.Surname, person.Tel)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Person with Name %s not found in the annuaire.\n", Name)
	}
}

func UpdatePerson(Name, Surname, Tel string) {
	found := false
	for person := range Annuaire {
		if person.Name == Name && person.Surname == Surname {
			delete(Annuaire, person)
			newPerson := Person{Name: Name, Surname: Surname, Tel: Tel}
			Annuaire[newPerson] = ""
			fmt.Printf("Updated person: %s %s, New Tel: %s\n", Name, Surname, Tel)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Person %s %s not found in the annuaire.\n", Name, Surname)
	}
}
