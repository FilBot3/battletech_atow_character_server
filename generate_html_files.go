package atow

import (
	"bufio"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

func GenerateHtmlPage(character_name string) (bool, error) {
	// Initalize an instance of the Character Sheet
	var character_sheet CharacterSheet
	// Slurp the JSON File into memory.
	json_file, err := ioutil.ReadFile(character_name)
	// Check if we received an error, and print message.
	if err != nil {
		log.Println("Could not Read specified file")
		return false, err
	}
	// Parse or Unmarshal the slurped file into the struct instance
	err = json.Unmarshal(json_file, &character_sheet)
	// Check for errors again.
	if err != nil {
		log.Println("Could not unmarshal JSON data into byte string.")
		return false, err
	}

	// Create a file to write to.
	character_file, err := os.OpenFile(character_name+".html", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("Could not open file for writing, or could not be created.")
		return false, err
	}
	defer character_file.Close()

	// Make a buffer to put the text in.
	template_mem_buffer := bufio.NewWriter(character_file)

	// Create a new Template.
	template_example := template.New("BT AToW Example")
	// Parse the Template String or File passed in.
	// Loading HtmlTemplate from the html_template.go file.
	template_example, err = template_example.Parse(HtmlTemplate)
	if err != nil {
		log.Println("Could not parse template")
		return false, err
	}
	// Execute the template.
	// This writes the template to the File buffer created earlier.
	template_example.Execute(template_mem_buffer, character_sheet)
	// Write the file to disk.
	template_mem_buffer.Flush()

	return true, nil
}
