package hash

import (
	"testing"

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

	value := ht.Get("key1")

	assert.Equal(t, 1, value)

	ht.Remove("key1")

	assert.Equal(t, 0, ht.Get("key1"))
}
