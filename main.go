package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
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
	//fmt.Println(character_sheet)

	// Create a file to write to.
	character_file, err := os.OpenFile("character_sheet.html", os.O_WRONLY|os.O_CREATE, 0666)
	checkerr(err, "Unable to create file.")
	defer character_file.Close()

	// Make a buffer to put the text in.
	template_mem_buffer := bufio.NewWriter(character_file)

	// Create a new Template.
	template_example := template.New("BT AToW Example")
	// Parse the Template String or File passed in.
	template_example, err = template_example.Parse(HtmlTemplate)
	// Execute the template.
	template_example.Execute(template_mem_buffer, character_sheet)
	template_mem_buffer.Flush()
}
