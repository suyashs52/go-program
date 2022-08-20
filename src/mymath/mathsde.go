//package mymath provides math solutions
package mymath

// Sum add unlimited value of type int
func Sum(x ...int) int {
	t := 0

	for i := 0; i < len(x); i++ {
		t += x[i]
	}
	return t

}
