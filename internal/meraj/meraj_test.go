package meraj

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInputJson(t *testing.T) {
	date := "2022-09-10T00:00:00.000Z"
	rsp := getTehranToNajafInputJson(date)
	assert.Contains(t, rsp, date)
}

func TestCheckAvailability(t *testing.T) {
	res := Response{}
	json.Unmarshal([]byte(Failed), &res)
	isAvailable := checkAvailability(&res)
	assert.Equal(t, false, isAvailable)

	json.Unmarshal([]byte(Success1), &res)
	isAvailable = checkAvailability(&res)
	assert.Equal(t, true, isAvailable)
}

func TestCheckFlight(t *testing.T) {
	date := "2022-09-09T00:00:00.000Z"
	res, err := checkFlight(date)
	assert.NoError(t, err)
	isAvailable := checkAvailability(res)
	assert.Equal(t, false, isAvailable)
}
