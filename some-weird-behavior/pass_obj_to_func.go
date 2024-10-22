package main

import "fmt"

func changeArr(arr []int) {
	fmt.Printf("%p\n", &arr)
	for i := range arr {
		fmt.Printf("%p, ", &arr[i])
	}
	fmt.Println()
	arr[0] = 0
}

func changeMap(m map[string]int) {
	m["hehehe"] = 10
}

func main() {

	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", &arr)

	for i := range arr {
		fmt.Printf("%p, ", &arr[i])
	}

	fmt.Println()
	changeArr(arr)

	fmt.Println(arr)

	m := map[string]int{}
	changeMap(m)
	fmt.Println(m)
}
