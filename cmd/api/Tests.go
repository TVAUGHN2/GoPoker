package main

import (
	"fmt"
)

// *** Test Methods ***//
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
