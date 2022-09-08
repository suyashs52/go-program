package main

import "fmt"

func main() {

	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	var str string
	for i := 0; i < size; i++ {

		fmt.Scanln(&str)
		fmt.Println("sdfasdf--------")
	}

	arr[0] = 1

	fmt.Println(str)

}
