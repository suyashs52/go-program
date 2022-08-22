package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("James")

	if s != "Welcome dear James" {
		t.Error("got", s, "expected", "Welcome dear James")
	}
}

func ExampleGreet() {
	fmt.Println(Greet("James"))
	//Output:
	//Welcome dear James
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("James")
	}

}
