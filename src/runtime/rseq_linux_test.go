// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && (amd64 || arm64)

package runtime_test

import (
	"runtime"
	"testing"
)

func TestRseqGetcpuid(t *testing.T) {
	id := runtime.Getcpuid()
	if id < 0 {
		t.Skip("rseq not registered for this thread (kernel < 4.18 or registration failed)")
	}
	ncpu := runtime.NumCPU()
	if int(id) >= ncpu {
		t.Errorf("getcpuid() = %d, want in [0, %d)", id, ncpu)
	}
}

// TestRseqGetcpuidStable verifies that getcpuid returns a plausible value
// across repeated calls on a locked OS thread.
func TestRseqGetcpuidStable(t *testing.T) {
	done := make(chan struct{})
	go func() {
		defer close(done)
		runtime.LockOSThread()
		defer runtime.UnlockOSThread()

		id := runtime.Getcpuid()
		if id < 0 {
			t.Log("rseq not supported, skipping stability check")
			return
		}
		ncpu := runtime.NumCPU()
		// With a locked OS thread the CPU may still change (preempted and
		// rescheduled on a different core), so we only validate the range.
		for i := 0; i < 10000; i++ {
			got := runtime.Getcpuid()
			if got < 0 || int(got) >= ncpu {
				t.Errorf("getcpuid() = %d, want in [0, %d) at iteration %d", got, ncpu, i)
				return
			}
		}
	}()
	<-done
}
