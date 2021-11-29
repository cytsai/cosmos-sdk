package rand

import (
	"fmt"
	"sort"
	"testing"
	_ "unsafe"
	mrand "math/rand"
)

const (
	BINS = 64
	BITS = 16
)

var perm []int
var hash []int

func updateHash(i int) {
	b := i % BINS
	j := i / BINS
	if j > hash[b] {
		hash[b] = j
	}
}

//go:linkname cover testing.cover
var cover testing.Cover
var coverKeys []string
var max int

func initCoverage() {
	for k, counter := range cover.Counters {
		coverKeys = append(coverKeys, k)
		max += len(counter)
	}
	sort.Strings(coverKeys)
	perm = mrand.New(mrand.NewSource(0)).Perm(max)
	hash = make([]int, BINS)
	for i := range hash {
		hash[i] = -1
	}
}

func getCoverage() int {
	i, c := 0, 0
	for _, k := range coverKeys {
		for _, b := range cover.Counters[k] {
			if b > 0 {
				c++
				updateHash(perm[i])
			}
			i++
		}
	}
	return c
}

func PrintCoverage() {
	fmt.Printf("COVERAGE %d\n", getCoverage())
}

func PrintCoverageStatus() {
	for _, v := range hash {
		fmt.Printf("%d,", v % BITS)
	}
	fmt.Printf("\n")
}
