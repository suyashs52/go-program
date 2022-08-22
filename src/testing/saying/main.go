//package that contain greeting function
package saying

import "fmt"

//Greet with Welcome on pass string value
func Greet(s string) string {
	/// go test ./...
	///..\..\bin\godoc -http=:8081
	///http://localhost:8081/pkg/src.com/user/testing/saying/
	//go test -bence . //all bench mark go test -bench Greet
	////go help -testflag
	////go test -cover
	///go test -coverprofile c.out
	//go tool cover -html=c.out
	return fmt.Sprint("Welcome dear ", s)
}
