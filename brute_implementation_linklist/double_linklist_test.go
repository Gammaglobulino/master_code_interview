package brute_implementation_linklist

import (
	"../brute_implementation_linklist"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDoubleLinkList(t *testing.T) {
	dll := brute_implementation_linklist.NewDoubleLinkList(10)
	dll.Append(20)
	assert.EqualValues(t, 2, dll.Size())
	assert.EqualValues(t, "10->20->nil", dll.Show())
}
