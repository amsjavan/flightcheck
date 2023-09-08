package airtour

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckAvailability(t *testing.T) {
	res := Response{}
	date := "1401/06/25"
	json.Unmarshal([]byte(Success), &res)
	isAvailable := checkAvailability(&res, date)
	assert.True(t, isAvailable)

	date = "1401/06/20"
	json.Unmarshal([]byte(Success), &res)
	isAvailable = checkAvailability(&res, date)
	assert.False(t, isAvailable)
}

func TestCheckFlight(t *testing.T) {
	date := "1401/06/20"
	res, err := checkFlight()
	assert.NoError(t, err)
	isAvailable := checkAvailability(res, date)
	assert.False(t, isAvailable)

	date = "1401/06/25"
	res, err = checkFlight()
	assert.NoError(t, err)
	isAvailable = checkAvailability(res, date)
	assert.True(t, isAvailable)
}
