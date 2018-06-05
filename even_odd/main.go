package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, num := range numbers {
		if num%2 == 0 {
			e := fmt.Sprintf("%d is Even!", num)
			fmt.Println(e)
		} else {
			o := fmt.Sprintf("%d is Odd!", num)
			fmt.Println(o)
		}
	}
}
