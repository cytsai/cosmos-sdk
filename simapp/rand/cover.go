package rand

import (
	"fmt"
	"sort"
	"testing"
	_ "unsafe"
)

// from https://github.com/segmentio/fasthash
const (
	offset64 = uint64(14695981039346656037)
	prime64  = uint64(1099511628211)
)

func addUint64(h uint64, u uint64) uint64 {
	/*
	h = (h ^ ((u >> 56) & 0xFF)) * prime64
	h = (h ^ ((u >> 48) & 0xFF)) * prime64
	h = (h ^ ((u >> 40) & 0xFF)) * prime64
	h = (h ^ ((u >> 32) & 0xFF)) * prime64
	h = (h ^ ((u >> 24) & 0xFF)) * prime64
	h = (h ^ ((u >> 16) & 0xFF)) * prime64*/
	h = (h ^ ((u >>  8) & 0xFF)) * prime64
	h = (h ^ ( u        & 0xFF)) * prime64
	return h
}

//go:linkname cover testing.cover
var cover testing.Cover
var coverKeys []string
var coverHash uint64
var total float32

func initCoverage() {
	t := 0
	for k, counter := range cover.Counters {
		coverKeys = append(coverKeys, k)
		t += len(counter)
	}
	sort.Strings(coverKeys)
	total = float32(t)
}

func getCoverage() float32 {
	if total == 0.0 {
		return 0.0
	}
	p := 0
	coverHash = offset64
	for _, k := range coverKeys {
		for i, c := range cover.Counters[k] {
			if c > 0 {
				p++
				coverHash = addUint64(coverHash, uint64(i))
			}
		}
	}
	return float32(p) / total
}

func PrintCoverage() {
	//fmt.Printf("COVERAGE %g\n", testing.Coverage())
	fmt.Printf("COVERAGE %g\n", getCoverage())
}

func PrintCoverageStatus() {
	fmt.Printf("%016X\n", coverHash)
}
