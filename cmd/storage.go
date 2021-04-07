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
	Name      string
	Filename  string
	Alt       string
	Tags      []string
	Groups    []string
	Published bool
}

// Generate DB Connection
func getDB() *scribble.Driver {

	dir := "./db"

	db, err := scribble.New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	return db
}

// Function for reading meta data from user to db
func readStringInput(text string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s: ", text)
	s, _ := reader.ReadString('\n')
	s = strings.Replace(s, "\n", "", 1)
	return s
}

// Add new photo data to the db store
func Add(filename string) string {

	var photo Photo

	db := getDB()

	fmt.Println(photo.Name)
	photo.Name = readStringInput("Name")
	photo.Filename = filename

	db.Write("photos", photo.Name, photo)

	return photo.Name
}

// Read all fish from the database, unmarshaling the response.
func List() {

	db := getDB()

	records, _ := db.ReadAll("photos")
	for _, f := range records {
		p := Photo{}
		if err := json.Unmarshal([]byte(f), &p); err != nil {
			fmt.Println("Error", err)
		}
		fmt.Println(p)
	}

}

func Show(name string) {
	fmt.Println(getPhoto(name))
}

// Read all fish from the database, unmarshaling the response.
func getPhoto(name string) Photo {
	db := getDB()

	p := Photo{}
	if err := db.Read("photos", name, &p); err != nil {
		fmt.Println("Error", err)
	}

	return p
}
