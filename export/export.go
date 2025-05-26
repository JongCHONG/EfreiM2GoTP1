package export

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Tel     string `json:"tel"`
}

func ExportAnnuaireJSON(filename string, contacts []Person) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(contacts); err != nil {
		return err
	}
	fmt.Printf("Annuaire exporté dans %s\n", filename)
	return nil
}
