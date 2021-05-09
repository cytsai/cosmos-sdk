package rand

import (
        mrand "math/rand"
	"os"
        "fmt"
	"runtime"
	"testing"
	"strings"
)

func printState(n int64) {
	fmt.Printf("\n")
	fmt.Printf("COVERAGE %g\n", testing.Coverage())
	fmt.Printf("STATE %d ", n)
	pc := make([]uintptr, 20)
	frames := runtime.CallersFrames(pc[:(runtime.Callers(2, pc) - 2)])
	for {
		frame, more := frames.Next()
		fmt.Printf("%s;", strings.TrimPrefix(frame.Function, "github.com/cosmos/cosmos-sdk/"))
		if !more {
			break
		}
	}
	fmt.Printf("\n")
}

func (r *Rand) Intn(n int) int {
	if r.guide == nil {
		return r.base.Intn(n)
	}
	printState(int64(n))
	var rand int
	if _, err := fmt.Fscanf(r.guide, "%d\n", &rand); err != nil {
		panic(err)
	}
	fmt.Printf("ACTION %d\n", rand)
	return rand
}

func (r *Rand) Int63n(n int64) int64 {
	if r.guide == nil {
		return r.base.Int63n(n)
	}
	printState(n)
	var rand int64
	if _, err := fmt.Fscanf(r.guide, "%d\n", &rand); err != nil {
		panic(err)
	}
	fmt.Printf("ACTION %d\n", rand)
	return rand
}

func (r *Rand) Float64() float64 {
	if r.guide == nil {
		return r.base.Float64()
	}
	printState(0)
	var rand float64
	if _, err := fmt.Fscanf(r.guide, "%g\n", &rand); err != nil {
		panic(err)
	}
	fmt.Printf("ACTION %g\n", rand)
	return rand
}

func (r *Rand) Int63() int64 {
	return r.base.Int63()
}

func (r *Rand) Perm(n int) []int {
	// TODO: consider using Intn
	return r.base.Perm(n)
}

func (r *Rand) Read(p []byte) (n int, err error) {
	return r.base.Read(p)
}

func (r *Rand) GetMathRand() *mrand.Rand {
	// TODO: check importance
	return r.base
}

type Rand struct {
	base  *mrand.Rand
	guide *os.File
}

type Source    = mrand.Source
var  NewSource = mrand.NewSource
var  Int       = mrand.Int
var  Intn      = mrand.Intn
var  Int31     = mrand.Int31
var  Int63     = mrand.Int63
var  Uint32    = mrand.Uint32
var  Uint64    = mrand.Uint64
var  Shuffle   = mrand.Shuffle

func New(src Source) *Rand {
	return &Rand{base: mrand.New(src)}
}

func NewGuided(src Source, guidePath string) *Rand {
	guide, err := os.Open(guidePath)
	if err != nil {
		panic(err)
	}
	return &Rand{base: mrand.New(src), guide: guide}
}
