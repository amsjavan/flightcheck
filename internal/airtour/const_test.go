package airtour

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessResponse(t *testing.T) {
	res := Response{}
	json.Unmarshal([]byte(Success), &res)
	assert.False(t, res.HasError)
	assert.True(t, res.Result.Successful)
	date := "1401/06/25"
	assert.Equal(t, date, res.Result.FlightOffers[0].Date)
}
