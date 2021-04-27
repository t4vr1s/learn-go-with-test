package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	// tengo
	repeated := Repeat("a", 5)
	// esperado
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("expected %q got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	iter := Repeat("a", 4)
	fmt.Println(iter)
	//Output: aaaa

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10000)
	}
}
