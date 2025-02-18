package main

import "strconv"

func fizzBuzz(n int) []string {
	array := make([]string, n)
	for i := 1; i < n+1; i++ {
		var toAppend string
		if i%3 == 0 {
			toAppend += "Fizz"
		}
		if i%5 == 0 {
			toAppend += "Buzz"
		}
		if i%3 != 0 && i%5 != 0 {
			toAppend += strconv.Itoa(i)
		}
		array[i-1] = toAppend
	}
	return array
}
