package tourOfGo

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		temp := first
		first = second
		second = temp + first
		return temp
	}
}

func CheckFibonacci() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
