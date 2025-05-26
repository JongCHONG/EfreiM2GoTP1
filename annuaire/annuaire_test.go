package annuaire

import (
	"testing"
)

// Test de l'ajout d'une personne
func TestAddPerson(t *testing.T) {
	// Nettoyage de l'annuaire pour le test
	Annuaire = make(map[Person]string)

	AddPerson("Alice", "Smith", "0601020304")
	if len(Annuaire) != 1 {
		t.Errorf("L'annuaire devrait contenir 1 contact, trouvé %d", len(Annuaire))
	}
	// Vérifie que la personne ajoutée existe bien
	found := false
	for person := range Annuaire {
		if person.Name == "Alice" && person.Surname == "Smith" && person.Tel == "0601020304" {
			found = true
		}
	}
	if !found {
		t.Error("Le contact Alice Smith n'a pas été trouvé après l'ajout")
	}
}

// Test de la recherche d'une personne
func TestSearchPerson(t *testing.T) {
	// Prépare un annuaire avec un contact
	Annuaire = make(map[Person]string)
	AddPerson("Bob", "Martin", "0708091011")

	found := false
	for person := range Annuaire {
		if person.Name == "Bob" && person.Surname == "Martin" {
			found = true
		}
	}
	if !found {
		t.Error("Le contact Bob Martin devrait être trouvé dans l'annuaire")
	}
}
