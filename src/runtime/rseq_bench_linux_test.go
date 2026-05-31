// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && (amd64 || arm64)

package runtime_test

// Benchmarks comparing rseq-based getcpuid against the traditional
// procPin/procUnpin mechanism for obtaining a stable per-CPU index.
//
// procPin disables preemption (increments m.locks) and returns the P ID.
// getcpuid reads a kernel-maintained TLS field — a single memory load with
// no write to shared state and no preemption disable.
//
// Run with:
//
//	CGO_ENABLED=0 go test -run='^$' -bench=BenchmarkGetcpuid \
//	    -benchtime=5s -count=5 -cpu=1,2,4,8,16,32
//
// The -cpu flag controls GOMAXPROCS. At high core counts the rseq approach
// scales linearly while procPin contends on the m.locks write.

import (
	"runtime"
	"sync"
	"testing"
)

// BenchmarkGetcpuidRseq measures the cost of reading the current CPU ID
// via the rseq TLS field. This is a pure load with no side effects.
func BenchmarkGetcpuidRseq(b *testing.B) {
	if runtime.Getcpuid() < 0 {
		b.Skip("rseq not registered (CGO_ENABLED=1 with glibc, or kernel < 4.18)")
	}
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = runtime.Getcpuid()
		}
	})
}

// BenchmarkGetcpuidProcPin measures the cost of obtaining a stable per-CPU
// index via procPin + procUnpin. This disables and re-enables preemption on
// every call, writing to m.locks twice.
func BenchmarkGetcpuidProcPin(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			pid := runtime.ProcPin()
			runtime.ProcUnpin()
			_ = pid
		}
	})
}

// BenchmarkGetcpuidContendedRseq exercises getcpuid under goroutine contention.
// Each goroutine spins reading the CPU ID while others do the same, showing
// that the rseq load has zero cross-core contention.
func BenchmarkGetcpuidContendedRseq(b *testing.B) {
	if runtime.Getcpuid() < 0 {
		b.Skip("rseq not registered")
	}
	procs := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	ready := make(chan struct{})
	perGoroutine := b.N / procs
	if perGoroutine < 1 {
		perGoroutine = 1
	}
	b.ResetTimer()
	for i := 0; i < procs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-ready
			for j := 0; j < perGoroutine; j++ {
				_ = runtime.Getcpuid()
			}
		}()
	}
	close(ready)
	wg.Wait()
}

// BenchmarkGetcpuidContendedProcPin exercises procPin under goroutine
// contention. Each goroutine pins itself, which writes to m.locks, causing
// cache-line traffic between cores sharing the same M.
func BenchmarkGetcpuidContendedProcPin(b *testing.B) {
	procs := runtime.GOMAXPROCS(0)
	var wg sync.WaitGroup
	ready := make(chan struct{})
	perGoroutine := b.N / procs
	if perGoroutine < 1 {
		perGoroutine = 1
	}
	b.ResetTimer()
	for i := 0; i < procs; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-ready
			for j := 0; j < perGoroutine; j++ {
				pid := runtime.ProcPin()
				runtime.ProcUnpin()
				_ = pid
			}
		}()
	}
	close(ready)
	wg.Wait()
}
