package limit_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/corex-io/limit"
)

type loop int

func (i loop) Run() {
	fmt.Println(i, time.Now().Format("2006/01/02 15:04:05.000"))
}

func Benchmark_Limit(b *testing.B) {
	l := limit.New(limit.Max(5000000000000))
	for i := 0; i < b.N; i++ {
		l.Do(loop(i).Run)
	}
	l.Wait()
}
