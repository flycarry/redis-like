package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {
	testTable := []struct {
		command string
		result  string
	}{
		{"set hello world", "+OK"},
		{"get hello", "world"},
		{"set hello nice", "+OK"},
		{"get hello", "nice"},
		{"del hello", "+OK"},
		{"del hello", "-error: key do not exist"},
		{"get hello", "-error: key do not exist"},
		{"get hello world", "-error: invalid number of parameters"},
		{"lget hello", "-error: method not support"},
	}

	testFun := func() {
		for _, everyTest := range testTable {
			result := Process(everyTest.command)
			assert.Equal(t, everyTest.result, result, "failed")
		}
	}
	testFun()
}

func BenchmarkProcess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Process("set hello world")
		Process("get hello")
		Process("del hello")
	}
}
