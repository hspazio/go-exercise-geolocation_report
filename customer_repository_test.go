package main

import (
	"fmt"
	"testing"

	"github.com/hspazio/customers_in_range/geopoint"
	"github.com/stretchr/testify/assert"
)

func TestCustomerRepository(t *testing.T) {
	assert := assert.New(t)

	t.Run("lists all customers", func(t *testing.T) {
		repo := &customerRepository{"customers.txt"}

		customers, err := repo.allCustomers()
		assert.NoError(err)
		assert.NotEmpty(customers)
	})

	t.Run("lists all customers within a certain distance from a point", func(t *testing.T) {
		repo := &customerRepository{"customers.txt"}

		rangeInKm := 100.0
		p := geopoint.Degrees{53.3393, -6.2576841}

		customers, err := repo.withinRangeInKm(rangeInKm, p)
		assert.NoError(err)
		assert.NotEmpty(customers)
		for _, customer := range customers {
			distance := geopoint.DistanceInKm(p, customer.Location())
			cond := func() bool { return distance <= rangeInKm }
			assert.Condition(cond, fmt.Sprintf("expected %f to be less than %f", distance, rangeInKm))
		}
	})
}
