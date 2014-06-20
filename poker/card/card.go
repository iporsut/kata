package card

import (
    "github.com/iporsut/kata/poker/card/suite"
    "github.com/iporsut/kata/poker/card/value"
)

type Card struct {
    value value.Value
    suite suite.Suite
}

func (c Card) Value() value.Value {
    return c.value
}

func (c Card) Suite() suite.Suite {
    return c.suite
}

func New(value value.Value, suite suite.Suite) Card {
    return Card{value, suite}
}
