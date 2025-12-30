package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", repeated, expected)
	}
}

func TestFullName(t *testing.T) {
	created := FullName("Jon", "Mateer")
	fullName := "Jon Mateer"

	if created != fullName {
		t.Errorf("expected %q but got %q", created, fullName)
	}
}

func ExampleRepeat() {
	chars := Repeat("a", 5)
	fmt.Println(chars)
	// Output: aaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 1)
		FullName("Jon", "Mateer")
	}
}
