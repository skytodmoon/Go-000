package sliding

import (
	"testing"
	"time"
)

func TestMaxSliding(t *testing.T) {
	number := NewNumber()
	now := time.Now()
	for _, item := range []float64{1, 2, 3, 4, 5} {
		number.Increment(item)
		time.Sleep(1 * time.Second)
	}

	if number.Max(now) != 5 {
		t.FailNow()
	}

}
func TestSumSliding(t *testing.T) {
	number := NewNumber()
	now := time.Now()
	for _, item := range []float64{1, 2, 3, 4, 5, 6, 7, 8, 9} {
		number.Increment(item)
		time.Sleep(1 * time.Second)
	}

	if number.Sum(now) != 45 {
		t.FailNow()
	}

}
func TestAvgSliding(t *testing.T) {
	number := NewNumber()
	now := time.Now()
	for _, item := range []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12} {
		number.Increment(item)
		time.Sleep(1 * time.Second)
	}

	if number.Avg(now) != 7.5 {
		t.FailNow()
	}

}

func BenchmarkRollingNumberIncrement(b *testing.B) {
	n := NewNumber()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		n.Increment(1)
	}
}

func BenchmarkRollingNumberUpdateMax(b *testing.B) {
	n := NewNumber()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		n.UpdateMax(float64(i))
	}
}
