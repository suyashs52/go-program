package main

import (
	"fmt"
	"runtime"
)

//godoc.org
//z:=4
var z = 4
var k int = 15
var a string = `test string`
var Y string

type testcase int //own type underline int
var y1 testcase

func main() {
	fmt.Print("Hello World")
	var x = 4
	y := 5
	fmt.Print(" ", x, x+y)
	fmt.Print(z)
	z = 9
	fmt.Printf("\n%T\n", k)
	fmt.Printf("\n%b \n", k) //binary
	fmt.Printf("\n%x\n", k)  //hexa decimal
	fmt.Printf("\n%#x\n", k) //0x as format

	s := fmt.Sprintf("%b\t%x\t%o\t%v", k, k, k, k) //string printing
	fmt.Print("-----------")
	fmt.Print(s)

	y1 = 20

	fmt.Printf("\n%T\n", y1)
	//	y1=y; //give error as y1 is testcase type'
	y = int(y1) //it is conversion not casting

	//boolean
	var b bool = 42 == 7.987
	fmt.Print(b)

	//var c  int8 =-129 //byte will give error as 8 byte only
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	oo := "Hello play ground"

	bs := []byte(oo) //string is slice of byte
	//fmt.Println(oo)
	fmt.Println(bs) //H ascii code is 72
	fmt.Printf("%t%T\n", bs)
	fmt.Println("--------------")
	for i := 0; i < len(oo); i++ {
		fmt.Printf("%#U ", oo[i]) //hexadecimal utf value
	}

	for i, v := range oo {
		fmt.Println(i, v)
	}

	const (
		a  = iota //reset to 0
		b1        //1
		c1        //2

	)
	const kl = 10
	const (
		a2 = iota
		b2
		c2 = iota
		d2
	)

	fmt.Println(c2)

}
