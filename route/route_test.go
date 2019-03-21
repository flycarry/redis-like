package route

import "testing"

type request_response struct {
	request, response string
}

var tt []request_response = []request_response{
	{"get liufei", "not exist"},
	{"set liufei hello", "ok"},
	{"get liufei", "hello"},
	{"del liufei", "delete success"},
	{"get liufei", "not exist"},
}

func TestDojob(t *testing.T) {
	for _, s := range tt {
		r := DoReply(s.request)
		if r != s.response {
			t.Fatal(s, r)
		}
	}
}
func BenchmarkRun(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, s := range tt {
			_ = DoReply(s.request)
		}
	}
}
