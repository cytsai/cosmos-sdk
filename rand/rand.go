package rand

import (
        "fmt"
	"runtime"
        mathrand "math/rand"
)

func trace() {
	pc := make([]uintptr, 20)
	n := runtime.Callers(2, pc) - 2
	frames := runtime.CallersFrames(pc[:n])
	for {
		frame, more := frames.Next()
		//fmt.Printf("%d %s %s %d\n", frame.PC, frame.Function, frame.File, frame.Line)
		fmt.Printf("%d %s\n", frame.PC, frame.Function)
		if !more {
			break
		}
	}
	fmt.Printf("========================================\n")
}


func (r *Rand) Intn(n int) int {
	rand := r.base.Intn(n)
	if r.guided {
		fmt.Printf("%d %d\n", rand, n)
		trace()
	}
	return rand
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
	Counter map[[16]uintptr]int
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
	return &Rand{base: mathrand.New(src)}
}

func NewGuided(src Source) *Rand {
	return &Rand{base: mathrand.New(src), guided: true, Counter: make(map[[16]uintptr]int)}
}
