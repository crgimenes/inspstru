package main

import (
	"github.com/crgimenes/inspstru"
)

func main() {
	type Address struct {
		City  string
		State string
	}
	type Person struct {
		Name    string
		Age     int
		Address Address
		Tags    map[string]string
		Scores  []int
	}

	p := Person{
		Name: "Alice",
		Age:  30,
		Address: Address{
			City:  "Wonderland",
			State: "Fantasy",
		},
		Tags: map[string]string{
			"Role":   "Adventurer",
			"Status": "Active",
		},
		Scores: []int{100, 98, 95},
	}

	inspstru.PrintElements(p)
}
