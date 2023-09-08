package meraj

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFailedResponse(t *testing.T) {
	res := Response{}
	json.Unmarshal([]byte(Failed), &res)
	assert.Equal(t, 0, len(res.Result[0].FlightGroups))
	assert.Equal(t, res.Success, true)
}

func TestSuccessResponse(t *testing.T) {
	res := Response{}
	json.Unmarshal([]byte(Success1), &res)
	// assert.Equal(t, "====", res.Result)
	assert.NotEqual(t, 0, len(res.Result[0].FlightGroups))
	assert.Equal(t, res.Success, true)
}
