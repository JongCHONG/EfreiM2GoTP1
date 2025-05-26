package annuaire

import "fmt"

type Person struct {
	name    string
	surname string
	tel     string
}

var Annuaire = map[Person]string{
	{name: "John", surname: "Doe", tel: "123-456-7890"}: "",
}

func PrintAnnuaire() {
	for person := range Annuaire {
		fmt.Printf("Name: %s Surname: %s, Tel: %s\n", person.name, person.surname, person.tel)
	}
}

func AddPerson(name, surname, tel string) {
	person := Person{name: name, surname: surname, tel: tel}
	if _, exists := Annuaire[person]; exists {
		fmt.Printf("Person %s %s already exists in the annuaire.\n", name, surname)
		return
	}
	Annuaire[person] = ""
	fmt.Printf("Added person: %s %s, Tel: %s\n", name, surname, tel)
}

func RemovePerson(name, surname string) {
	found := false
	for person := range Annuaire {
		if person.name == name && person.surname == surname {
			delete(Annuaire, person)
			fmt.Printf("Removed person: %s %s\n", name, surname)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Person %s %s not found in the annuaire.\n", name, surname)
	}
}

func SearchPerson(name string) {
	found := false
	for person := range Annuaire {
		if person.name == name {
			fmt.Printf("Found person: %s %s, Tel: %s\n", person.name, person.surname, person.tel)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Person with name %s not found in the annuaire.\n", name)
	}
}

func UpdatePerson(name, surname, tel string) {
	found := false
	for person := range Annuaire {
		if person.name == name && person.surname == surname {
			delete(Annuaire, person)
			newPerson := Person{name: name, surname: surname, tel: tel}
			Annuaire[newPerson] = ""
			fmt.Printf("Updated person: %s %s, New Tel: %s\n", name, surname, tel)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Person %s %s not found in the annuaire.\n", name, surname)
	}
}
