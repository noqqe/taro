package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	scribble "github.com/nanobox-io/golang-scribble"
)

// a fish
type Photo struct {
	URL       string
	Name      string
	Alt       string
	Tags      []string
	Groups    []string
	Published bool
}

func readStringInput(text string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", text)
	s, _ := reader.ReadString('\n')
	s = strings.Replace(s, "\n", "", 1)
	return s
}

func Add() {

	dir := "./db"

	db, err := scribble.New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	var photo Photo

	fmt.Println(photo.Name)
	photo.Name = readStringInput("Name")
	photo.URL = readStringInput("URL")

	db.Write("photos", photo.Name, photo)
}

func List() {

	// Read all fish from the database, unmarshaling the response.
	dir := "./db"

	db, err := scribble.New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	records, err := db.ReadAll("photos")
	for _, f := range records {
		p := Photo{}
		if err := json.Unmarshal([]byte(f), &p); err != nil {
			fmt.Println("Error", err)
		}
		fmt.Println(p)
	}

	// fishies := []Fish{}
	// for _, f := range records {
	// 	fishFound := Fish{}
	// 	if err := json.Unmarshal([]byte(f), &fishFound); err != nil {
	// 		fmt.Println("Error", err)
	// 	}
	// 	fishies = append(fishies, fishFound)
	// }

	// // Delete a fish from the database
	// if err := db.Delete("fish", "onefish"); err != nil {
	// 	fmt.Println("Error", err)
	// }
	//
	// // Delete all fish from the database
	// if err := db.Delete("fish", ""); err != nil {
	// 	fmt.Println("Error", err)
	// }

}
