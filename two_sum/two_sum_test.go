package two_sum

import (
	"../two_sum"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestTwoSum(t *testing.T) {
	sourceInts := []int{1, 5, 2, 3}
	targetSum := 7
	result := []int{0, 1}
	indexes, err := two_sum.TwoSum(sourceInts, targetSum)
	if err != nil {
		log.Fatal(err)
	}
	assert.EqualValues(t, result, indexes)
}
