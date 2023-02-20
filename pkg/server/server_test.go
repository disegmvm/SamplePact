package server

import (
	"github.com/pact-foundation/pact-go/dsl"
	"log"

	//
	// some imports
	//
	"github.com/pact-foundation/pact-go/types"
	"testing"
)

func TestServerPact_Verification(t *testing.T) {

	/*router := gin.Default()
	router.GET("/users/:userId", GetUserByID)
	router.Run(":8080")*/

	// initialize PACT DSL
	pact := dsl.Pact{
		Provider: "example-server",
	}

	// verify Contract on server side
	log.Println("[[[[[[[[[[]]]]]]]]]]]]]] test")
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://127.0.0.1:8080",
		PactURLs:        []string{"../client/pactx/example-client-example-server.json"},
	})

	if err != nil {
		t.Log(err)
	}
	log.Println("[[[[[[[[[[]]]]]]]]]]]]]] test PASS")
}
