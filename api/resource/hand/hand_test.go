package hand

import (
	"fmt"

	"github.com/tvaughn2/GoPoker/api/resource/card"
)

// *** Test Methods ***//
func noPair() *Hand {
	return NewHand([]*card.Card{card.NewCard(card.Cv2, card.Hearts), card.NewCard(card.CvJ, card.Spades),
		card.NewCard(card.CvQ, card.Spades), card.NewCard(card.CvK, card.Spades), card.NewCard(card.CvA, card.Spades)})
}
func pair() *Hand {
	return NewHand([]*card.Card{card.NewCard(card.Cv2, card.Hearts), card.NewCard(card.Cv2, card.Spades),
		card.NewCard(card.CvQ, card.Spades), card.NewCard(card.CvK, card.Spades), card.NewCard(card.CvA, card.Spades)})
}
func twoPairs() *Hand {
	return NewHand([]*card.Card{card.NewCard(card.Cv2, card.Hearts), card.NewCard(card.Cv2, card.Spades),
		card.NewCard(card.CvQ, card.Spades), card.NewCard(card.CvQ, card.Clubs), card.NewCard(card.CvA, card.Spades)})
}
func set() *Hand {
	return NewHand([]*card.Card{card.NewCard(card.Cv2, card.Hearts), card.NewCard(card.Cv2, card.Spades),
		card.NewCard(card.CvQ, card.Spades), card.NewCard(card.Cv2, card.Clubs), card.NewCard(card.CvA, card.Spades)})
}
func quads() *Hand {
	return NewHand([]*card.Card{card.NewCard(card.Cv2, card.Hearts), card.NewCard(card.Cv2, card.Spades),
		card.NewCard(card.Cv2, card.Diamonds), card.NewCard(card.Cv2, card.Clubs), card.NewCard(card.CvA, card.Spades)})
}
func fullHouse() *Hand {
	return NewHand([]*card.Card{card.NewCard(card.Cv2, card.Hearts), card.NewCard(card.Cv2, card.Spades),
		card.NewCard(card.Cv2, card.Diamonds), card.NewCard(card.CvA, card.Clubs), card.NewCard(card.CvA, card.Spades)})
}
func straight() *Hand {
	return NewHand([]*card.Card{card.NewCard(card.Cv3, card.Spades), card.NewCard(card.Cv2, card.Hearts),
		card.NewCard(card.Cv4, card.Spades), card.NewCard(card.Cv5, card.Clubs), card.NewCard(card.Cv6, card.Spades)})
}
func aceLowStraight() *Hand {
	return NewHand([]*card.Card{card.NewCard(card.CvA, card.Hearts), card.NewCard(card.Cv2, card.Spades),
		card.NewCard(card.Cv3, card.Spades), card.NewCard(card.Cv4, card.Clubs), card.NewCard(card.Cv5, card.Spades)})
}
func flush() *Hand {
	return NewHand([]*card.Card{card.NewCard(card.Cv2, card.Hearts), card.NewCard(card.Cv4, card.Hearts),
		card.NewCard(card.CvQ, card.Hearts), card.NewCard(card.CvK, card.Hearts), card.NewCard(card.CvA, card.Hearts)})
}
func straightFlush() *Hand {
	return NewHand([]*card.Card{card.NewCard(card.Cv2, card.Hearts), card.NewCard(card.Cv3, card.Hearts),
		card.NewCard(card.Cv4, card.Hearts), card.NewCard(card.Cv5, card.Hearts), card.NewCard(card.Cv6, card.Hearts)})
}
func royalFlush() *Hand {
	return NewHand([]*card.Card{card.NewCard(card.CvT, card.Hearts), card.NewCard(card.CvJ, card.Hearts),
		card.NewCard(card.CvQ, card.Hearts), card.NewCard(card.CvK, card.Hearts), card.NewCard(card.CvA, card.Hearts)})
}
func tooManyCards() *Hand {
	return NewHand([]*card.Card{card.NewCard(card.CvT, card.Hearts), card.NewCard(card.CvJ, card.Hearts),
		card.NewCard(card.CvQ, card.Hearts), card.NewCard(card.CvK, card.Hearts), card.NewCard(card.CvA, card.Hearts), card.NewCard(card.Cv2, card.Hearts)})
}
func sameCards() *Hand {
	return NewHand([]*card.Card{card.NewCard(card.CvT, card.Hearts), card.NewCard(card.CvJ, card.Hearts),
		card.NewCard(card.CvQ, card.Hearts), card.NewCard(card.CvK, card.Hearts), card.NewCard(card.CvT, card.Hearts)})
}

func handValueTestsPass() bool {
	noPair := noPair()
	pair := pair()
	twoPairs := twoPairs()
	set := set()
	quads := quads()
	fullHouse := fullHouse()
	straight := straight()
	flush := flush()
	straightFlush := straightFlush()
	royalFlush := royalFlush()
	aceLowStraight := aceLowStraight()
	tooManyCards := tooManyCards()
	sameCards := sameCards()
	testsPass := true

	if noPair.Value != HighCard {
		fmt.Printf("ERROR: High Card != %s\n", noPair.Face)
		testsPass = false
	}
	if pair.Value != Pair {
		fmt.Printf("ERROR: Pair = %s\n", pair.Face)
		testsPass = false
	}
	if twoPairs.Value != TwoPairs {
		fmt.Printf("ERROR: Two Pair = %s\n", twoPairs.Face)
		testsPass = false
	}
	if set.Value != Set {
		fmt.Printf("ERROR: Three of a Kind = %s\n", set.Face)
		testsPass = false
	}
	if quads.Value != Quads {
		fmt.Printf("ERROR: Four of a Kind = %s\n", quads.Face)
		testsPass = false
	}
	if fullHouse.Value != FullHouse {
		fmt.Printf("ERROR: Full House = %s\n", fullHouse.Face)
		testsPass = false
	}
	if straight.Value != Straight {
		fmt.Printf("ERROR: Straight = %s\n", straight.Face)
		testsPass = false
	}
	if flush.Value != Flush {
		fmt.Printf("ERROR: Flush = %s\n", flush.Face)
		testsPass = false
	}
	if straightFlush.Value != StraightFlush {
		fmt.Printf("ERROR: Straight Flush = %s\n", straightFlush.Face)
		testsPass = false
	}
	if royalFlush.Value != RoyalFlush {
		fmt.Printf("ERROR: Royal Flush = %s\n", royalFlush.Face)
		testsPass = false
	}
	if aceLowStraight.Value != Straight {
		fmt.Printf("ERROR: Ace Low Straight = %s\n", aceLowStraight.Face)
		testsPass = false
	}
	if tooManyCards != nil {
		fmt.Printf("ERROR: Too Many Cards ==  %t\n", tooManyCards == nil)
		testsPass = false
	}
	if sameCards != nil {
		fmt.Printf("ERROR: Same Cards == %t\n", sameCards == nil)
		testsPass = false
	}

	return testsPass
}

func RunTests() bool {
	testsPass := handValueTestsPass()

	if testsPass {
		fmt.Println("All Tests passed!")
	} else {
		fmt.Println("Test(s) failed!")
	}

	return testsPass
}
