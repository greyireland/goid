package goid

import (
	"testing"
)

func TestID(t *testing.T) {
	id := ID()
	if id == 0 {
		t.Fatalf("get goroutine id failed")
	}
}
func TestGoID(t *testing.T) {
	id := GoID()
	if id == 0 {
		t.Fatalf("get slow id failed")
	}

}
func BenchmarkID(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ID()
		}
	})

}
func BenchmarkGoID(b *testing.B) {
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			GoID()
		}
	})
}
