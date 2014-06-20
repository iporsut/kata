package suite_test

import (
    "fmt"
    "github.com/iporsut/kata/poker/card/suite"
    "testing"
)

func TestPrintSuite(t *testing.T) {
    data := []struct {
        suite    suite.Suite
        expected string
    }{
        {suite.Spades, "S"},
        {suite.Hearts, "H"},
        {suite.Diamonds, "D"},
        {suite.Clubs, "C"},
    }

    for _, d := range data {
        if fmt.Sprint(d.suite) != d.expected {
            t.Errorf("Expect print %v but show %v", d.expected, d.suite)
        }
    }
}
