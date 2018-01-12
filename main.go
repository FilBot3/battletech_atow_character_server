package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func checkerr(err error, message string) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(message)
	}
}

func main() {
	// Initalize an instance of the Character Sheet
	var character_sheet CharacterSheet
	// Slurp the JSON File into memory.
	json_file, err := ioutil.ReadFile("example_character_sheets/AToW_RecordSheet.json")
	// Check if we received an error, and print message.
	checkerr(err, "Could not load file.")
	// Parse or Unmarshal the slurped file into the struct instance
	err = json.Unmarshal(json_file, &character_sheet)
	// Check for errors again.
	checkerr(err, "Could not unmarshal JSON byte data.")
	// Print our result. It'll look like junk because its mapped to the
	// structure so there won't be any visible names.
	fmt.Println(character_sheet)
}
