package main

import (
	"fmt"
	"log"

	"github.com/hspazio/customers_in_range/geopoint"
)

func main() {
	reference := geopoint.Degrees{53.3393, -6.2576841}
	rangeInKm := 100.0
	repo := &customerRepository{"customers.txt"}

	report := customerInviteReport{
		reference: reference,
		rangeInKm: rangeInKm,
		repo:      repo,
	}

	lines, err := report.generate()
	if err != nil {
		log.Fatal(err)
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}
