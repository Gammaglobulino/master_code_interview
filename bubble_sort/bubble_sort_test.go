package bubble_sort

import (
	"../bubble_sort"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	vec := []int{1, 2, 6, 7, 8, 5}
	sorted := []int{1, 2, 5, 6, 7, 8}
	bubble_sort.Sort(vec)
	assert.ElementsMatch(t, vec, sorted)
	fmt.Println(vec)
	fmt.Println(sorted)
}
