package runtime_test

import (
	"testing"
	_ "unsafe"
)

var nanotimeSink int64

//go:linkname cputicks runtime.cputicks
func cputicks() int64

func BenchmarkCputicks(b *testing.B) {
	var v int64
	for range b.N {
		v = cputicks()
	}

	nanotimeSink = v
}

//go:linkname nanotime runtime.nanotime
func nanotime() int64

func BenchmarkNanotime(b *testing.B) {
	var v int64
	for range b.N {
		v = nanotime()
	}

	nanotimeSink = v
}
