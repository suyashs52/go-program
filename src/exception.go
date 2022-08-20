package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	var ans1, ans2 string

	fmt.Print("Name: ")

	_, err := fmt.Scan(&ans1)
	if err != nil {
		panic(err)
	}
	fmt.Print("Fav Food: ")
	_, err = fmt.Scan(&ans2)
	if err != nil {
		panic(err)
	}
	fmt.Println(ans1, ans2)

	writeFile()
	readFile()

	fmt.Println("------------\nDefer return type example")
	y := deferincr()
	fmt.Println(y)
	fmt.Println("------------\nDefer recover ")
	fRecoverEx()
	fmt.Println("------------\nThrowing error ")

	_, err = sqrts(-10.0)
	if err != nil {
		log.Println(err)

	}

	fmt.Println("------------\n Custom error ")

	cus := custError{
		info: "need coffee",
	}
	errorfoo(cus)
}

func writeFile() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	defer fmt.Println("closed")
	r := strings.NewReader("Wassup")

	io.Copy(f, r)

}

func deferincr() (i int) {
	defer func() {
		i++
	}()
	return 1
}

func readFile() {
	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		//	log.SetOutput(f) write into file
		log.Println("error occured", err)
		//	log.Fatalln("fatal error occur", err)
		//panic(err)
		return

	}

	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		log.Println("error occured", err)
		return
	}

	fmt.Println(string(bs))

}

func fRecoverEx() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover in f ", r)
		}
	}()

	fmt.Println("calling g")
	g(0)

	fmt.Println("return normally from g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("panicking...")
		panic(fmt.Sprintf("value %v", i))
	}
	defer fmt.Println("defer in g ", i)
	fmt.Println("printing in g", i)
	g(i + 1)
}

func sqrts(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math error: sqare root of negative number")
	}

	return 42, nil
}
func sqrt(f float64) (float64, error) {
	if f < 0 {

		//return 0, fmt.Errorf("error in square %v", f)
		//return 0,fmt.Println("error in square func")
		return 0, custError{info: "sqrt error"}
	}

	return 42, nil
}

type custError struct {
	info string
}

func (c custError) Error() string {
	return fmt.Sprintf("here is the error %v", c.info)
}

func errorfoo(e error) {
	fmt.Println(e, "\n")
	fmt.Println(e.(custError).info) //custErrror work as assertion
}
