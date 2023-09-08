package iranair

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessResponse(t *testing.T) {

	res := Response{}
	json.Unmarshal([]byte(Success), &res)
	assert.Equal(t, res.Result.SalesDate, "05/09/2022 15:08:00:000")
	assert.Equal(t, res.Success, true)
	assert.Equal(t, res.Result.Journeies[0].DTOs[0].SeatFull, false)
	assert.Equal(t, res.Result.Journeies[0].DTOs[0].SeatAvailability, true)

}
