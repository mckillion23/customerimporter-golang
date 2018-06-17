package main

import (
	"fmt"
	"sort"

	"../customerimporter"
)

func main() {
	domainMap := customerimporter.GetUniqueDomainNamesWithCount("../customers.csv")

	sortedCounts := make([]string, 0, len(domainMap))
	for k := range domainMap {
		sortedCounts = append(sortedCounts, k)
	}
	sort.Strings(sortedCounts)

	for _, k := range sortedCounts {
		fmt.Printf("Domain: %s, Customer Count: %d\n", k, domainMap[k])
	}
}
