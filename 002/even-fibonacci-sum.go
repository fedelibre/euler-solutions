package main

import "fmt"

func main() {
	var s []int
	var new int
	for s = []int{1, 2}; new < 4000000; {
		last := s[len(s)-1]
		secondlast := s[len(s)-2]
		new = secondlast + last
		// The condition in the for statement will allow for a number > 4milion
		// So I need this if statement to be safe.
		if new < 4000000 {
			s = append(s, new)
		}
	}
	//fmt.Println("Fibonacci sequence:", s)

	var sEven []int
	for _, number := range s {
		if number%2 == 0 {
			sEven = append(sEven, number)
		}
	}
	//fmt.Println("Even numbers in Fibonacci sequence:", sEven)

	evenSum := 0
	for _, value := range sEven {
		evenSum += value
	}
	fmt.Println(evenSum)
}
