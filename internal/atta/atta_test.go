package atta

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUrl(t *testing.T) {
	url := "https://api.app.ataair.ir/v1/booking-api/AvailableFlights?source=IKA&target=NJF&day=20&month=6&adultQty=1&childQty=0&infantQty=0"
	day := int32(20)
	rsp := getToNajafUrl(day)
	assert.Equal(t, rsp, url)
}

func TestCheckAvailability(t *testing.T) {
	res := Response{}
	json.Unmarshal([]byte(Failed), &res)
	isAvailable := checkAvailability(&res)
	assert.Equal(t, false, isAvailable)

	// json.Unmarshal([]byte(Success1), &res)
	// isAvailable = checkAvailability(&res)
	// assert.Equal(t, true, isAvailable)
}

func TestCheckFlight(t *testing.T) {
	res, err := checkToNajafFlight(20)
	assert.NoError(t, err)
	isAvailable := checkAvailability(res)
	assert.Equal(t, false, isAvailable)

	res, err = checkToNajafFlight(4)
	assert.NoError(t, err)
	isAvailable = checkAvailability(res)
	assert.Equal(t, false, isAvailable)
}
