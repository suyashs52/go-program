package main

import "fmt"

func main() {
	fmt.Println("control flow")

	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
	}

	x := 1
	fmt.Println("\nwith Condition only")
	for x < 5 {
		fmt.Println(x)
		x++
	}

	fmt.Println("\nwith if ")
	for {
		if x < 10 {
			fmt.Println(x)
			x++
		} else {
			break
		}
	}

	for x := 42; x == 42; {
		fmt.Println("42")
		break
	}
	x = 42
	if x == 40 {
		fmt.Println("Our value is 40")
	} else if x == 42 {
		fmt.Println("in else if")
	}

	switch {
	case false:
		fmt.Println("This should  not print ")
	case 2 == 2:
		fmt.Println("This value should print , value is 2")
		fallthrough
	case 3 == 3:
		fmt.Println("3==3")
	case 4 == 4:
		fmt.Println(4 == 4)
	case 3 == 3, 4 == 6:
		fmt.Println("multiple experssion")
	}

	str := "james bond"

	if str == "james bond" {
		fmt.Println("reference is same")
	}
}
