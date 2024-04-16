package hand_test

import (
	"fmt"

	"github.com/tvaughn2/GoPoker/api/resource/card"
	"github.com/tvaughn2/GoPoker/api/resource/hand"
)

// *** Test Methods ***//
func noPair() *hand.Hand {
	return hand.NewHand([]*card.Card{card.NewCard(card.Cv2, card.Hearts), card.NewCard(card.CvJ, card.Spades),
		card.NewCard(card.CvQ, card.Spades), card.NewCard(card.CvK, card.Spades), card.NewCard(card.CvA, card.Spades)})
}
func pair() *hand.Hand {
	return hand.NewHand([]*card.Card{card.NewCard(card.Cv2, card.Hearts), card.NewCard(card.Cv2, card.Spades),
		card.NewCard(card.CvQ, card.Spades), card.NewCard(card.CvK, card.Spades), card.NewCard(card.CvA, card.Spades)})
}
func twoPairs() *hand.Hand {
	return hand.NewHand([]*card.Card{card.NewCard(card.Cv2, card.Hearts), card.NewCard(card.Cv2, card.Spades),
		card.NewCard(card.CvQ, card.Spades), card.NewCard(card.CvQ, card.Clubs), card.NewCard(card.CvA, card.Spades)})
}
func set() *hand.Hand {
	return hand.NewHand([]*card.Card{card.NewCard(card.Cv2, card.Hearts), card.NewCard(card.Cv2, card.Spades),
		card.NewCard(card.CvQ, card.Spades), card.NewCard(card.Cv2, card.Clubs), card.NewCard(card.CvA, card.Spades)})
}
func quads() *hand.Hand {
	return hand.NewHand([]*card.Card{card.NewCard(card.Cv2, card.Hearts), card.NewCard(card.Cv2, card.Spades),
		card.NewCard(card.Cv2, card.Diamonds), card.NewCard(card.Cv2, card.Clubs), card.NewCard(card.CvA, card.Spades)})
}
func fullHouse() *hand.Hand {
	return hand.NewHand([]*card.Card{card.NewCard(card.Cv2, card.Hearts), card.NewCard(card.Cv2, card.Spades),
		card.NewCard(card.Cv2, card.Diamonds), card.NewCard(card.CvA, card.Clubs), card.NewCard(card.CvA, card.Spades)})
}
func straight() *hand.Hand {
	return hand.NewHand([]*card.Card{card.NewCard(card.Cv3, card.Spades), card.NewCard(card.Cv2, card.Hearts),
		card.NewCard(card.Cv4, card.Spades), card.NewCard(card.Cv5, card.Clubs), card.NewCard(card.Cv6, card.Spades)})
}
func aceLowStraight() *hand.Hand {
	return hand.NewHand([]*card.Card{card.NewCard(card.CvA, card.Hearts), card.NewCard(card.Cv2, card.Spades),
		card.NewCard(card.Cv3, card.Spades), card.NewCard(card.Cv4, card.Clubs), card.NewCard(card.Cv5, card.Spades)})
}
func flush() *hand.Hand {
	return hand.NewHand([]*card.Card{card.NewCard(card.Cv2, card.Hearts), card.NewCard(card.Cv4, card.Hearts),
		card.NewCard(card.CvQ, card.Hearts), card.NewCard(card.CvK, card.Hearts), card.NewCard(card.CvA, card.Hearts)})
}
func straightFlush() *hand.Hand {
	return hand.NewHand([]*card.Card{card.NewCard(card.Cv2, card.Hearts), card.NewCard(card.Cv3, card.Hearts),
		card.NewCard(card.Cv4, card.Hearts), card.NewCard(card.Cv5, card.Hearts), card.NewCard(card.Cv6, card.Hearts)})
}
func royalFlush() *hand.Hand {
	return hand.NewHand([]*card.Card{card.NewCard(card.CvT, card.Hearts), card.NewCard(card.CvJ, card.Hearts),
		card.NewCard(card.CvQ, card.Hearts), card.NewCard(card.CvK, card.Hearts), card.NewCard(card.CvA, card.Hearts)})
}
func tooManyCards() *hand.Hand {
	return hand.NewHand([]*card.Card{card.NewCard(card.CvT, card.Hearts), card.NewCard(card.CvJ, card.Hearts),
		card.NewCard(card.CvQ, card.Hearts), card.NewCard(card.CvK, card.Hearts), card.NewCard(card.CvA, card.Hearts), card.NewCard(card.Cv2, card.Hearts)})
}
func sameCards() *hand.Hand {
	return hand.NewHand([]*card.Card{card.NewCard(card.CvT, card.Hearts), card.NewCard(card.CvJ, card.Hearts),
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

	if noPair.Value != hand.HighCard {
		fmt.Printf("ERROR: High Card != %s\n", noPair.Face)
		testsPass = false
	}
	if pair.Value != hand.Pair {
		fmt.Printf("ERROR: Pair = %s\n", pair.Face)
		testsPass = false
	}
	if twoPairs.Value != hand.TwoPairs {
		fmt.Printf("ERROR: Two Pair = %s\n", twoPairs.Face)
		testsPass = false
	}
	if set.Value != hand.Set {
		fmt.Printf("ERROR: Three of a Kind = %s\n", set.Face)
		testsPass = false
	}
	if quads.Value != hand.Quads {
		fmt.Printf("ERROR: Four of a Kind = %s\n", quads.Face)
		testsPass = false
	}
	if fullHouse.Value != hand.FullHouse {
		fmt.Printf("ERROR: Full House = %s\n", fullHouse.Face)
		testsPass = false
	}
	if straight.Value != hand.Straight {
		fmt.Printf("ERROR: Straight = %s\n", straight.Face)
		testsPass = false
	}
	if flush.Value != hand.Flush {
		fmt.Printf("ERROR: Flush = %s\n", flush.Face)
		testsPass = false
	}
	if straightFlush.Value != hand.StraightFlush {
		fmt.Printf("ERROR: Straight Flush = %s\n", straightFlush.Face)
		testsPass = false
	}
	if royalFlush.Value != hand.RoyalFlush {
		fmt.Printf("ERROR: Royal Flush = %s\n", royalFlush.Face)
		testsPass = false
	}
	if aceLowStraight.Value != hand.Straight {
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
