package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessForList(t *testing.T) {
	testTable := []struct {
		command string
		result  string
	}{
		{"lpush hello 1", "+OK"},
		{"lpush hello 2 3 4 5 6", "+OK"},
		{"rpush hello 2 3 4 5 6", "+OK"},
		{"lpop hello", "6"},
		{"rpop hello", "6"},
		{"lrange hello 0 2", "1) 5\r\n2) 4"},
	}
	for _, test := range testTable {
		result := Process(test.command)
		assert.Equal(t, test.result, result, "test failed")
	}
}
