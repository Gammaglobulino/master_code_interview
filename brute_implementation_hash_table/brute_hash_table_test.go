package brute_implementation_hash_table

import (
	"../brute_implementation_hash_table"
	"github.com/stretchr/testify/assert"
	"log"
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

	log.Println(indexPos1)
	log.Println(indexPos2)
	log.Println(indexPos3)
	log.Println(indexPos4)
}
func TestInsertKVelement(t *testing.T) {
	hasht := brute_implementation_hash_table.NewHasTable(10)
	assert.True(t, hasht.Insert("Printer", 1))
	assert.True(t, hasht.Insert("Scanner", 1))
	log.Println(hasht)
}
