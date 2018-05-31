package main

import (
	"sort"

	"github.com/hspazio/customers_in_range/geopoint"
)

type customerInviteReport struct {
	reference geopoint.Degrees
	rangeInKm float64
	repo      *customerRepository
}

func (r customerInviteReport) generate() ([]reportLine, error) {
	var results []reportLine
	customers, err := r.repo.withinRangeInKm(r.rangeInKm, r.reference)
	if err != nil {
		return nil, err
	}
	for _, c := range customers {
		results = append(results, reportLine{c.UserID, c.Name})
	}
	sort.Sort(byID(results))
	return results, nil
}

type reportLine struct {
	id   int
	name string
}

type byID []reportLine

func (r byID) Len() int           { return len(r) }
func (r byID) Less(i, j int) bool { return r[i].id < r[j].id }
func (r byID) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
