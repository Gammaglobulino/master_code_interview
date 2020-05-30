package selection_sort

import (
	"../selection_sort"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	vec := []int{1, 2, 6, 7, 8, 5}
	sorted := []int{1, 2, 5, 6, 7, 8}
	selection_sort.Sort(vec)
	assert.ElementsMatch(t, vec, sorted)
	fmt.Println(vec)
	fmt.Println(sorted)
}
