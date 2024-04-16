package hand_test

import (
	"testing"

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

func TestHighCard(t *testing.T) {
	noPair := noPair()
	if noPair.Value != hand.HighCard {
		t.Errorf("ERROR: High Card != %s\n", noPair.Face)
	}
}

func TestPair(t *testing.T) {
	pair := pair()
	if pair.Value != hand.Pair {
		t.Errorf("ERROR: Pair = %s\n", pair.Face)
	}
}

func TestTwoPairs(t *testing.T) {
	twoPairs := twoPairs()
	if twoPairs.Value != hand.TwoPairs {
		t.Errorf("ERROR: Two Pair = %s\n", twoPairs.Face)
	}
}

func TestSet(t *testing.T) {
	set := set()
	if set.Value != hand.Set {
		t.Errorf("ERROR: Three of a Kind = %s\n", set.Face)
	}
}

func TestQuads(t *testing.T) {
	quads := quads()
	if quads.Value != hand.Quads {
		t.Errorf("ERROR: Four of a Kind = %s\n", quads.Face)
	}
}

func TestFullHouse(t *testing.T) {
	fullHouse := fullHouse()
	if fullHouse.Value != hand.FullHouse {
		t.Errorf("ERROR: Full House = %s\n", fullHouse.Face)
	}
}

func TestStraight(t *testing.T) {
	straight := straight()
	if straight.Value != hand.Straight {
		t.Errorf("ERROR: Straight = %s\n", straight.Face)
	}
}

func TestFlush(t *testing.T) {
	flush := flush()
	if flush.Value != hand.Flush {
		t.Errorf("ERROR: Flush = %s\n", flush.Face)
	}
}

func TestStraightFlush(t *testing.T) {
	straightFlush := straightFlush()
	if straightFlush.Value != hand.StraightFlush {
		t.Errorf("ERROR: Straight Flush = %s\n", straightFlush.Face)
	}
}

func TestRoyalFlush(t *testing.T) {
	royalFlush := royalFlush()
	if royalFlush.Value != hand.RoyalFlush {
		t.Errorf("ERROR: Royal Flush = %s\n", royalFlush.Face)
	}
}

func TestAceLowStraight(t *testing.T) {
	aceLowStraight := aceLowStraight()
	if aceLowStraight.Value != hand.Straight {
		t.Errorf("ERROR: Ace Low Straight = %s\n", aceLowStraight.Face)
	}
}

func TestTooManyCards(t *testing.T) {
	tooManyCards := tooManyCards()
	if tooManyCards != nil {
		t.Errorf("ERROR: Too Many Cards ==  %t\n", tooManyCards == nil)
	}
}

func TestSameCards(t *testing.T) {
	sameCards := sameCards()
	if sameCards != nil {
		t.Errorf("ERROR: Same Cards == %t\n", sameCards == nil)
	}
}
