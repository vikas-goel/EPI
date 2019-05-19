// You are reading a sequence of strings. You know a priori that more than
// half the strings are repetitions of a single string (the "majority element")
// but the positions where the majority element occurs are unknown. Write a
// program that makes a single pass over the sequence and identifies the
// majority element. For example, if the input is (b,a,c,a,a,b,a,a,c,a), then
// a is the majority element (it appears in 6 out of the 10 places).

package main

import "fmt"

func Majority(sequence []rune) (major rune) {
	if len(sequence) == 0 {
		return
	}

	// In the majority sequence, the majority element is more than the
	// the total of all other elements. So, we can categorize into two
	// buckets, viz. majority and non-majority. This makes it easy to
	// track the frequency in terms of black and white. Start with the
	// first element as majority.
	major = sequence[0]
	for i, frequency := 1, 1; i < len(sequence); i++ {
		if frequency == 0 {
			major = sequence[i]
			frequency++
		} else if major == sequence[i] {
			frequency++
		} else {
			frequency--
		}
	}

	return
}

func main() {
	seq := []rune{'A', 'B', 'C', 'A', 'A', 'B', 'A', 'A', 'C', 'A'}
	fmt.Printf("Majority for sequence %c is %c.\n", seq, Majority(seq))
}
