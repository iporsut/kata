package poker

import (
    "github.com/iporsut/kata/poker/card"
)

type Cards []card.Card

func (cards Cards) Len() int {
    return len(cards)
}

func (cards Cards) Less(i, j int) bool {
    return cards[i].Value() <= cards[j].Value()
}

func (cards Cards) Swap(i, j int) {
    cards[i], cards[j] = cards[j], cards[i]
}
func IsFlush(cards []card.Card) bool {
    suite := cards[0].Suite()
    for _, card := range cards {
        if card.Suite() != suite {
            return false
        }
    }
    return true
}
