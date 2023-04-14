package reader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/kilianp07/CassandraCRUD/utils/structs"
)

func read() ([]structs.Restaurant, error) {
	// Read JSON file into byte slice
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return
	}

	// Unmarshal JSON data into slice of MorrisParkBakeShop structs
	var bakeries []structs.Restaurant
	err = json.Unmarshal(data, &bakeries)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}
}
