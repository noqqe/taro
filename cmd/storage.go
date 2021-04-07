package cmd

import (
	"bufio"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"

	scribble "github.com/nanobox-io/golang-scribble"
)

// a fish
type Photo struct {
	Id        string
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

// Function for reading comma separated list into db
func readArrayInput(text string) []string {
	return strings.Split(readStringInput(text), ",")
}

func calcHash(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

// Add new photo data to the db store
func Add(filename string) string {

	var photo Photo

	db := getDB()

	photo.Id = calcHash(filename)
	photo.Name = readStringInput("Name")
	photo.Tags = readArrayInput("Tags")
	photo.Filename = path.Base(filename)

	db.Write("photos", photo.Name, photo)

	return photo.Id
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
