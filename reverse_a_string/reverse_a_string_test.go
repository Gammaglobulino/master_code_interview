package reverse_a_string

import (
	"../reverse_a_string"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	sourceString := "Andrea"
	assert.EqualValues(t, "aerdnA", reverse_a_string.ReverseString(sourceString))
}
