package storage

import (
	"testing"
)

type params_result struct {
	params []string
	result string
}

var tt []params_result = []params_result{
	{[]string{"get", "liufei"}, "not exist"},
	{[]string{"set", "liufei", "nihao"}, "ok"},
	{[]string{"get", "liufei"}, "nihao"},
	{[]string{"del", "liufei"}, "delete success"},
	{[]string{"get", "liufei"}, "not exist"},
	{[]string{"lpush", "hello", "world"}, "1"},
	{[]string{"lpop", "hello"}, "world"},
	{[]string{"lpop", "hello"}, "not exist"},
	{[]string{"rpush", "hello", "world"}, "1"},
	{[]string{"rpop", "hello"}, "world"},
	{[]string{"rpop", "hello"}, "not exist"},
	{[]string{"lpush", "hello", "world"}, "1"},
	{[]string{"rpush", "hello", "test"}, "2"},
	{[]string{"lrange", "hello", "0","2"}, "test\nworld"},
	{[]string{"lrange", "hello", "0","1"}, "test"},
	{[]string{"lrange", "hello", "1","2"}, "world"},
	{[]string{"lrange", "hello", "-2","2"}, "test\nworld"},
	{[]string{"lrange", "hello", "-2","-1"}, "test"},
	{[]string{"lrange", "hello", "-2","-2"}, ""},
	{[]string{"lpop", "hello"}, "world"},
	{[]string{"lpop", "hello"}, "test"},
	{[]string{"rpop", "hello"}, "not exist"},
	{[]string{"lpop", "hello"}, "not exist"},
}

func TestDojob(t *testing.T) {
	for _, s := range tt {
		r, err := GetResult(s.params...)
		if err != nil {
			if err.Error() != s.result {
				t.Fatal(s, err)
			}
		} else {
			if r != s.result {
				t.Fatal(s, r)
			}
		}
	}
}
func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, s := range tt {
			_, _ = GetResult(s.params...)
		}
	}
}
