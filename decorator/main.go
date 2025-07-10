package main

import "log"

func main() {

	f := wrapLogger(addTwoNumbers)

	f(1, 2)

}

type addTwoNumberFunc func(a, b int) int

func wrapLogger(function addTwoNumberFunc) addTwoNumberFunc {
	return func(a, b int) int {
		fn := func(a, b int) (res int) {
			defer func() {
				log.Printf("input values, a: %d, b: %d. Result: %d", a, b, res)
			}()

			return function(a, b)
		}

		return fn(a, b)
	}
}

func addTwoNumbers(a, b int) int {
	return a + b
}
