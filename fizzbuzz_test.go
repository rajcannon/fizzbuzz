package fizzbuzz_test

import (
	"fmt"
	"testing"

	"github.com/rajcannon/fizzbuzz"
)

//function name always beings with Tesxx takes inout of *testing.T and returns nothing
func TestFizz(t *testing.T) {
	result, ok := fizzbuzz.Fizzbuzz(1)

	if ok {
		t.Logf("The value %v should not have returned true", 1)
		t.FailNow()
	}

	result, ok = fizzbuzz.Fizzbuzz(3)

	if !ok {
		t.Logf("The value %v should  have returned true", 3)
		t.Fail()
	}

	if result != "fizz" {
		t.Logf("The result should have been fizz ")
		t.Fail()
	}
}

func ExampleFizzbuzz() {
	result, _ := fizzbuzz.Fizzbuzz(3)
	fmt.Println(result)
	// Output: fizz
}
