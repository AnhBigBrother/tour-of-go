package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

func fibonacci() func() int {
	idx := 0
	i := 0
	j := 1
	return func() int {
		res := i
		if idx == 0 {
			i, j = 1, 1
		} else {
			temp := j
			j = j + i
			i = temp
		}
		idx++
		return res
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
