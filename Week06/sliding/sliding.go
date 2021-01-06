package sliding

import (
	"fmt"
	"sync"
	"time"
)

var BUCKETSNUMBER float64 = 10
var BUCKETSTIME int64 = 10

type Number struct {
	Buckets map[int64]*numberBucket
	Mutex   *sync.RWMutex
}

type numberBucket struct {
	Value float64
}

// 初始化
func NewNumber() *Number {
	r := &Number{
		Buckets: make(map[int64]*numberBucket),
		Mutex:   &sync.RWMutex{},
	}
	return r
}

func (r *Number) getCurrentBucket() *numberBucket {
	now := time.Now().Unix()
	var bucket *numberBucket
	var ok bool
	if bucket, ok = r.Buckets[now]; !ok {
		bucket = &numberBucket{}
		r.Buckets[now] = bucket
	}
	return bucket
}

func (r *Number) removeOldBuckets() {
	now := time.Now().Unix()

	for timestamp := range r.Buckets {
		if timestamp <= now-BUCKETSTIME {
			delete(r.Buckets, timestamp)
		}
	}
}

func (r *Number) Increment(i float64) {
	if i == 0 {
		return
	}
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	b := r.getCurrentBucket()
	b.Value += i
	fmt.Println(i)
	r.removeOldBuckets()

}

func (r *Number) UpdateMax(n float64) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	b := r.getCurrentBucket()
	if n > b.Value {
		b.Value = n
	}
	r.removeOldBuckets()
}

func (r *Number) Sum(now time.Time) float64 {
	sum := float64(0)

	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	for timestamp, bucket := range r.Buckets {
		if timestamp >= now.Unix()-BUCKETSTIME {
			sum += bucket.Value
		}
	}
	fmt.Println("sum:", sum)
	return sum
}

func (r *Number) Max(now time.Time) float64 {
	var max float64

	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	for timestamp, bucket := range r.Buckets {
		if timestamp >= now.Unix()-BUCKETSTIME {
			if bucket.Value > max {
				max = bucket.Value
			}
		}
	}
	fmt.Println("max:", max)
	return max
}

func (r *Number) Avg(now time.Time) float64 {
	avg := r.Sum(now) / BUCKETSNUMBER
	fmt.Println("avg:", avg)
	return avg
}
