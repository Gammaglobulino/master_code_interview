package google_code_interview_p1

import (
	"../google_code_interview_p1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindCoupleWithSameSum(t *testing.T) {
	sliceA := []int{1, 2, 4, 6}
	sum := 8
	assert.True(t, google_code_interview_p1.FindCoupleWithSameSum(sliceA, sum))
}

func TestFindCoupleWithoutSameSum(t *testing.T) {
	sliceA := []int{1, 2, 4, 3}
	sum := 8
	assert.False(t, google_code_interview_p1.FindCoupleWithSameSum(sliceA, sum))
}
