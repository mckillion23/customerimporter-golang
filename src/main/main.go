package main

import (
	"fmt"
	"sort"

	"../customerimporter"
)

func main() {
	domainMap := customerimporter.GetUniqueDomainNamesWithCount("../customers.csv")

	sortedNames := make([]string, 0, len(domainMap))
	for k := range domainMap {
		sortedNames = append(sortedNames, k)
	}
	sort.Strings(sortedNames)

	for _, k := range sortedNames {
		fmt.Printf("Domain: %s, Customer Count: %d\n", k, domainMap[k])
	}
}
