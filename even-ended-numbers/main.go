package main

import (
	"fmt"
	"strconv"
)

func main() {

	//how many even-ended numbers result from multiplying two four digit numbers
	count := 0
	for a := 1000; a < 10000; a++ {
		for b := a; b < 10000; b++ {
			n := a * b
			s := strconv.FormatInt(int64(n), 10)

			startStr := s[0]
			endStr := s[len(s)-1]

			start, _ := strconv.ParseFloat(string(startStr), 64)
			end, _ := strconv.ParseFloat(string(endStr), 64)

			// if (int64(start)%2 == 0) && (int64(end)%2 == 0) {
			// 	count++
			// }

			if start == end {
				count++
			}

		}

	}

	fmt.Println(count)

}
