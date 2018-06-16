package customerimporter

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strings"
)

// GetAllDomainNames - Returns a list of domains from the CSV
func GetAllDomainNames(fileName string) []string {
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

// GetDomainNameWithCount -
func GetDomainNameWithCount(domains []string) map[string]int {
	domainFrequency := make(map[string]int)
	for _, item := range domains {
		_, exist := domainFrequency[item]
		if exist {
			domainFrequency[item]++
		} else {
			domainFrequency[item] = 1
		}
	}
	return domainFrequency
}

func getFileReader(fileName string) *csv.Reader {
	csvFile, _ := os.Open(fileName)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	return reader
}
