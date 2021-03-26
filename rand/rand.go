package rand

import (
        //"fmt"
	//"runtime"
        mathrand "math/rand"
)

// Intn 56
// Int63 6
// Int63n 9
// Float64 4
// Perm 2

func (r *Rand) Intn(n int) int {
	/*r.counter += 1
	if (r.counter % 1000) == 0 {
		fmt.Println("\n", r.counter, "\n")
	}*/
	//PCs := make([]uintptr, 16)
	//N := runtime.Callers(2, PCs)
	//fmt.Println(n, N, PCs)
        return r.base.Intn(n)
}

func (r *Rand) Int63() int64 {
	return r.base.Int63()
}

func (r *Rand) Int63n(n int64) int64 {
	return r.base.Int63n(n)
}

func (r *Rand) Float64() float64 {
	return r.base.Float64()
}

func (r *Rand) Perm(n int) []int {
	return r.base.Perm(n)
}

func (r *Rand) Read(p []byte) (n int, err error) {
	return r.base.Read(p)
}

func (r *Rand) GetMathRand() *mathrand.Rand {
	return r.base
}

type Rand struct {
	base    *mathrand.Rand
	guided  bool
	counter int
}

type Source    = mathrand.Source
var  NewSource = mathrand.NewSource
var  Int       = mathrand.Int
var  Intn      = mathrand.Intn
var  Int31     = mathrand.Int31
var  Int63     = mathrand.Int63
var  Uint32    = mathrand.Uint32
var  Uint64    = mathrand.Uint64
var  Shuffle   = mathrand.Shuffle

func New(src Source) *Rand {
	return &Rand{mathrand.New(src), false, 0}
}

func NewGuided(src Source) *Rand {
	return &Rand{mathrand.New(src), true, 0}
}
