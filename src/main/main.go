package main

import (
	"fmt"
	"sort"

	"../customerimporter"
)

func main() {
	domains := customerimporter.GetDomains("../customers.csv")
	fmt.Println("Total Length= ", len(domains))

	domainMap := domainCount(domains)
	fmt.Println("Unique Domains= ", len(domainMap))

	sortedCounts := make([]string, 0, len(domainMap))
	for k := range domainMap {
		sortedCounts = append(sortedCounts, k)
	}
	sort.Strings(sortedCounts)

	for _, k := range sortedCounts {
		fmt.Printf("Domain: %s, Customer Count: %d\n", k, domainMap[k])
	}
}

func domainCount(domains []string) map[string]int {

	domainFrequency := make(map[string]int)

	for _, item := range domains {
		// check if the item/element exist in the domainFrequency map

		_, exist := domainFrequency[item]

		if exist {
			domainFrequency[item]++ // increase counter by 1 if already in the map
		} else {
			domainFrequency[item] = 1 // else start counting from 1
		}
	}

	return domainFrequency
}
