package card

import (
	"strings"
	// "encoding/json"
)

//*** Card Logic | TODO: Move to own file ***//

type CardSuit string

const (
	Clubs    CardSuit = "c"
	Diamonds CardSuit = "d"
	Hearts   CardSuit = "h"
	Spades   CardSuit = "s"
)

var CardSuitNames = map[CardSuit]string{
	Clubs:    "Clubs",
	Diamonds: "Diamonds",
	Hearts:   "Hearts",
	Spades:   "Spades",
}

// type CardSuit struct {
// 	Name string
// 	ShortName string
// }

// var (
// 	Clubs = CardSuit{"Clubs", "c"}
// 	Diamonds = CardSuit{"Diamonds", "d"}
// 	Hearts = CardSuit{"Hearts", "h"}
// 	Spades = CardSuit{"Spades", "s"}
// )

// Numerical values make comparison easier
type CardValue int

const (
	Cv2 CardValue = 2 + iota
	Cv3
	Cv4
	Cv5
	Cv6
	Cv7
	Cv8
	Cv9
	CvT
	CvJ
	CvQ
	CvK
	CvA
)

type CardFace struct {
	Name      string
	ShortName string
}

var CardFaces = map[CardValue]CardFace{
	Cv2: CardFace{"Two", "2"},
	Cv3: CardFace{"Three", "3"},
	Cv4: CardFace{"Four", "4"},
	Cv5: CardFace{"Five", "5"},
	Cv6: CardFace{"Six", "6"},
	Cv7: CardFace{"Seven", "7"},
	Cv8: CardFace{"Eight", "8"},
	Cv9: CardFace{"Nine", "9"},
	CvT: CardFace{"Ten", "T"},
	CvJ: CardFace{"Jack", "J"},
	CvQ: CardFace{"Queen", "Q"},
	CvK: CardFace{"King", "K"},
	CvA: CardFace{"Ace", "A"},
}

type Card struct {
	Value CardValue `json: "Value"`
	Suit  CardSuit  `json: "Suit"`
}

// Treat as calc field to to use Card in REST
func (this *Card) Face() CardFace {
	return CardFaces[this.Value]
}

func StringifyCards(cards []*Card) string {
	var str strings.Builder
	str.WriteString(cards[0].ShortName())

	for _, card := range cards[1:] {
		str.WriteString(" ")
		str.WriteString(card.ShortName())
	}

	return str.String()
}

func NewCard(cv CardValue, cs CardSuit) *Card {
	return &Card{cv, cs}
}

func (this *Card) ShortName() string {
	return this.Face().ShortName + string(this.Suit)
}

func (this *Card) IsInvalidCard() bool {
	return this.Value == 0 || string(this.Suit) == ""
}

func (this *Card) SameAs(c *Card) bool {
	return this.ShortName() == c.ShortName()
}

func (this *Card) Equals(c *Card) bool {
	return this.Value == c.Value
}
func (this *Card) GreaterThan(c *Card) bool {
	return this.Value > c.Value
}
func (this *Card) GreaterThanOrEqual(c *Card) bool {
	return this.Value >= c.Value
}
func (this *Card) LessThan(c *Card) bool {
	return this.Value < c.Value
}
func (this *Card) LessThanOrEqual(c *Card) bool {
	return this.Value <= c.Value
}

//*** End Card Logic ***//
