package hand

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tvaughn2/GoPoker/api/resource/card"
)

type HandValueResult struct {
	Value HandValue `json:"Value"`
	Face  HandFace  `json:"Face"`
}

func HandValueProcessor(w http.ResponseWriter, r *http.Request) {
	var payload []card.Card
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&payload); err != nil {
		json.NewEncoder(w).Encode("Unable to decode!")
		//respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	var cards []*card.Card
	for i := 0; i < len(payload); i++ {
		cards = append(cards, &payload[i])
	}

	fmt.Println("Stringifiy: " + card.StringifyCards(cards))

	hand := NewHand(cards)

	// if err := hand; err != nil {
	// 	json.NewEncoder(w).Encode("Unable to create hand!")
	//     //respondWithError(w, http.StatusInternalServerError, err.Error())
	//     return
	// }

	handResult := HandValueResult{hand.Value, hand.Face}
	json.NewEncoder(w).Encode(handResult)

}
