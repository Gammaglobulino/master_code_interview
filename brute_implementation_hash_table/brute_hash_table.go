package brute_implementation_hash_table

import (
	"bytes"
	"encoding/gob"
)

type keyValuePair struct {
	key   string
	value int
}

type HashTable struct {
	buckets  [][]keyValuePair
	size     int
	capacity int
}

func (ht *HashTable) DeepCopy() *HashTable {
	copiedBuffer := &bytes.Buffer{}
	newEncoder := gob.NewEncoder(copiedBuffer)
	newEncoder.Encode(ht)
	copiedHashTable := NewHasTable(ht.capacity * 2)
	newDecoder := gob.NewDecoder(copiedBuffer)
	_ = newDecoder.Decode(copiedHashTable)
	return copiedHashTable
}

func NewHasTable(size int) *HashTable {
	return &HashTable{make([][]keyValuePair, size), 0, size}
}

func (ht *HashTable) Hash(key string) uint32 {
	var h uint32
	for _, c := range key {
		h += uint32(c)
		h += h << 10
		h ^= h >> 6
	}
	h += h << 3
	h ^= h >> 11
	h += h << 15
	return h
}

func (ht *HashTable) getIndex(key string) int {
	return int(ht.Hash(key)) % ht.capacity
}

func (ht *HashTable) Size() int {
	return ht.size
}
func (ht *HashTable) Cap() int {
	return ht.capacity
}

func (ht *HashTable) growMe() {
	newGrownTable := ht.DeepCopy()
	ht = newGrownTable
}

func (ht *HashTable) Insert(key string, value int) bool {
	index := ht.getIndex(key)
	chain := ht.buckets[index]
	for i, kv := range chain {
		if kv.key == key {
			node := &chain[i]
			node.value = value
			return false
		}
	}
	if ht.size == ht.capacity {
		ht.growMe()
	}
	node := keyValuePair{
		key:   key,
		value: value,
	}
	chain = append(chain, node)
	ht.buckets[index] = chain
	ht.size++
	return true
}
