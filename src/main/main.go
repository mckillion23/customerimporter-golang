package main

import (
	"fmt"
	"sort"

	"../customerimporter"
)

func main() {
	domainNames := customerimporter.GetAllDomainNames("../customers.csv")
	fmt.Println("Total Domain Names = ", len(domainNames))

	domainMap := customerimporter.GetDomainNameWithCount(domainNames)
	fmt.Println("Total Unique Domains = ", len(domainMap))

	sortedCounts := make([]string, 0, len(domainMap))
	for k := range domainMap {
		sortedCounts = append(sortedCounts, k)
	}
	sort.Strings(sortedCounts)

	for _, k := range sortedCounts {
		fmt.Printf("Domain: %s, Customer Count: %d\n", k, domainMap[k])
	}
}
