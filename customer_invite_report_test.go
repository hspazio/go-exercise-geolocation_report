package main

import (
	"fmt"
	"testing"

	"github.com/hspazio/customers_in_range/geopoint"
	"github.com/stretchr/testify/assert"
)

func TestCustomerInviteReport(t *testing.T) {
	assert := assert.New(t)

	t.Run("generates a report", func(t *testing.T) {
		reference := geopoint.Degrees{53.3393, -6.2576841}
		rangeInKm := 100.0
		repo := &customerRepository{"customers.txt"}

		report := customerInviteReport{
			reference: reference,
			rangeInKm: rangeInKm,
			repo:      repo,
		}

		lines, err := report.generate()
		assert.NoError(err)

		lastID := 0
		for _, line := range lines {
			cond := func() bool { return lastID < line.id }
			assert.Condition(cond, fmt.Sprintf("expected %d to be less than %d", lastID, line.id))
			assert.NotEmpty(line.name)
			lastID = line.id
		}
	})
}
