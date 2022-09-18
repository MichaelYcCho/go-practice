package saying

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Michael")
	if s != "Welcome my dear Michael" {
		t.Error("Expected", "Welcome my dear Michael", "Got", s)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Michael"))
	// Output:
	// Welcome my dear Michael
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Michael")
	}

}
