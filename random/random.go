package random

import (
	"math/rand"
)

// IntMinMax gets pseudo-random number in range [min..max).
func IntMinMax(r *rand.Rand, min, max int) int {
	if min >= max {
		panic("IntMinMax: invalid interval")
	}
	return min + r.Intn(max-min)
}

func RandFillBytes(r *rand.Rand, bs []byte) {
	const (
		bitsPerByte    = 8
		bytesPerUint64 = 8
	)
	var x uint64 // bits accumulator
	var n int    // number of random bytes
	for i := range bs {
		if n == 0 {
			x = r.Uint64()
			n = bytesPerUint64
		}
		bs[i] = byte(x)
		x >>= bitsPerByte
		n--
	}
}
