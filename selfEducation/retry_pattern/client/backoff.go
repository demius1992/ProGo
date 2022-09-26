package main

import (
	"math"
	"math/rand"
	"sync/atomic"
	"time"
)

type Backoff struct {
	attempt uint64
	Factor  float64
	Jitter  bool
	Min     time.Duration
	Max     time.Duration
}

func (b *Backoff) Duration() time.Duration {
	d := b.ForAttempt(float64(atomic.AddUint64(&b.attempt, 1) - 1))
	return d
}

const maxInt64 = float64(math.MaxInt64 - 512)

func (b *Backoff) ForAttempt(attempt float64) time.Duration {
	min := b.Min
	if min <= 0 {
		min = 100 * time.Millisecond
	}
	max := b.Max
	if max <= 0 {
		max = 10 * time.Second
	}
	if min >= max {
		return max
	}
	factor := b.Factor
	if factor <= 0 {
		factor = 2
	}
	minf := float64(min)
	durf := minf * math.Pow(factor, attempt)

	if b.Jitter {
		durf = rand.Float64()*(durf-minf) + minf
	}

	if durf > maxInt64 {
		return max
	}
	dur := time.Duration(durf)
	//keep within bounds
	if dur < min {
		return min
	}
	if dur > max {
		return max
	}
	return dur
}

func (b *Backoff) Reset() {
	atomic.StoreUint64(&b.attempt, 0)
}

func (b *Backoff) Attempt() float64 {
	return float64(atomic.LoadUint64(&b.attempt))
}

func (b *Backoff) Copy() *Backoff {
	return &Backoff{
		Factor: b.Factor,
		Jitter: b.Jitter,
		Min:    b.Min,
		Max:    b.Max,
	}
}
