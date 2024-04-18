package router

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/rest/gorillamux"
	"github.com/tvaughn2/GoPoker/api/resource/hand"
)

func openApiCheck(err error) bool {
	if err != nil {
		fmt.Println("Unable to generate Open Api json document. Swagger UI may be unavailable or outdated.")
		return false
	}

	return true
}

func HandleRequests() {

	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/handvalue", hand.NewHandValueHandler()).Methods("POST")

	// Setup OpenAPI schema.
	refl := openapi3.NewReflector()
	refl.SpecSchema().SetTitle("GoPoker API")
	refl.SpecSchema().SetVersion("v1.0.0")
	refl.SpecSchema().SetDescription("This is a poker calculation engine written in Go.")

	getMux := router.Methods(http.MethodGet).Subrouter()
	opts := middleware.RedocOpts{SpecURL: "/openapi.json"}
	sh := middleware.Redoc(opts, nil)

	getMux.Handle("/docs", sh)
	getMux.Handle("/openapi.json", http.FileServer(http.Dir("./docs/")))

	refl.Spec.AddOperation(http.MethodGet, "/docs", openapi3.Operation{})
	refl.Spec.AddOperation(http.MethodGet, "/openapi.json", openapi3.Operation{})

	// Walk the router with OpenAPI collector.
	c := gorillamux.NewOpenAPICollector(refl)
	_ = router.Walk(c.Walker)

	// Get the resulting schema.
	json, _ := refl.Spec.MarshalJSON()

	f, err := os.Create("./docs/openapi.json")
	validOpenApiFile := openApiCheck(err)

	if validOpenApiFile {
		_, err = f.WriteString(string(json))
		openApiCheck(err)
	}
	f.Close()

	fmt.Println("Listening on port 3000")

	log.Fatal(http.ListenAndServe(":3000", router))
}
