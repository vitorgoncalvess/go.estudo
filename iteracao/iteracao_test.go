package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a", 6)
	esperado := "aaaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticoes)
	}
}

func ExampleRepetir() {
	repeticoes := Repetir("e", 3)
	fmt.Println(repeticoes)
	// Output: eee
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 5)
	}
}
