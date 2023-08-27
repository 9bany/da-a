package hash

import (
	"github.com/9bany/ds-a/hash"
)

type hashtable struct {
	array []*node
}

func (h *hashtable) renewArray(newLength int) {
	narr := make([]*node, 1+len(h.array)+newLength)
	copy(narr, h.array)
	h.array = narr
}

func (h *hashtable) Add(key string, value int) {
	index := hash.Hash(key)
	if len(h.array) < index {
		h.renewArray(index)
	}
	if h.array[index] == nil {
		h.array[index] = &node{
			key:   key,
			value: value,
		}
	} else {
		lastNode := lastNode(h.array[index])
		lastNode.next = &node{
			key:   key,
			value: value,
		}
	}

}

func (h *hashtable) Get(key string) int {
	index := hash.Hash(key)
	node := linkedLookup(h.array[index], key)
	if node == nil {
		return 0
	}
	return node.value
}

func (h *hashtable) Remove(key string) {
	index := hash.Hash(key)
	lookup(h.array[index])
	new := linkedRemoveNode(h.array[index], key)
	lookup(new)
	h.array[index] = new
}
