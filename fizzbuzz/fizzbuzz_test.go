package fizzbuzz_test

import (
	"fmt"
	"github.com/iporsut/kata/fizzbuzz"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testExpectedList := []struct {
		number int
		expect string
	}{
		{1, "1"},
		{2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"},
		{8, "8"},
		{9, "Fizz"},
		{10, "Buzz"},
		{11, "11"},
		{12, "Fizz"},
		{13, "13"},
		{14, "14"},
		{15, "FizzBuzz"},
	}

	for _, testExpected := range testExpectedList {
		result := fizzbuzz.New(testExpected.number)
		if fmt.Sprint(result) != testExpected.expect {
			t.Errorf("Expect %s but %s", testExpected.expect, fmt.Sprint(result))
		}
	}
}
