package main

import (
	"fmt"
)

func main() {

	//Print the numbers between 1 and 20, one per line
	//If number is divisible by 3, print "fizz"
	//If number is divisible by 5, print "buzz"
	//If number is divisible by both 3 and 5 print "fizzbuzz"
	//Otherwise just print number

	for i := 1; i <= 20; i++ {

		if i%3 == 0 && i%5 == 0 {
			//If number is divisible by both 3 and 5 print "fizzbuzz"
			fmt.Println("fizzbuzz")
			//fmt.Printf("%v = fizzbuzz\n", i)
			continue
		}

		if i%3 == 0 {
			//If number is divisible by 3, print "fizz"
			fmt.Println("fizz")
			//fmt.Printf("%v = fizz\n", i)
			continue
		}

		if i%5 == 0 {
			//If number is divisible by 5, print "buzz"
			fmt.Println("buzz")
			//fmt.Printf("%v = buzz\n", i)
			continue
		}
	}

}
