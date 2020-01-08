package example

import (
	"testing"
	"time"
)

func BenchmarkCustomDate(b *testing.B) {

	for i := 0; i < b.N; i++ {
		customDate(time.Time{})
	}
}
