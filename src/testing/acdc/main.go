//package acdc to evaluate many math function
package acdc

//Sum adds the unlimited number of value of type int
func Sum(x ...int) int {
	sum := 0
	///// go test ./... to run all sub folder test
	//////C:\Users\suysingh\Documents\goworkspace> bin/golint ./...
	/////goworkspace\bin>golint ../src
	////go vet ./...
	for _, v := range x {
		sum += v
	}

	return sum
}
