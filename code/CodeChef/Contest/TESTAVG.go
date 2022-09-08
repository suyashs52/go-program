package main

import "fmt"

func main() {
	// your code goes here
	//arr := []int{4, 2, -2, 8}

	//fmt.Println(findSubarrays(arr))

	for n := 4; n < 1000; n++ {
		fmt.Println(isStrictlyPalindromic(n))
	}
	//fmt.Println(isStrictlyPalindromic(5))
	fmt.Println(false)

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

func isStrictlyPalindromic(n int) bool {

	for i := 2; i < n-1; i++ {
		val := convertToBase(i, n)
		//fmt.Println(val)
		if isPalind(val) == false {
			return false
		}
	}
	return true

}

func findSubarrays(nums []int) bool {
	var m map[int]bool
	m = make(map[int]bool)

	for i := 0; i < len(nums)-1; i++ {
		_, ok := m[nums[i]+nums[i+1]]

		if ok {
			return true
		} else {
			m[nums[i]+nums[i+1]] = true
		}
	}

	return false
}

func convertToBase(base int, num int) []int {

	count := 0
	size := (num / 2) + 1

	arr := make([]int, 0, size+5)

	//arr= append(arr,0)
	count++

	for num > 0 {
		arr = append(arr, num%base)
		//arr[count] = num % base
		num = num / base
		count++
	}
	arr = append([]int{count}, arr...)
	return arr
}

func isPalind(val []int) bool {
	l := 1
	h := val[0] - 1

	for h > l {
		if val[h] != val[l] {
			return false
		}
		l++
		h--
	}

	return true
}
