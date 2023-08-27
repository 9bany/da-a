package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashValue(t *testing.T) {
	value := Hash("1234567890")
	assert.LessOrEqual(t, value, hash_length)
}
