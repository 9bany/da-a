package hash

import (
	"testing"

	"github.com/9bany/ds-a/hash"
	"github.com/stretchr/testify/assert"
)

func TestNewArray(t *testing.T) {
	ht := hashtable{}
	assert.Equal(t, 0, len(ht.array))
	ht.renewArray(10)
	assert.Equal(t, 11, len(ht.array))

}

func TestAddValue(t *testing.T) {
	ht := hashtable{}
	ht.Add("key1", 1)

	value := ht.Get("key1")

	assert.Equal(t, 1, value)
}

func TestRemoveValue(t *testing.T) {
	ht := hashtable{}
	ht.Add("key1", 1)
	ht.Add("key2", 2)
	ht.Add("key3", 3)
	ht.Add("key4", 4)

	assert.Equal(t, 4, lenNode(ht.array[hash.Hash("key1")]))
	
	ht.Remove("key2")

	assert.Equal(t, 1, ht.Get("key1"))
	assert.Equal(t, 0, ht.Get("key2"))
	assert.Equal(t, 3, ht.Get("key3"))
	assert.Equal(t, 4, ht.Get("key4"))

	assert.Equal(t, 3, lenNode(ht.array[hash.Hash("key1")]))
}
