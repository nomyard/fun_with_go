package main

import "fmt"

func main() {
	//create a new slice with the type of int- store the numbers 0 through 10
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//for each number in the numbers slice
	for _, num := range numbers {
		//if the number is evenly divisible by 2 it's even
		if num%2 == 0 {
			//store string concatenation inside a variable
			//since we are unable to do type conversion within Println
			e := fmt.Sprintf("%d is Even!", num)
			//print that string out to stdout (the monitor)
			fmt.Println(e)
		} else {
			//if it isn't even. Then it's odd

			//store the string concatenation in a variable
			o := fmt.Sprintf("%d is Odd!", num)
			//print the result to stdout
			fmt.Println(o)
		}
	}
}
