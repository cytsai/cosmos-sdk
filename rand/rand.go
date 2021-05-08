package rand

import (
        mrand "math/rand"
	"os"
        "fmt"
        "bufio"
	"strings"
	"runtime"
	"testing"
	"strconv"
)

func printState(n int) {
	fmt.Printf("\n")
	fmt.Printf("COVERAGE %g\n", testing.Coverage())
	fmt.Printf("STATE %d ", n)
	pc := make([]uintptr, 20)
	frames := runtime.CallersFrames(pc[:(runtime.Callers(2, pc) - 2)])
	for {
		frame, more := frames.Next()
		fmt.Printf("%d ", frame.PC) //("%d %s %s %d\n", frame.PC, frame.Function, frame.File, frame.Line)
		if !more {
			break
		}
	}
	fmt.Printf("\n")
}

func getInt(guideRNG *bufio.Reader) int64 {
	s, err := guideRNG.ReadString('\n')
	if err != nil {
		panic(err)
	}
	n, err := strconv.ParseInt(strings.TrimSuffix(s, "\n"), 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}

func (r *Rand) Intn(n int) int {
	if !r.guided {
		return r.base.Intn(n)
	}
	printState(n)
	rand := int(getInt(r.guideRNG))
	//fmt.Printf("ACTION %d\n", rand)
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

func (r *Rand) GetMathRand() *mrand.Rand {
	return r.base
}

type Rand struct {
	base     *mrand.Rand
	guided   bool
	guideRNG *bufio.Reader
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

func NewGuided(src Source, guide string) *Rand {
	guidePipe, err := os.Open(guide)
	if err != nil {
		panic(err)
	}
	return &Rand{base: mrand.New(src), guided: true, guideRNG: bufio.NewReader(guidePipe)}
}
