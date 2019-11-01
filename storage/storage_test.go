package storage

import (
	"fmt"
	"log"
	"sync"
	"testing"
)

const (
	concurrent = 200
)

type params_result struct {
	params []string
	result string
}

var tt []params_result = []params_result{
	//{[]string{"get", "hello"}, "not exist"},
	{[]string{"set", "hello", "nihao"}, "ok"},
	//{[]string{"get", "hello"}, "nihao"},
	//{[]string{"del", "hello"}, "delete success"},
	//{[]string{"get", "hello"}, "not exist"},
	//{[]string{"lpush", "hello", "world"}, "1"},
	//{[]string{"lpop", "hello"}, "world"},
	//{[]string{"lpop", "hello"}, "not exist"},
	//{[]string{"rpush", "hello", "world"}, "1"},
	//{[]string{"rpop", "hello"}, "world"},
	//{[]string{"rpop", "hello"}, "not exist"},
	//{[]string{"lpush", "hello", "world"}, "1"},
	//{[]string{"rpush", "hello", "test"}, "2"},
	//{[]string{"lrange", "hello", "0","2"}, "world\ntest"},
	//{[]string{"lrange", "hello", "0","1"}, "world"},
	//{[]string{"lrange", "hello", "1","2"}, "test"},
	//{[]string{"lrange", "hello", "-2","2"}, "world\ntest"},
	//{[]string{"lrange", "hello", "-2","-1"}, "world"},
	//{[]string{"lrange", "hello", "-2","-2"}, ""},
	//{[]string{"lpop", "hello"}, "world"},
	//{[]string{"lindex", "hello","1"}, "test"},
	//{[]string{"linsert", "hello","test","nihao"}, "ok"},
	//{[]string{"lindex", "hello","1"}, "nihao"},
	//{[]string{"rpop", "hello"}, "test"},
	//{[]string{"lindex", "hello","2"}, "error index"},
	//{[]string{"lpop", "hello"}, "nihao"},
	//{[]string{"rpop", "hello"}, "not exist"},
	//{[]string{"lpop", "hello"}, "not exist"},
}

func TestDoJob(t *testing.T) {
	for _, s := range tt {
		r, err := GetResult(s.params...)
		if err != nil {
			if err.Error() != s.result {
				t.Fatal(s, err)
			} else {
				log.Println(s, r)
			}
		} else {
			if r != s.result {
				t.Fatal(s, r)
			} else {
				log.Println(s, r)
			}
		}
	}
}

func TestConcurrent(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(concurrent)
	m := 0
	for i := 0; i < concurrent; i++ {
		go func(i int) {
			for _, s := range tt {
				s.params[1] = "job" + fmt.Sprint(i)
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
			m++
			log.Println(m)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, s := range tt {
			_, _ = GetResult(s.params...)
		}
	}
}
