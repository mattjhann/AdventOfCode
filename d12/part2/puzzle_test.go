package part2

import (
	"testing"
)

func BenchmarkChallengeInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoPuzzle("challenge_input.txt")
	}
}

func BenchmarkProdInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoPuzzle("prod_input.txt")
	}
}

func BenchmarkTestInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoPuzzle("test_input.txt")
	}
}
