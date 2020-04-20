package main

import(
	"fmt"
	"sort"
)
/* TODOs: 
 *  1. Move logic to  separate files. 
 * 	2. Setup REST api.
 *  3. Add concurrency.
 *  4. Apply idiomatic Golang test harness.
 *	5. Convert fmt to logging framework.
 */

//*** Card Logic | TODO: Move to own file ***//

type CardSuit struct {
	Name string
	ShortName string
}

var (
	Clubs = CardSuit{"Clubs", "c"}
	Diamonds = CardSuit{"Diamonds", "d"}
	Hearts = CardSuit{"Hearts", "h"}
	Spades = CardSuit{"Spades", "s"}
)

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
	Name string
	ShortName string
}
var CardFaces = map[CardValue]CardFace {
	Cv2: CardFace{"Two", "2"},
	Cv3 : CardFace{"Three", "3"},
	Cv4 : CardFace{"Four", "4"},
	Cv5 : CardFace{"Five", "5"},
	Cv6 : CardFace{"Six", "6"},
	Cv7 : CardFace{"Seven", "7"},
	Cv8 : CardFace{"Eight", "8"},
	Cv9 : CardFace{"Nine", "9"},
	CvT : CardFace{"Ten", "T"},
	CvJ : CardFace{"Jack", "J"},
	CvQ : CardFace{"Queen", "Q"},
	CvK : CardFace{"King", "K"},
	CvA : CardFace{"Ace", "A"},
}


type Card struct {
	Value CardValue
	Suit CardSuit
	Face CardFace
}


func newCard(cv CardValue, cs CardSuit) *Card {
	return &Card{cv, cs, CardFaces[cv]}
}

func (this *Card) ShortName() string {
	return this.Face.ShortName + this.Suit.ShortName
}

