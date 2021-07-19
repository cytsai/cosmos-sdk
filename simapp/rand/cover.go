package rand

import (
	"fmt"
	"sort"
	"testing"
	"math/big"
	_ "unsafe"
)

//go:linkname cover testing.cover
var cover testing.Cover
var coverKeys []string
var coverMap big.Int

func getCoverage() float32 {
	if coverKeys == nil {
		for k, _ := range cover.Counters {
			coverKeys = append(coverKeys, k)
		}
		sort.Strings(coverKeys)
	}
	p, t := 0, 0
	for _, k := range coverKeys {
		for _, c := range cover.Counters[k] {
			if c > 0 {
				p++
				coverMap.SetBit(&coverMap, t, 1)
			}
			t++
		}
	}
	coverMap.SetBit(&coverMap, t, 1)
	if t == 0 {
		return 0.0
	}
	return float32(p) / float32(t)
}

func PrintCoverage() {
	//fmt.Printf("COVERAGE %g\n", getCoverage())
	fmt.Printf("COVERAGE %g\n", testing.Coverage())
}

func PrintCoverageMap() {
	fmt.Println(coverMap.Text(16))
}
