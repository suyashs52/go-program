package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

func main() {
	fmt.Println("json")
	p1 := person{
		First: "james",
		Last:  "bond",
		Age:   34,
	}
	p2 := person{
		First: "miss",
		Last:  "penny",
		Age:   23,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people) //convert to byte

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bs)
	fmt.Println(string(bs)) //because first chat as lower case didn't count for field

	//unmarshall
	s := `[{"First":"james","Last":"bond","Age":34},{"First":"miss","Last":"penny","Age":23}]`
	bs = []byte(s)
	fmt.Printf("\n%T\n", bs)
	people = []person{}
	err = json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}
	for i, v := range people {
		fmt.Println(" Person Number", i+1)
		fmt.Println(v)
	}

	fmt.Fprintln(os.Stdout, "Hello play ground")
	io.WriteString(os.Stdout, "Hello , Playground\n")

	is := []int{5, 4, 2, 4, 1}
	sort.Ints(is)
	fmt.Println(is)

	p3 := person{
		First: "Q",
		Last:  "A",
		Age:   22,
	}
	people = append(people, p3)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}

type person struct {
	First string
	Last  string
	Age   int
}
type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
