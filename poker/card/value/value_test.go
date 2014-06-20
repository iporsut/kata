package value_test

import (
    "fmt"
    "github.com/iporsut/kata/poker/card/value"
    "testing"
)

func TestCompare(t *testing.T) {
    data := []struct {
        lesser  value.Value
        greater value.Value
    }{
        {value.Two, value.Three},
        {value.Three, value.Four},
        {value.Four, value.Five},
        {value.Five, value.Six},
        {value.Six, value.Seven},
        {value.Seven, value.Eight},
        {value.Eight, value.Nine},
        {value.Ten, value.Jack},
        {value.Jack, value.Queen},
        {value.Queen, value.King},
        {value.King, value.Ace},
    }

    for _, d := range data {
        if (d.lesser < d.greater) != true {
            t.Errorf("Expect %v less than %v is true but not.", d.lesser, d.greater)
        }
    }

    for _, d := range data {
        if (d.greater < d.lesser) != false {
            t.Errorf("Expect %v less than %v is false but not.", d.greater, d.lesser)
        }
    }

    for _, d := range data {
        if (d.greater > d.lesser) != true {
            t.Errorf("Expect %v greater than %v is true but not.", d.greater, d.lesser)
        }
    }

    for _, d := range data {
        if (d.lesser > d.greater) != false {
            t.Errorf("Expect %v greater than %v is false but not.", d.lesser, d.greater)
        }
    }
}

func TestCardString(t *testing.T) {
    data := []struct {
        value value.Value
        str   string
    }{
        {value.Two, "2"},
        {value.Three, "3"},
        {value.Four, "4"},
        {value.Five, "5"},
        {value.Six, "6"},
        {value.Seven, "7"},
        {value.Eight, "8"},
        {value.Nine, "9"},
        {value.Ten, "T"},
        {value.Jack, "J"},
        {value.Queen, "Q"},
        {value.King, "K"},
        {value.Ace, "A"},
    }

    for _, d := range data {
        if str := fmt.Sprint(d.value); str != d.str {
            t.Errorf("Expect string of value to %v but to %v", d.str, str)
        }
    }

}
