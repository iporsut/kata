package value

import (
    "strconv"
)

type Value int

const (
    Two Value = iota + 2
    Three
    Four
    Five
    Six
    Seven
    Eight
    Nine
    Ten
    Jack
    Queen
    King
    Ace
)

func (v Value) String() (str string) {
    switch v {
    case Two, Three, Four, Five, Six, Seven, Eight, Nine:
        str = strconv.Itoa(int(v))
    case Ten:
        str = "T"
    case Jack:
        str = "J"
    case Queen:
        str = "Q"
    case King:
        str = "K"
    case Ace:
        str = "A"
    }
    return
}
