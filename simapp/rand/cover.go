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
var total int

func initCoverage() {
	for k, counter := range cover.Counters {
		coverKeys = append(coverKeys, k)
		total += len(counter)
	}
	sort.Strings(coverKeys)
}

func getCoverage() int {
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
	return p
}

func PrintCoverage() {
	fmt.Printf("COVERAGE %d\n", getCoverage())
}

func PrintCoverageStatus() {
	fmt.Printf("%016X\n", coverHash)
}
