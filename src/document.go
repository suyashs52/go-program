package main

import (
	"fmt"

	"src.com/user/mymath"
)

func main() {
	fmt.Println("documantation")
	fmt.Println(mymath.Sum(3, 4, 5, 6))
}

//run go doc inside package folder
// go doc Sum for reading doc about Sum function
//go doc fmt.Printf
//go help doc
//godoc -http=:8083 //need to install godoc using go install golang.org/x/tools/cmd/godoc@latest
//go doc -http=:8083
//bin/godoc -http=:8083
