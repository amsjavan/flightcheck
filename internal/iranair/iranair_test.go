package iranair

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInputJson(t *testing.T) {
	date := "11/09/2022"
	rsp := getInputJson(date)
	assert.Contains(t, rsp, date)
}