func (this *Card) IsInvalidCard() bool {
	return this.Value == 0 || this.Suit.Name == ""
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
var HandFaces = map[HandValue]HandFace {
	HighCard: "High Card",
	Pair : "Pair",
	TwoPairs : "Two Pairs",
	Set : "Three of a Kind",
	Straight : "Straight",
	Flush  : "Flush",
	FullHouse  : "Full House",
	Quads  : "Four of a Kind",
	StraightFlush  : "Straight Flush",
	RoyalFlush : "Royal Flush",
}

type Hand struct {
	Type HandValue
	Face HandFace
	Cards []*Card
}

func IsInvalidHand(orderedCards []*Card) bool {
	if len(orderedCards) != 5 { return true }

	prevCard := orderedCards[0]
	if prevCard.IsInvalidCard() { return true }

	for _, card := range orderedCards[1:] {
		if card.IsInvalidCard() { return true}
		if card.SameAs(prevCard) { return true }
		prevCard = card
	}

	return false
}

func newHand(cards []*Card) *Hand {
	// Order cards for illogical check and straight matching
	sort.Slice(cards, func(i, j int) bool { 
		return cards[i].Value < cards[j].Value 
	})

	if IsInvalidHand(cards) {
		fmt.Println("Invalid hand input!"); 
		return nil;
	}

	var HandValue HandValue = newHandValue(cards)
	return &Hand{HandValue, HandFaces[HandValue], cards}
}

// Due to frequency, sacrificing readability for optimization
func newHandValue(orderedCards []*Card) HandValue {
	var result HandValue = Pair

	// Map by card value for quicker matching
	// Impossible to have > 4 of same card value
	var cardValues = make(map[CardValue][]*Card)
	for _, card := range orderedCards { 
		cardValues[card.Value] = append(cardValues[card.Value], card)
	}

	switch len(cardValues) {
	case 5:
		var isFlush bool = true;
		var isStraight bool = true;
		var prevCard *Card = orderedCards[0]
		for i, card := range orderedCards[1:] {
			if card.Suit != prevCard.Suit { isFlush = false;}
			if card.Value - 1 != prevCard.Value { 
				// Edge case of ace-5 straight
				if !(i == 3 && card.Value == CvA && orderedCards[0].Value == Cv2) {
					isStraight = false; 
				}
			}
			prevCard = card
		}
		switch isFlush {
		case true:
			switch isStraight {
			case true:
				if (orderedCards[0].Value == CvT) {
					result = RoyalFlush
				} else {
					result = StraightFlush
				}
			case false:
				result = Flush
			}
		case false:
			if (isStraight) {
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
			if len(cardValue) > 2 { result = Set }
		}
	case 2:
		result = FullHouse
		for _, cardValue := range cardValues {
			if len(cardValue) > 3 { result = Quads }
		}
	}
	
	return result	
}
//*** End Hand Logic ***//

//*** Test Methods ***//
func noPair() *Hand {
	return newHand([]*Card{newCard(Cv2, Hearts), newCard(CvJ, Spades), 
		newCard(CvQ, Spades), newCard(CvK, Spades), newCard(CvA, Spades)})
}
func pair() *Hand {
	return newHand([]*Card{newCard(Cv2, Hearts), newCard(Cv2, Spades), 
		newCard(CvQ, Spades), newCard(CvK, Spades), newCard(CvA, Spades)})
}
func twoPairs() *Hand {
	return newHand([]*Card{newCard(Cv2, Hearts), newCard(Cv2, Spades), 
		newCard(CvQ, Spades), newCard(CvQ, Clubs), newCard(CvA, Spades)})
}
func set() *Hand {
	return newHand([]*Card{newCard(Cv2, Hearts), newCard(Cv2, Spades), 
		newCard(CvQ, Spades), newCard(Cv2, Clubs), newCard(CvA, Spades)})
}
func quads() *Hand {
	return newHand([]*Card{newCard(Cv2, Hearts), newCard(Cv2, Spades), 
		newCard(Cv2, Diamonds), newCard(Cv2, Clubs), newCard(CvA, Spades)})
}
func fullHouse() *Hand {
	return newHand([]*Card{newCard(Cv2, Hearts), newCard(Cv2, Spades), 
		newCard(Cv2, Diamonds), newCard(CvA, Clubs), newCard(CvA, Spades)})
}
func straight() *Hand {
	return newHand([]*Card{newCard(Cv3, Spades), newCard(Cv2, Hearts),
		newCard(Cv4, Spades), newCard(Cv5, Clubs), newCard(Cv6, Spades)})
}
func aceLowStraight() *Hand {
	return newHand([]*Card{newCard(CvA, Hearts), newCard(Cv2, Spades), 
		newCard(Cv3, Spades), newCard(Cv4, Clubs), newCard(Cv5, Spades)})
}
func flush() *Hand {
	return newHand([]*Card{newCard(Cv2, Hearts), newCard(Cv4, Hearts), 
		newCard(CvQ, Hearts), newCard(CvK, Hearts), newCard(CvA, Hearts)})
}
func straightFlush() *Hand {
	return newHand([]*Card{newCard(Cv2, Hearts), newCard(Cv3, Hearts), 
		newCard(Cv4, Hearts), newCard(Cv5, Hearts), newCard(Cv6, Hearts)})
}
func royalFlush() *Hand {
	return newHand([]*Card{newCard(CvT, Hearts), newCard(CvJ, Hearts), 
		newCard(CvQ, Hearts), newCard(CvK, Hearts), newCard(CvA, Hearts)})
}
func tooManyCards() *Hand {
	return newHand([]*Card{newCard(CvT, Hearts), newCard(CvJ, Hearts), 
		newCard(CvQ, Hearts), newCard(CvK, Hearts), newCard(CvA, Hearts), newCard(Cv2, Hearts)})
}
func sameCards() *Hand {
	return newHand([]*Card{newCard(CvT, Hearts), newCard(CvJ, Hearts), 
		newCard(CvQ, Hearts), newCard(CvK, Hearts), newCard(CvT, Hearts)})
}

func main() {
	var noPair = noPair()
	var pair = pair()
	var twoPairs = twoPairs()
	var set = set()
	var quads = quads()
	var fullHouse = fullHouse()
	var straight = straight()
	var flush = flush()
	var straightFlush = straightFlush()
	var royalFlush = royalFlush()
	var aceLowStraight = aceLowStraight()

	var tooManyCards = tooManyCards()
	var sameCards = sameCards()

	fmt.Printf("High Card = %s\n", noPair.Face)
	fmt.Printf("Pair = %s\n", pair.Face)
	fmt.Printf("Two Pair = %s\n", twoPairs.Face)
	fmt.Printf("Three of a Kind = %s\n", set.Face)
	fmt.Printf("Four of a Kind = %s\n", quads.Face)
	fmt.Printf("Full House = %s\n", fullHouse.Face)
	fmt.Printf("Straight = %s\n", straight.Face)
	fmt.Printf("Flush = %s\n", flush.Face)
	fmt.Printf("Straight Flush = %s\n", straightFlush.Face)
	fmt.Printf("Royal Flush = %s\n", royalFlush.Face)
	fmt.Printf("Ace Low Straight = %s\n", aceLowStraight.Face)
	fmt.Printf("Too Many Cards ==  %t\n", tooManyCards == nil)
	fmt.Printf("Same Cards == %t\n", sameCards == nil)
}