package hash

import (
	"github.com/9bany/ds-a/hash"
)

type hashtable struct {
	array []int
}

func (h *hashtable) renewArray(newLength int) {
	narr := make([]int, 1+len(h.array)+newLength)
	copy(narr, h.array)
	h.array = narr
}

func (h *hashtable) Add(key string, value int) {
	index := hash.Hash(key)
	if len(h.array) < index {
		h.renewArray(index)
	}
	h.array[index] = value
}

func (h *hashtable) Get(key string) int {
	index := hash.Hash(key)
	return h.array[index]
}

func (h *hashtable) Remove(key string) {
	index := hash.Hash(key)
	h.array[index] = 0
}
