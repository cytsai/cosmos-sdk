package rand

import (
	"os"
	"fmt"
	"runtime"
	"testing"
	"strings"
)

var guide *os.File
var interactive bool

func SetGuide(guidePath string) {
	fi, err := os.Stat(guidePath)
	if err != nil {
		panic(err)
	}
	guide, _ = os.Open(guidePath)
	interactive = !fi.Mode().IsRegular()
}

func PrintCoverage() {
	if interactive {
		fmt.Printf("COVERAGE %g\n", testing.Coverage())
	}
}

func printState(n uint64) {
	fmt.Printf("\n")
	fmt.Printf("COVERAGE %g\n", testing.Coverage())
	fmt.Printf("STATE %d ", n)
	pc := make([]uintptr, 32)
	frames := runtime.CallersFrames(pc[:(runtime.Callers(3, pc) - 2)])
	for {
		frame, more := frames.Next()
		fmt.Printf("%s;", strings.TrimPrefix(frame.Function, "github.com/cosmos/cosmos-sdk/"))
		if !more {
			break
		}
	}
	fmt.Printf("\n")
}

func guidedIntn(n uint64) uint64 {
	if n <= 0 {
		panic("invalid argument to guidedIntn")
	}
	if interactive {
		printState(n)
	}
	var rand uint64
	if _, err := fmt.Fscanf(guide, "%d\n", &rand); err != nil {
		panic(err)
	}
	if interactive {
		fmt.Printf("ACTION %d\n", rand)
	} else if rand < 0 || rand >= n {
		panic("invalid guide file content")
	}
	return rand
}

func guidedFloat() float64 {
	if interactive {
		printState(0)
	}
	var rand float64
	if _, err := fmt.Fscanf(guide, "%g\n", &rand); err != nil {
		panic(err)
	}
	if interactive {
		fmt.Printf("ACTION %g\n", rand)
	} else if rand < 0.0 || rand >= 1.0 {
		panic("invalid guide file content")
	}
	return rand
}
