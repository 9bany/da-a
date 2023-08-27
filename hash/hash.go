package hash

const hash_length = 1000

// need a uniform function
func Hash(value string) int {
	return len(value) % hash_length
}
