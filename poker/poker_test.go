package poker_test

import (
    "github.com/iporsut/kata/poker"
    "github.com/iporsut/kata/poker/card"
    "github.com/iporsut/kata/poker/card/suite"
    "github.com/iporsut/kata/poker/card/value"
    "sort"
    "testing"
)

func TestFlush(t *testing.T) {
    data := []struct {
        cards          []card.Card
        expectedResult bool
    }{
        {
            []card.Card{
                card.New(value.Ten, suite.Spades),
                card.New(value.Jack, suite.Spades),
                card.New(value.Queen, suite.Spades),
                card.New(value.King, suite.Spades),
                card.New(value.Ace, suite.Spades),
            },
            true,
        },

        {
            []card.Card{
                card.New(value.Ten, suite.Spades),
                card.New(value.Jack, suite.Spades),
                card.New(value.Queen, suite.Hearts),
                card.New(value.King, suite.Spades),
                card.New(value.Ace, suite.Spades),
            },
            false,
        },
    }

    for _, d := range data {
        if flush := poker.IsFlush(d.cards); flush != d.expectedResult {
            t.Errorf("Expect is flush %v  but %v", d.expectedResult, flush)
        }
    }
}

func TestSortCardByValue(t *testing.T) {
    data := []struct {
        cards       []card.Card
        sortedCards []card.Card
    }{
        {
            []card.Card{
                card.New(value.Ace, suite.Hearts),
                card.New(value.King, suite.Hearts),
                card.New(value.Queen, suite.Hearts),
                card.New(value.Jack, suite.Hearts),
                card.New(value.Ten, suite.Hearts),
            },
            []card.Card{
                card.New(value.Ten, suite.Hearts),
                card.New(value.Jack, suite.Hearts),
                card.New(value.Queen, suite.Hearts),
                card.New(value.King, suite.Hearts),
                card.New(value.Ace, suite.Hearts),
            },
        },

        {
            []card.Card{
                card.New(value.King, suite.Hearts),
                card.New(value.Queen, suite.Hearts),
                card.New(value.Jack, suite.Hearts),
                card.New(value.Ten, suite.Hearts),
                card.New(value.Nine, suite.Hearts),
            },
            []card.Card{
                card.New(value.Nine, suite.Hearts),
                card.New(value.Ten, suite.Hearts),
                card.New(value.Jack, suite.Hearts),
                card.New(value.Queen, suite.Hearts),
                card.New(value.King, suite.Hearts),
            },
        },
    }

    for _, d := range data {
        sort.Sort(poker.Cards(d.cards))
        sortedCards := d.cards
        for i, sortedCard := range sortedCards {
            if sortedCard != d.sortedCards[i] {
                t.Errorf("Expect sorted to %v but %v", d.sortedCards, sortedCards)
            }
        }
    }
}
