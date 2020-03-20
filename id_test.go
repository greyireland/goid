package goid

import (
	"testing"
)

func TestGoID(t *testing.T) {
	id := ID()
	if id == 0 {
		t.Fatalf("get goroutine id failed")
	}
}
func TestGoIDSlow(t *testing.T) {
	id := GoID()
	if id == 0 {
		t.Fatalf("get slow id failed")
	}

}
func BenchmarkGoID(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ID()
		}
	})

}
func BenchmarkGoIDSlow(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			GoID()
		}
	})
}
