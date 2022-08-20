package main

import "fmt"

func main() {

	foo()
	bar("James")
	s1 := woo("miss pennies")
	fmt.Println(s1)
	x, y, z := mouse("Ian", "Musk")
	fmt.Println(x, y, z)

	fmt.Println(pegion(2, 3, 4, 5, 6))

	xi := []int{1, 2, 3, 4, 5}       //slice
	defer fmt.Println(pegion(xi...)) //so it convert slice to unlimited param
	fmt.Println(pegion())            //defer ll execute at last

	persons{
		first: "mike",
	}.speak() //receiver method

	p1 := persons{
		first: "Dr.",
	}
	p1.speak()

	bars(p1) // as bars taking interface agrument, any struck which
	//implement speack() method ll consider implementing human interface

	//annonymous func

	func(x int) {
		fmt.Println("annonymous i am ", x)
	}(10)

	f := func() {
		fmt.Println("Void argument func")
	}
	f()

	f1 := bigbrother()
	fmt.Println(f1())

	fmt.Printf("%T\n", f)

	//callback
	ii := []int{1, 2, 3, 4, 5}

	s := sum(ii...)

	fmt.Println(s)

	fmt.Println(even(sum, ii...))

	xinc := incrementor()
	//xinc is annonymous func

	fmt.Println(xinc())
	fmt.Println(xinc())
	fmt.Println(xinc())
	fmt.Println(xinc())
	fmt.Println(xinc()) //scope of x is here

	//recursion
	fmt.Println(factorial(5))
}

func foo() {
	fmt.Println("I am from foo")
}
func bar(s string) {

	fmt.Println("Hello", s)

}
func woo(s string) string {
	return fmt.Sprint("Hello from woo", s)
}

func mouse(first string, last string) (string, bool, bool) {
	a := fmt.Sprint(first, last, `, says "Hello"`)
	b := false
	c := true

	return a, b, c
}

func pegion(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0

	for _, v := range x {
		sum += v
	}
	return sum
}

type human interface {
	speak()
}

type persons struct {
	first string
}

func (s persons) speak() { //receiver
	fmt.Println(s.first, "speak")
}

func bars(h human) {
	switch h.(type) {
	case persons:
		fmt.Println("I am interface bars", h)
	}

}

func bigbrother() func() int { //return annonymous functino
	return func() int {
		return 451
	}
}

func sum(x ...int) int {

	total := 0

	for _, v := range x {
		total += v
	}

	return total
}

func even(f func(x ...int) int, vi ...int) int {
	var v []int

	for _, val := range vi {
		if val%2 == 0 {
			v = append(v, val)
		}
	}

	return f(v...)

}

func incrementor() func() int { //return function
	var x int

	return func() int {
		x++
		return x
	}
}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
