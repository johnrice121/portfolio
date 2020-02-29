package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"
)

var start, end uint

const maxTestRange = math.MaxUint64

// SmallestNum is given an int A
// Returns the smallest integer comprised of 1s and 0s of which A is a factor.
// Returns a string description of the results
// Returns any error found, nil by default
func SmallestNum(A uint64) (uint64, string, error) {
	var num uint64
	var nStr string
	for num = 1; num < maxTestRange; num++ {
		nStr = fmt.Sprintf("%b", num)
		bInt, err := strconv.Atoi(nStr)
		if err != nil {
			log.Fatalf("Unable to convert num to binary int: %v\n", err)
		}

		bUint := uint64(bInt)
		if bUint%A == 0 {
			return bUint, fmt.Sprintf("%d is a factor of %s\n", A, nStr), nil
		}
	}
	errStr := fmt.Sprintf("Unable to find binary integer less than %s where %d is a factor", nStr, A)
	return 0, errStr, fmt.Errorf(errStr)
}

func init() {
	flag.UintVar(&start, "A", 0, "Factor A, single value or start of range")
	flag.UintVar(&end, "B", 0, "Factor B, end of value range")
}

func main() {
	description := "------------------------------------------------\n"
	description += "See README.md for a description of this program.\n"
	description += "Program returns the smallest integer comprised of 1s and 0s of which A is a factor\n"
	description += "You may specify a single factor: -A 100,\n or a range of factors: -A 1 -B 100\n"
	description += "------------------------------------------------\n"

	flag.Parse()
	if flag.NFlag() < 1 {
		fmt.Print(description)
		fmt.Println("Please choose a factor")
		flag.Usage()
		return
	}
	if (flag.NFlag() == 1 && start >= maxTestRange) ||
		(flag.NFlag() == 2 && (start >= maxTestRange || end >= maxTestRange)) {
		fmt.Print(description)
		fmt.Println("Please choose smaller factors.")
		flag.Usage()
		return
	}
	if (flag.NFlag() == 1 && start < 1) ||
		(flag.NFlag() == 2 && (start < 1 || end < 1)) {
		fmt.Print(description)
		fmt.Println("Please choose factors greater than 0.")
		flag.Usage()
		return
	}

	if end < start {
		num, _, err := SmallestNum(uint64(start))
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(num)
	} else {
		for ; start <= end; start++ {
			_, str, err := SmallestNum(uint64(start))
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print(str)
		}
	}
}
