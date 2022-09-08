package main

import "fmt"

func main() {
	// your code goes here
	var testCase int
	fmt.Scan(&testCase)

	for i := 0; i < testCase; i++ {
		var lens int
		fmt.Scan(&lens)

		var data [lens]int
		//data := make([]int, lens, lens)

		for j := 0; j < lens; j++ {
			fmt.Scan(&data[j])
		}

		var m = make(map[int]int)
		//len: 5 so most freq should be less than 4
		//N>=M*2-1 so feasible
		b := true
		for j := 0; j < lens; j++ {
			//fmt.Println(data[j])
			m[data[j]] = m[data[j]] + 1

			if (m[data[j]]*2 - 1) > lens {
				b = false
				break
			}

		}

		if b {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}

}
