package main

import (
	"fmt"
	"testing"
)

const (
	testStart uint64 = 1
	testEnd   uint64 = 395
)

// TestFactors checks to make sure output is actually a factor of the given A
func TestFactors(t *testing.T) {
	factors := make([]uint64, testEnd-testStart)
	for f := range factors {
		factors[f] = uint64(f) + testStart
	}

	t.Logf("Test range: %d to %d", testStart, testEnd)

	for _, A := range factors {
		result, _, err := SmallestNum(A)
		if err != nil {
			t.Errorf("%v", err)
			break
		}

		if result%A != 0 {
			t.Errorf("Failed, %d is not a factor of %d", A, result)
		}
	}
}

// TestOnesZeros checks to make sure output int only contains 1s and 0s
func TestOnesZeros(t *testing.T) {
	factors := make([]uint64, testEnd-testStart)
	for f := range factors {
		factors[f] = uint64(f) + testStart
	}

	t.Logf("Test range: %d to %d", testStart, testEnd)

	for _, A := range factors {
		ok := true
		result, _, err := SmallestNum(A)
		if err != nil {
			t.Errorf("%v", err)
			break
		}
		resultStr := fmt.Sprintf("%d", result)
		for i := 0; i < len(resultStr); i++ {
			if resultStr[i] != '1' && resultStr[i] != '0' {
				ok = false
				break
			}
		}
		if !ok {
			t.Errorf("Failed, SmallestNum(%d) = %s Not all 1s and 0s", A, resultStr)
		}
	}
}
