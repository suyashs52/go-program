package main

import "fmt"

func main() {
	// your code goes here
	var i, j, k, t int
	fmt.Scanf("%d", &t)

	for s := 0; s < t; s++ {
		fmt.Scanf("%d", &i)
		fmt.Scanf("%d", &j)
		fmt.Scanf("%d", &k)
		fmt.Println("Hello")
		if (i+j < 70) || (j+k < 70) || (i+k < 70) {
			fmt.Println("FAIL")
		} else {
			fmt.Println("PASS")

		}
	}

}
