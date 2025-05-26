package importation

import (
	"Annuaire/annuaire"
	"encoding/json"
	"fmt"
	"os"
)

func ImportAnnuaireJSON(filename string) ([]annuaire.Person, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	var contacts []annuaire.Person
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&contacts); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	for _, contact := range contacts {
		annuaire.AddPerson(contact.Name, contact.Surname, contact.Tel)
	}

	fmt.Printf("Imported %d contacts from %s\n", len(contacts), filename)
	return contacts, nil
}
