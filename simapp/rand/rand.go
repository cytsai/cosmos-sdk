package rand

import (
	"math"
	mrand "math/rand"
)

type Source = mrand.Source
type Rand struct {
	base *mrand.Rand
}

var NewSource = mrand.NewSource
func New(src Source) *Rand {
	return &Rand{base: mrand.New(src)}
}

func (r *Rand) Intn(n int) int {
	if guide == nil {
		return r.base.Intn(n)
	}
	return int(guidedIntn(uint64(n)))
}

func (r *Rand) Int63n(n int64) int64 {
	if guide == nil {
		return r.base.Int63n(n)
	}
	return int64(guidedIntn(uint64(n)))
}

func (r *Rand) Int63() int64 {
	if guide == nil {
		return r.base.Int63()
	}
	return int64(guidedIntn(uint64(math.MaxInt64) + 1))
}

func (r *Rand) Float64() float64 {
	if guide == nil {
		return r.base.Float64()
	}
	return guidedFloat()
}

var Intn    = mrand.Intn
var Int     = mrand.Int
var Int31   = mrand.Int31
var Int63   = mrand.Int63
var Uint32  = mrand.Uint32
var Uint64  = mrand.Uint64
var Shuffle = mrand.Shuffle

func (r *Rand) Perm(n int) []int {
	return r.base.Perm(n)
}

func (r *Rand) Read(p []byte) (n int, err error) {
	return r.base.Read(p)
}

func (r *Rand) GetMathRand() *mrand.Rand {
	return r.base
}
