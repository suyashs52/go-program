package main

import "fmt"

func main() {

	var testCase int

	fmt.Scanf("%d", &testCase)
	var result []string

	result = make([]string, testCase)

	for i := 0; i < testCase; i++ {
		var clock int

		fmt.Scanf("%d", &clock)
		var parts int
		fmt.Scan(&parts)
		fmt.Scanf("%d", &parts)

		var types = make([]int, clock, clock)

		for j := 0; j < clock; j++ {

			fmt.Scanf("%d", &types[j])
		}

		if parts*2 < clock {
			result[i] = "NO"

		} else {

			var m map[int]int
			m = make(map[int]int)
			n := true
			for j := 0; j < clock; j++ {
				m[types[j]] += 1

				if m[types[j]] > 2 {
					n = false
					break

				}
			}

			if n == false {
				result[i] = "NO"

			} else {
				result[i] = "YES"

			}
		}

	}
	for i := 0; i < testCase; i++ {
		fmt.Printf("Case #%v: %v\n", i+1, result[i])
	}

}
