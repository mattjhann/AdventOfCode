package part2

import (
	"testing"
)

func TestTestInput(t *testing.T) {
	result := DoPuzzle("test_input.txt", 25)
	if result != 55312 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 55312)
	}
}

func TestProdInput(t *testing.T) {
	result := DoPuzzle("prod_input.txt", 25)
	if result != 198089 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 198089)
	}
}

func BenchmarkProdInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoPuzzle("prod_input.txt", 25)
	}
}

func BenchmarkTestInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoPuzzle("test_input.txt", 75)
	}
}
