package insertion_sort

import (
	"../insertion_sort"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	vec := []int{1, 2, 6, 7, 8, 5, 3, 20, 17, 24, 13}
	sorted := []int{1, 2, 3, 5, 6, 7, 8, 13, 17, 20, 24}
	insertion_sort.Sort(vec)
	assert.ElementsMatch(t, vec, sorted)
	fmt.Println(vec)
	fmt.Println(sorted)
}
