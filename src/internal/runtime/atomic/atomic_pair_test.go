// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package atomic_test

import (
	"internal/runtime/atomic"
	"testing"
	"unsafe"
)

// alignedPair128 returns a 16-byte-aligned *[2]uint64 inside buf. buf
// must be at least 3 uint64s wide so a 16-byte-aligned 16-byte region fits.
func alignedPair128(buf *[3]uint64) *[2]uint64 {
	if uintptr(unsafe.Pointer(&buf[0]))&15 == 0 {
		return (*[2]uint64)(buf[:2])
	}
	return (*[2]uint64)(buf[1:])
}

// misalignedPair128 returns a *[2]uint64 inside buf that is 8-byte aligned
// but not 16-byte aligned, to test the panic path.
func misalignedPair128(buf *[3]uint64) *[2]uint64 {
	if uintptr(unsafe.Pointer(&buf[0]))&15 == 0 {
		return (*[2]uint64)(buf[1:])
	}
	return (*[2]uint64)(buf[:2])
}

func TestCas128(t *testing.T) {
	var buf [3]uint64
	pair := alignedPair128(&buf)

	// Successful CAS from (0, 0) to (1, 2).
	if !atomic.Cas128(pair, 0, 0, 1, 2) {
		t.Fatal("Cas128: should have succeeded from zero")
	}
	if pair[0] != 1 || pair[1] != 2 {
		t.Fatalf("Cas128 corrupt write: got (%d, %d), want (1, 2)", pair[0], pair[1])
	}

	// Mismatch on low half: should fail without writing.
	if atomic.Cas128(pair, 0, 2, 9, 9) {
		t.Fatal("Cas128: should have failed on low-half mismatch")
	}
	// Mismatch on high half: should fail without writing.
	if atomic.Cas128(pair, 1, 0, 9, 9) {
		t.Fatal("Cas128: should have failed on high-half mismatch")
	}
	if pair[0] != 1 || pair[1] != 2 {
		t.Fatalf("Cas128 wrote on failed CAS: got (%d, %d), want (1, 2)", pair[0], pair[1])
	}

	// Concurrent test: 32 goroutines each bump (lo, hi) -> (lo+1, hi-1)
	// 1000 times. The invariant lo + hi == initialHi holds iff every
	// successful CAS updated both halves together.
	const initialHi = uint64(0xdeadbeefcafebabe)
	pair[0], pair[1] = 0, initialHi

	const G, N = 32, 1000
	done := make(chan struct{})
	for g := 0; g < G; g++ {
		go func() {
			for i := 0; i < N; i++ {
				for {
					lo := atomic.Load64(&pair[0])
					hi := atomic.Load64(&pair[1])
					if atomic.Cas128(pair, lo, hi, lo+1, hi-1) {
						break
					}
				}
			}
			done <- struct{}{}
		}()
	}
	for g := 0; g < G; g++ {
		<-done
	}
	if got, want := atomic.Load64(&pair[0]), uint64(G*N); got != want {
		t.Errorf("low half: got %d, want %d", got, want)
	}
	if got, want := atomic.Load64(&pair[1]), initialHi-uint64(G*N); got != want {
		t.Errorf("high half: got %#x, want %#x", got, want)
	}
}

func TestCas128Unaligned(t *testing.T) {
	var buf [3]uint64
	misaligned := misalignedPair128(&buf)

	defer func() {
		err := recover()
		const want = "unaligned 128-bit atomic operation"
		if err == nil {
			t.Fatal("Cas128 on misaligned address did not panic")
		}
		if s, _ := err.(string); s != want {
			t.Fatalf("Cas128: got panic %q, want %q", err, want)
		}
	}()
	atomic.Cas128(misaligned, 0, 0, 0, 0)
}

// alignedPair128p returns a 16-byte-aligned *[2]unsafe.Pointer inside buf.
func alignedPair128p(buf *[3]unsafe.Pointer) *[2]unsafe.Pointer {
	if uintptr(unsafe.Pointer(&buf[0]))&15 == 0 {
		return (*[2]unsafe.Pointer)(buf[:2])
	}
	return (*[2]unsafe.Pointer)(buf[1:])
}

// misalignedPair128p returns a *[2]unsafe.Pointer inside buf that is
// pointer-aligned but not 16-byte aligned, to test the panic path.
func misalignedPair128p(buf *[3]unsafe.Pointer) *[2]unsafe.Pointer {
	if uintptr(unsafe.Pointer(&buf[0]))&15 == 0 {
		return (*[2]unsafe.Pointer)(buf[1:])
	}
	return (*[2]unsafe.Pointer)(buf[:2])
}

func TestCas128p(t *testing.T) {
	var buf [3]unsafe.Pointer
	pair := alignedPair128p(&buf)

	var a, b, c, d int
	pa, pb, pc, pd := unsafe.Pointer(&a), unsafe.Pointer(&b), unsafe.Pointer(&c), unsafe.Pointer(&d)

	// Successful CAS from (nil, nil) to (pa, pb).
	if !atomic.Cas128p(pair, nil, nil, pa, pb) {
		t.Fatal("Cas128p: should have succeeded from nil")
	}
	if pair[0] != pa || pair[1] != pb {
		t.Fatal("Cas128p corrupt write")
	}

	// Mismatch on low half: should fail without writing.
	if atomic.Cas128p(pair, nil, pb, pc, pd) {
		t.Fatal("Cas128p: should have failed on low-half mismatch")
	}
	// Mismatch on high half: should fail without writing.
	if atomic.Cas128p(pair, pa, nil, pc, pd) {
		t.Fatal("Cas128p: should have failed on high-half mismatch")
	}
	if pair[0] != pa || pair[1] != pb {
		t.Fatal("Cas128p wrote on failed CAS")
	}

	// Concurrent test: G goroutines all race to CAS from (nil, nil) to (pa, pb).
	// Exactly one should win — the atomic guarantee ensures no double-write.
	pair[0], pair[1] = nil, nil
	const G = 32
	wins := make(chan bool, G)
	for g := 0; g < G; g++ {
		go func() {
			wins <- atomic.Cas128p(pair, nil, nil, pa, pb)
		}()
	}
	var total int
	for g := 0; g < G; g++ {
		if <-wins {
			total++
		}
	}
	if total != 1 {
		t.Errorf("Cas128p: %d goroutines won the race, want exactly 1", total)
	}
	if pair[0] != pa || pair[1] != pb {
		t.Error("Cas128p: pair not set correctly after concurrent CAS")
	}
}

func TestCas128pUnaligned(t *testing.T) {
	var buf [3]unsafe.Pointer
	misaligned := misalignedPair128p(&buf)

	defer func() {
		err := recover()
		const want = "unaligned 128-bit atomic operation"
		if err == nil {
			t.Fatal("Cas128p on misaligned address did not panic")
		}
		if s, _ := err.(string); s != want {
			t.Fatalf("Cas128p: got panic %q, want %q", err, want)
		}
	}()
	atomic.Cas128p(misaligned, nil, nil, nil, nil)
}

func BenchmarkCas128(b *testing.B) {
	var buf [3]uint64
	pair := alignedPair128(&buf)
	sink = &pair[0]
	for i := 0; i < b.N; i++ {
		atomic.Cas128(pair, 0, 0, 0, 0)
	}
}

func BenchmarkCas128Parallel(b *testing.B) {
	var buf [3]uint64
	pair := alignedPair128(&buf)
	sink = &pair[0]
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			atomic.Cas128(pair, 0, 0, 0, 0)
		}
	})
}
