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

	// initialize PACT DSL
	pact := dsl.Pact{
		Provider: "Denis testit server ebanii",
	}

	// verify Contract on server side
	log.Println("OOOOOOOOOOOOO nachalo vefity OOOOOOOOOOOOOOOOOOO")
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://127.0.0.1:8080",
		PactURLs:        []string{"../client/pacts/klient_ebanii-provaider_ebanii.json"},
	})

	if err != nil {
		t.Log(err)
	}
	log.Println("OOOOOOOOOOOOOOOOOOOOO konec verify (pass???????????) OOOOOOOOOOOOOOOOO")
}
