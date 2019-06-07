package main

import "fmt"

// Implement a function that converts a spreadsheet column id to the
// corresponding integer, with "A" corresponding to 1. For example, you should
// return 4 for "D", 27 for "AA", 702 for "ZZ", etc.
func ColumnDecode(column string) (id int) {
	for i := 0; i < len(column); i++ {
		id = id * 26 + int(column[i] - 'A' + 1)
	}

	return
}

// Solve the same problem with "A" corresponding to 0.
func ColumnDecodeVariant(column string) int {
	return ColumnDecode(column) - 1
}

// Implement a function that converts an integer to the spreadsheet column id.
func ColumnEncode(id int) (column string) {
	for letter := 0; id > 0; id = (id - letter) / 26 {
		letter = id % 26

		// For letter Z.
		if letter == 0 {
			letter = 26
		}

		column = string('A' + letter - 1) + column
	}

	return
}

func main() {
	fmt.Print("Column[ ")
	for _, col := range []string{"A", "B", "Y", "Z", "AA", "ZY", "ZZ"} {
		id := ColumnDecode(col)
		rev := ColumnEncode(id)
		fmt.Printf("%v->%v->%v ", col, id, rev)
	}
	fmt.Println("]")
}
