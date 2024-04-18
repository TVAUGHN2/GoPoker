package hand

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/swaggest/openapi-go"
	"github.com/swaggest/rest"
	"github.com/swaggest/rest/gorillamux"
	"github.com/swaggest/rest/nethttp"
	"github.com/swaggest/rest/request"
	"github.com/tvaughn2/GoPoker/api/resource/card"
)

type HandValueResult struct {
	Value HandValue `json:"Value"`
	Face  HandFace  `json:"Face"`
}

func NewHandValueHandler() *handValueHandler {
	decoderFactory := request.NewDecoderFactory()
	decoderFactory.ApplyDefaults = true
	decoderFactory.SetDecoderFunc(rest.ParamInPath, gorillamux.PathToURLValues)

	return &handValueHandler{
		dec: decoderFactory.MakeDecoder(http.MethodPost, []card.Card{}, nil),
	}
}

type handValueHandler struct {
	// Automated request decoding is not required to collect OpenAPI schema,
	// but it is good to have to establish a single source of truth and to simplify request reading.
	dec nethttp.RequestDecoder
}

// ServeHTTP implements http.Handler.
func (h *handValueHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var payload []card.Card
	if err := h.dec.Decode(r, &payload, nil); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var cards []*card.Card
	for i := 0; i < len(payload); i++ {
		cards = append(cards, &payload[i])
	}

	fmt.Println("Stringifiy: " + card.StringifyCards(cards))

	hand := NewHand(cards)

	handResult := HandValueResult{hand.Value, hand.Face}

	j, err := json.Marshal(handResult)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	_, _ = w.Write(j)
}

// SetupOpenAPIOperation declares OpenAPI schema for the handler.
func (h *handValueHandler) SetupOpenAPIOperation(oc openapi.OperationContext) error {
	oc.SetTags("HandValue")
	oc.SetSummary("Provides hand value rankings based on card info passed in.")
	oc.SetDescription("This endpoint aggregates request in structured way.")

	oc.AddReqStructure([]card.Card{})
	oc.AddRespStructure(HandValueResult{})
	oc.AddRespStructure(nil, openapi.WithContentType("text/html"), openapi.WithHTTPStatus(http.StatusBadRequest))
	oc.AddRespStructure(nil, openapi.WithContentType("text/html"), openapi.WithHTTPStatus(http.StatusInternalServerError))

	return nil
}
