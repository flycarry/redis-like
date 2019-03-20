package little

import "testing"

func TestRun(t *testing.T) {
    sum := run()
    if sum == 0 {
        t.Error("failed")
    }

}

func BenchmarkRun(b *testing.B) {
    for i := 0; i < b.N; i++ {
        run()
    }
}
