package puzzle

import (
	"os"
	"testing"
)

func BenchmarkReadFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		os.ReadFile("challenge_input.txt")
	}
}

func BenchmarkGridCreate(b *testing.B) {
	text, _ := os.ReadFile("challenge_input.txt")
	for i := 0; i < b.N; i++ {
		TextToGrid(string(text))
	}
}

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
