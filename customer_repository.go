package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/hspazio/customers_in_range/geopoint"
)

type customer struct {
	UserID    int    `json:"user_id"`
	Name      string `json:"name"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func (c *customer) Location() (g geopoint.Degrees) {
	lat, err := strconv.ParseFloat(c.Latitude, 64)
	lon, err := strconv.ParseFloat(c.Longitude, 64)
	if err != nil {
		return
	}
	return geopoint.Degrees{lat, lon}
}

type customerRepository struct {
	file string
}

func (r *customerRepository) allCustomers() ([]customer, error) {
	b, err := ioutil.ReadFile(r.file)
	if err != nil {
		return nil, err
	}
	customers := []customer{}

	lines := bytes.Split(b, []byte("\n"))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		var c customer
		if err := json.Unmarshal(line, &c); err != nil {
			return nil, err
		}

		customers = append(customers, c)
	}
	return customers, nil
}

func (r *customerRepository) withinRangeInKm(rangeInKm float64, p geopoint.Point) ([]customer, error) {
	var results []customer
	customers, err := r.allCustomers()
	if err != nil {
		return nil, err
	}
	for _, customer := range customers {
		if geopoint.DistanceInKm(p, customer.Location()) <= rangeInKm {
			results = append(results, customer)
		}
	}
	return results, nil
}
