// Package customerimporter reads from the given customers.csv file and returns a
// sorted data structure of your choice of email domains along with the number
// of customers with e-mail addresses for each domain.
// Any errors should be logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

// Person Comment
type Person struct {
	Firstname string
	Lastname  string
	Email     string
	Gender    string
	IPAddress string
}

// GetPeople Comment
func GetPeople(fileName string) []Person {
	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var people []Person
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		people = append(people, Person{
			Firstname: line[0],
			Lastname:  line[1],
			Email:     line[2],
			Gender:    line[3],
			IPAddress: line[4],
		})
	}
	return people
}
