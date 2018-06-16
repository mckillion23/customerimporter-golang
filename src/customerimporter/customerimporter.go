package customerimporter

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

// Person - struct to parse each row of the CSV
type Person struct {
	Firstname string
	Lastname  string
	Email     string
	Gender    string
	IPAddress string
}

// GetPeople - Returns a list of people from the CSV
func GetPeople(fileName string) []Person {
	reader := getFileReader(fileName)
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

// GetEmails - Returns a list of emails from the CSV
func GetEmails(fileName string) []string {
	reader := getFileReader(fileName)
	var string []string
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		string = append(string, line[2])

	}
	return string
}

// GetDomains - Returns a list of domains from the CSV
func GetDomains(fileName string) []string {
	reader := getFileReader(fileName)
	var domains []string
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if line[2] != "" && strings.Contains(line[2], "@") {
			components := strings.Split(line[2], "@")
			if components[1] != "" {
				domains = append(domains, components[1])
			}
		} else {
			log.Println("Invalid email found. Value = " + line[2])
		}
	}
	return domains
}

// GetUniqueSortedDomains - Returns a list of domains from the CSV
func GetUniqueSortedDomains(fileName string) []string {
	reader := getFileReader(fileName)
	var domains []string
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if line[2] != "" && strings.Contains(line[2], "@") {
			components := strings.Split(line[2], "@")
			if components[1] != "" {
				domains = append(domains, components[1])
			}
		} else {
			log.Println("Invalid email found. Value = " + line[2])
		}
	}
	result := removeDuplicatesUnordered(domains)
	sort.Sort(sort.StringSlice(result))
	return result
}

// getFileReader - Returns file reader based off the file name
func getFileReader(fileName string) *csv.Reader {
	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	return reader
}

func removeDuplicatesUnordered(elements []string) []string {
	encountered := map[string]bool{}

	// Create a map of all unique elements.
	for v := range elements {
		encountered[elements[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key, _ := range encountered {
		result = append(result, key)
	}
	return result
}

// DomainCount - Count some
func DomainCount(list []string) map[string]int {

	duplicateFrequency := make(map[string]int)

	for _, item := range list {
		// check if the item/element exist in the duplicate_frequency map

		_, exist := duplicateFrequency[item]

		if exist {
			duplicateFrequency[item]++ // increase counter by 1 if already in the map
		} else {
			duplicateFrequency[item] = 1 // else start counting from 1
		}
	}
	return duplicateFrequency
}
