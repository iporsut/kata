package card_test

import (
    "github.com/iporsut/kata/poker/card"
    "github.com/iporsut/kata/poker/card/suite"
    "github.com/iporsut/kata/poker/card/value"
    "testing"
)

func TestNewCard(t *testing.T) {

    data := []struct {
        value value.Value
        suite suite.Suite
    }{
        {value.Two, suite.Spades},
        {value.Three, suite.Spades},
        {value.Four, suite.Spades},
        {value.Five, suite.Spades},
        {value.Six, suite.Spades},
        {value.Seven, suite.Spades},
        {value.Eight, suite.Spades},
        {value.Nine, suite.Spades},
        {value.Ten, suite.Spades},
        {value.Jack, suite.Spades},
        {value.Queen, suite.Spades},
        {value.King, suite.Spades},
        {value.Ace, suite.Spades},

        {value.Two, suite.Hearts},
        {value.Three, suite.Hearts},
        {value.Four, suite.Hearts},
        {value.Five, suite.Hearts},
        {value.Six, suite.Hearts},
        {value.Seven, suite.Hearts},
        {value.Eight, suite.Hearts},
        {value.Nine, suite.Hearts},
        {value.Ten, suite.Hearts},
        {value.Jack, suite.Hearts},
        {value.Queen, suite.Hearts},
        {value.King, suite.Hearts},
        {value.Ace, suite.Hearts},

        {value.Two, suite.Diamonds},
        {value.Three, suite.Diamonds},
        {value.Four, suite.Diamonds},
        {value.Five, suite.Diamonds},
        {value.Six, suite.Diamonds},
        {value.Seven, suite.Diamonds},
        {value.Eight, suite.Diamonds},
        {value.Nine, suite.Diamonds},
        {value.Ten, suite.Diamonds},
        {value.Jack, suite.Diamonds},
        {value.Queen, suite.Diamonds},
        {value.King, suite.Diamonds},
        {value.Ace, suite.Diamonds},

        {value.Two, suite.Clubs},
        {value.Three, suite.Clubs},
        {value.Four, suite.Clubs},
        {value.Five, suite.Clubs},
        {value.Six, suite.Clubs},
        {value.Seven, suite.Clubs},
        {value.Eight, suite.Clubs},
        {value.Nine, suite.Clubs},
        {value.Ten, suite.Clubs},
        {value.Jack, suite.Clubs},
        {value.Queen, suite.Clubs},
        {value.King, suite.Clubs},
        {value.Ace, suite.Clubs},
    }

    for _, d := range data {
        newCard := card.New(d.value, d.suite)
        if v := newCard.Value(); v != d.value {
            t.Errorf("Expect card value %v but %v.", v, d.value)
        }

        if s := newCard.Suite(); s != d.suite {
            t.Errorf("Expect card suite %v but %v.", s, d.suite)
        }
    }
}
