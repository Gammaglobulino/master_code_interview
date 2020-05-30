package brute_implementation_hash_table

import (
	"bytes"
	"encoding/gob"
)

type KeyValuePair struct {
	Key   string
	Value int
}

type HashTable struct {
	Buckets  [][]KeyValuePair
	Size     int
	Capacity int
}

func (ht *HashTable) DeepCopy() HashTable {
	copiedBuffer := &bytes.Buffer{}
	newEncoder := gob.NewEncoder(copiedBuffer)
	err := newEncoder.Encode(ht)
	if err != nil {
		panic(err)
	}
	copiedHashTable := NewHasTable(ht.Capacity)
	newDecoder := gob.NewDecoder(copiedBuffer)
	err = newDecoder.Decode(&copiedHashTable)
	if err != nil {
		panic(err)
	}
	return copiedHashTable
}

func NewHasTable(size int) HashTable {
	ht := HashTable{make([][]KeyValuePair, size), 0, size}

	//fill the table with empty values (needed for GOB serialization)
	for i := 0; i < len(ht.Buckets); i++ {
		ht.Buckets[i] = append(ht.Buckets[i], KeyValuePair{})
	}
	return ht
}

func (ht *HashTable) Hash(key string) uint32 { //
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
	return int(ht.Hash(key)) % ht.Capacity
}

func (ht *HashTable) Cap() int {
	return ht.Capacity
}

func (ht *HashTable) GrowMe() {
	extendedBackets := make([][]KeyValuePair, ht.Capacity)
	for i := ht.Capacity + 1; i < len(extendedBackets); i++ {
		extendedBackets[i] = append(extendedBackets[i], KeyValuePair{"", 0})
	}
	ht.Buckets = append(ht.Buckets, extendedBackets...)
	ht.Capacity = ht.Capacity * 2
}

func (ht *HashTable) Insert(key string, value int) bool {
	index := ht.getIndex(key)
	chain := ht.Buckets[index]
	for i, kv := range chain {
		if kv.Key == key {
			node := &chain[i]
			node.Value = value
			return false
		}
	}
	if ht.Size == ht.Capacity {
		ht.GrowMe()
	}
	node := KeyValuePair{
		Key:   key,
		Value: value,
	}
	chain = append(chain, node)
	ht.Buckets[index] = chain
	ht.Size++
	return true
}
