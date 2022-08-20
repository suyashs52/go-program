package main

import "fmt"

func main() {
	fmt.Println("pointer")
	a := 22

	fmt.Println(&a)
	fmt.Printf("\n%T\n", &a)

	b := &a

	fmt.Println(*b)

	*b = 45
	fmt.Println(a)

	x := 0

	pfoo(&x)

	fmt.Println(x)

	p := person{
		first: "MoneyPenny",
	}
	fmt.Println(p)
	changeName(&p)
	fmt.Println(p)
	chgName(p)
	fmt.Println(p.first)
}

type person struct {
	first string
}

func chgName(p person) {
	p.first = "Money Value"
}
func changeName(p *person) {
	p.first = "Money Peny"
	//(*p).first="Money Peny"
}
func pfoo(f *int) {
	*f++

}
