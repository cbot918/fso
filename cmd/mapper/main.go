package main

import (
	"encoding/json"
	"log"

	"github.com/cbot918/fss/internal/model"
	"github.com/cbot918/fss/internal/utils"
	u "github.com/cbot918/fss/utils"
)

func main() {

	byteValue, err := utils.ReadMap("http.json")
	if err != nil {
		log.Fatal(err)
	}

	var result *model.Planet
	json.Unmarshal([]byte(byteValue), &result)

	u.Lj(result)
	// fmt.Printf("%#v\n", result)
}
