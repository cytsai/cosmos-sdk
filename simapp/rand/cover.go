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

func sortCoverKeys() {
	t := 0
	for k, C := range cover.Counters {
		coverKeys = append(coverKeys, k)
		t += len(C)
	}
	//coverMap = make([]uint32, t)
	sort.Strings(coverKeys)
}

func getCoverage() float32 {
	p, t := 0, 0
	coverMap.SetInt64(0)
	for _, k := range coverKeys {
		//copy(coverMap[t:], cover.Counters[k])
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
	fmt.Printf("COVERAGE %g\n", getCoverage())
}

func PrintCoverageMap() {
	fmt.Println(coverMap.Text(16))
	//fmt.Println()
}
