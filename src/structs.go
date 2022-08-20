package main

import "fmt"

type person struct { //user define type person , underline type is struct
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	license bool
}

func main() {

	fmt.Println("struct")

	p1 := person{
		first: "john",
		last:  "miss",
	}
	fmt.Println(p1)

	sa := secretAgent{
		person:  p1,
		license: true,
	}

	fmt.Println(sa.person.first)
	fmt.Println(sa.first)
	a, b := 10, 20

	fmt.Println(a, b)

	//anonymous struck

	per := struct {
		first string
		last  string
		age   int
	}{
		//first: "test",
		last: "meatball",
	}

	fmt.Println(per.first, per.last)

	type foo int
	var bar foo = 10

	fmt.Println(bar)

	bar = 24
	const mash = 44
	bar = mash //no error as mash is annonymous type
	const nomash int = 34
	//bar = nomash //their is an error becuase nomash is of type int

	bar = foo(nomash)

}
