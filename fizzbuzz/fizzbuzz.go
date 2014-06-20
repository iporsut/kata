package fizzbuzz

import (
	"fmt"
)

type Fizzbuzz int

func New(number int) Fizzbuzz {
	return Fizzbuzz(number)
}

func (fb Fizzbuzz) String() string {
	switch {
	case (fb % 15) == 0:
		return "FizzBuzz"
	case (fb % 3) == 0:
		return "Fizz"
	case (fb % 5) == 0:
		return "Buzz"
	}
	return fmt.Sprint(int(fb))
}
