# Trial Problem

The following trial problem was given to me as part of a job application.

## Problem

Implement a function that can solve the following problem:

- [Smallest integer composed of only 1s and 0s which is a factor of a given integer 'A'](https://github.com/johnrice121/portfolio/blob/master/factorsOfOnesZeros)

## Examples

If the given number A is 4, the smallest number with 1s and 0s which is divisible by 4 is 100.

If the given number A is 3, the answer would be 111.

## Assumptions

The range of given factors will be integers greater than or equal to 1.

## Solution

This solution is written in go, to build and run it needs to be in the $GOPATH.

To build the solution:

`$ cd factors`

`$ go build`

To run the program:

`$ ./factors -A 10`

or run the program over a range.

`$ ./factors -A 1 -B 100`

Sum unit tests have been written, to run them:

`$ go test`

Currently the highest factor I'm able to process is 395.

## Process of Solution

Given a factor A, my function SmallestNum(A), first loops from 1 to math.MaxUint64 (18446744073709551615) and converts that number into its binary representation in a string type.

Next the program uses the Atoi function to convert the binary string into a base 10 integer.

Following that it does a modulus opperation on the number by the passed in argument A. If it's equal to 0, A is a factor.

The function returns three things; the base 10 number of only 1s and 0s, a string description of the results, and any error that may have been encountered (nil by default).
