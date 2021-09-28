package rand

import (
	"fmt"
	"testing"
)

func PrintCoverage() {
	fmt.Printf("COVERAGE %g\n", testing.Coverage())
}
