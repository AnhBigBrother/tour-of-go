package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func modifyPerson(p Person) {
	p.Age += 10
}

func changeArr(arr []int) {
	fmt.Printf("&arr pass to func: %p\n", &arr)
	fmt.Print("&arr_elements in func:")
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
	fmt.Println("\nIn Go, the value of an array is the address of its first element, and the next elements are next to that address.")
	fmt.Println("This means that when you pass an array to a function, you're actually passing a pointer to the first element of the array.")
	fmt.Println("--> You can actually change the value of each element in an array by passing that array to a function without using a pointer.")

	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("&arr in main: %p\n", &arr)

	fmt.Print("&arr_elements: ")
	for i := range arr {
		fmt.Printf("%p, ", &arr[i])
	}
	fmt.Println()
	fmt.Println("arr: ", arr)
	changeArr(arr)
	fmt.Println("arr: ", arr)

	fmt.Println("\nIt has the same behavior as other object-like types, for example: maps.")
	m := map[string]int{}
	fmt.Println(m)
	changeMap(m)
	fmt.Println(m)

	fmt.Println("\nBut not struct !")
	fmt.Println("The value of a struct-type variable in Go is a copy of the entire struct, including all its fields.")
	fmt.Println("This means that when you pass a struct to a function, you're passing a copy of that struct, and any modifications made to the struct within the function will not affect the original struct.")
	p := Person{Name: "Bro", Age: 24}
	fmt.Println(p, p.Age)
	modifyPerson(p)
	fmt.Println(p, p.Age)
}
