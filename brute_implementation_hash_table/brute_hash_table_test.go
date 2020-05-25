package brute_implementation_hash_table

import (
	"../brute_implementation_hash_table"
	"bytes"
	"encoding/gob"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashTable_HashFunction(t *testing.T) {
	hasht := brute_implementation_hash_table.NewHasTable(3)
	assert.EqualValues(t, 1484529, hasht.Hash("Andrea"))
}

func TestHashingCollisions(t *testing.T) {
	hasht := brute_implementation_hash_table.NewHasTable(10)
	indexPos1 := hasht.Hash("Printer") % 10
	indexPos2 := hasht.Hash("Scanner") % 10
	indexPos3 := hasht.Hash("Fax") % 10
	indexPos4 := hasht.Hash("Headset") % 10

	assert.NotEqual(t, indexPos1, indexPos2)
	assert.NotEqual(t, indexPos3, indexPos4)

}
func TestInsertKVelement(t *testing.T) {
	hasht := brute_implementation_hash_table.NewHasTable(10)
	assert.True(t, hasht.Insert("Printer", 1))
	assert.True(t, hasht.Insert("Scanner", 1))
}
func TestGrowCapacity(t *testing.T) {
	hasht := brute_implementation_hash_table.NewHasTable(10)
	assert.True(t, hasht.Insert("Printer", 1))
	assert.True(t, hasht.Insert("Scanner", 1))
	hasht.GrowMe()
}

func TestSliceOfSliceDeepCopy(t *testing.T) {
	sliceA := [][]brute_implementation_hash_table.KeyValuePair{{brute_implementation_hash_table.KeyValuePair{}}, {brute_implementation_hash_table.KeyValuePair{
		Key:   "andrea",
		Value: 0,
	},
		brute_implementation_hash_table.KeyValuePair{
			Key:   "andrea",
			Value: 2,
		},
	}, {brute_implementation_hash_table.KeyValuePair{
		Key:   "Stefano",
		Value: 2,
	}}}
	copiedBuffer := &bytes.Buffer{}
	newEncoder := gob.NewEncoder(copiedBuffer)
	err := newEncoder.Encode(sliceA)
	if err != nil {
		t.Fatal(err)
	}
	sliceB := [][]brute_implementation_hash_table.KeyValuePair{}
	newDecoder := gob.NewDecoder(copiedBuffer)
	_ = newDecoder.Decode(&sliceB)
	assert.EqualValues(t, sliceA, sliceB)

}

func TestHashTableCopy(t *testing.T) {
	htA := brute_implementation_hash_table.NewHasTable(5)
	htA.Insert("Andrea", 1)
	htA.Insert("Stefano", 2)
	copiedBuffer := &bytes.Buffer{}
	newEncoder := gob.NewEncoder(copiedBuffer)
	err := newEncoder.Encode(htA)
	if err != nil {
		t.Fatal(err)
	}
	htB := brute_implementation_hash_table.HashTable{nil, 0, 0}
	newDecoder := gob.NewDecoder(copiedBuffer)
	err = newDecoder.Decode(&htB)
	if err != nil {
		t.Fatal(err)
	}
	assert.EqualValues(t, htA, htB)
}

func TestHashTable_DeepCopy(t *testing.T) {
	htA := brute_implementation_hash_table.NewHasTable(5)
	htA.Insert("Andrea", 1)
	htA.Insert("Stefano", 2)
	htB := htA.DeepCopy()
	assert.EqualValues(t, htA, htB)
}

func TestHashTable_GrowMe(t *testing.T) {
	htA := brute_implementation_hash_table.NewHasTable(2)
	htA.Insert("Andrea", 1)
	htA.Insert("Stefano", 2)
	htA.Insert("luca", 2)
	htA.Insert("Paolo", 2)
	htA.Insert("Francesco", 2)
	assert.EqualValues(t, htA.Capacity, 8)
	assert.EqualValues(t, htA.Size, 5)
}
