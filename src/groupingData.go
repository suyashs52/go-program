package main

import "fmt"

func main() {
	fmt.Println("grouping data")
	var x [5]int

	fmt.Println(x)
	x[3] = 5
	fmt.Println(x)
	y := []int{1, 2, 3, 4, 5} //slice
	fmt.Println(y)

	fmt.Println(y[1])

	for i, v := range y {
		fmt.Println(i, v)
	}

	fmt.Println(y[1:2])

	y = append(y, 6, 7, 8)
	fmt.Println(y)
	z := []int{10, 11, 12}

	y = append(y, z...)
	fmt.Println(y)

	y = append(y[:2], y[4:]...)
	fmt.Println(y)

	a := make([]int, 10, 11) //type , len, capacity
	fmt.Println(a)
	fmt.Println(cap(a))
	fmt.Println(len(a))
	a = append(a, 10)
	fmt.Println(len(a)) //make
	a = append(a, 10)
	a = append(a, 10)
	fmt.Println(cap(a)) //double the size if len increase the size
	b := [][]int{a, y}  //2 d array
	fmt.Println(b)

	m := map[string]int{
		"james":       32,
		"money penny": 34,
		"suyah":       45,
	}

	fmt.Println(m)
	fmt.Println(m["suyash"])

	v, ok := m["james"]

	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["james"]; ok {
		fmt.Println("this is map ok value %v", v)
	}

	m["john"] = 33

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "jon")

}
