// Write a program which takes as input an array of characters, and removes
// each 'b' and replaces each 'a' by two 'd's. Specifically, along with the
// array, you are provided an integer-valued size. Size denotes the number of
// entries of the array that the operation is to be applied to. You do not
// have to worry preserving about subsequent entries. For example, if the array
// is (a,b,a,c, ) and the size is 4, then you can return (d,d,d,d,c). You can
// assume there is enough space in the array to hold the final result.

package main

import "fmt"

func ReplaceRemove(array []rune, size int) {
	// Remove b's in the first pass and keep count of a's.
	aCount, pos := 0, -1
	for i := 0; i < size; i++ {
		if array[i] == 'a' {
			aCount++
		}

		if array[i] != 'b' {
			pos++
			array[pos] = array[i]
		}
	}

	// Get the last index that would be after replacing a's.
	lastIndex := pos + aCount

	// Traverse from the rear end and keep shifting or replacing the
	// elements, as needed.
	for i := pos; i >= 0; i-- {
		if array[i] == 'a' {
			array[lastIndex], array[lastIndex-1] = 'd', 'd'
			lastIndex -= 2
		} else {
			array[lastIndex] = array[i]
			lastIndex--
		}
	}
}

func main() {
	size1, array1 := 7, []rune{'a','c','d','b','b','c','a'}
	size2, array2 := 4, []rune{'a','b','a','c',' '}

	fmt.Printf("%c -> ", array1)
	ReplaceRemove(array1, size1)
	fmt.Printf("%c\n", array1)

	fmt.Printf("%c -> ", array2)
	ReplaceRemove(array2, size2)
	fmt.Printf("%c\n", array2)
}
