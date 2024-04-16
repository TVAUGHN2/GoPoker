package hand

import (
	"fmt"
	"sort"

	"github.com/tvaughn2/GoPoker/api/resource/card"
)

//*** Hand Logic | TODO: Move to own file ***//

type HandValue int

const (
	HighCard HandValue = iota
	Pair
	TwoPairs
	Set
	Straight
	Flush
	FullHouse
	Quads
	StraightFlush
	RoyalFlush
)

type HandFace string

var HandFaces = map[HandValue]HandFace{
	HighCard:      "High Card",
	Pair:          "Pair",
	TwoPairs:      "Two Pairs",
	Set:           "Three of a Kind",
	Straight:      "Straight",
	Flush:         "Flush",
	FullHouse:     "Full House",
	Quads:         "Four of a Kind",
	StraightFlush: "Straight Flush",
	RoyalFlush:    "Royal Flush",
}

type Hand struct {
	Value HandValue
	Face  HandFace
	Cards []*card.Card
}

func IsInvalidHand(orderedCards []*card.Card) bool {
	if len(orderedCards) != 5 {
		return true
	}

	prevCard := orderedCards[0]
	if prevCard.IsInvalidCard() {
		return true
	}

	for _, card := range orderedCards[1:] {
		if card.IsInvalidCard() {
			return true
		}
		if card.SameAs(prevCard) {
			return true
		}
		prevCard = card
	}

	return false
}

func NewHand(cards []*card.Card) *Hand {
	// Order cards for illogical check and straight matching
	sort.Slice(cards, func(i, j int) bool {
		return cards[i].Value < cards[j].Value
	})

	if IsInvalidHand(cards) {
		fmt.Printf("Invalid hand input! %s\n", card.StringifyCards(cards))
		return nil
	}

	var HandValue HandValue = NewHandValue(cards)
	return &Hand{HandValue, HandFaces[HandValue], cards}
}

// Due to frequency, sacrificing readability for optimization
func NewHandValue(orderedCards []*card.Card) HandValue {
	var result HandValue = Pair

	// Map by card value for quicker matching
	// Impossible to have > 4 of same card value
	var cardValues = make(map[card.CardValue][]*card.Card)
	for _, card := range orderedCards {
		cardValues[card.Value] = append(cardValues[card.Value], card)
	}

	switch len(cardValues) {
	case 5:
		var isFlush bool = true
		var isStraight bool = true
		var prevCard *card.Card = orderedCards[0]
		for i, card := range orderedCards[1:] {
			if card.Suit != prevCard.Suit {
				isFlush = false
			}
			if card.Value-1 != prevCard.Value {
				// Edge case of ace-5 straight
				if !(i == 3 && card.Value == card.CvA && orderedCards[0].Value == card.Cv2) {
					isStraight = false
				}
			}
			prevCard = card
		}
		switch isFlush {
		case true:
			switch isStraight {
			case true:
				if orderedCards[0].Value == card.CvT {
					result = RoyalFlush
				} else {
					result = StraightFlush
				}
			case false:
				result = Flush
			}
		case false:
			if isStraight {
				result = Straight
			} else {
				result = HighCard
			}
		}
	case 4:
		result = Pair
	case 3:
		result = TwoPairs
		for _, cardValue := range cardValues {
			if len(cardValue) > 2 {
				result = Set
			}
		}
	case 2:
		result = FullHouse
		for _, cardValue := range cardValues {
			if len(cardValue) > 3 {
				result = Quads
			}
		}
	}

	return result
}

//*** End Hand Logic ***//
