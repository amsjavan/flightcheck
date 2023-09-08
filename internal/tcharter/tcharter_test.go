package tcharter

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestChec(t *testing.T) {
	msg := check(Saturday, CloseText)
	assert.Equal(t, msg, "")

	msg = check(Saturday, "else")
	assert.Equal(t, msg, FindSaturday)

	msg = check(Saturday, "else")
	assert.Equal(t, msg, "")

	msg = check(Saturday, CloseText)
	assert.Equal(t, msg, CloseSaturday)

	msg = check(Saturday, CloseText)
	assert.Equal(t, msg, "")

	//================================
	msg = check(Sunday, CloseText)
	assert.Equal(t, msg, "")

	msg = check(Sunday, "else")
	assert.Equal(t, msg, FindSunday)

	msg = check(Sunday, "else")
	assert.Equal(t, msg, "")

	msg = check(Sunday, CloseText)
	assert.Equal(t, msg, CloseSunday)

	msg = check(Sunday, CloseText)
	assert.Equal(t, msg, "")
}
