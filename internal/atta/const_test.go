package atta

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFailedResponse(t *testing.T) {
	res := Response{}
	json.Unmarshal([]byte(Failed), &res)
	assert.Equal(t, 0, len(res.Result))
	assert.Equal(t, res.Success, true)
}

func TestSuccessResponse(t *testing.T) {
	res := Response{}
	json.Unmarshal([]byte(Success), &res)
	// assert.Equal(t, "====", res.Result)
	assert.Equal(t, res.Result[0].CapacityTitle, "ظرفیت تکمیل است!")
	assert.Equal(t, res.Success, true)
}
