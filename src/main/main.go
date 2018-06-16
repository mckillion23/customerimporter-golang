package main

import (
	"fmt"

	"../customerimporter"
)

func main() {
	people := customerimporter.GetPeople("../customers.csv")
	fmt.Println(people)
}
