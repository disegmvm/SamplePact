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

	log.Println("[debug] test start")

	pact := dsl.Pact{
		Consumer: "PBB_SQS",
		Provider: "SKU_Extract_handler",
	}

	log.Println("[debug] start verification")
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL:            "http://127.0.0.1:8080",
		BrokerURL:                  "https://pen.pactflow.io",
		BrokerToken:                "jEQnxw7xWgYRv-3-G7Cx-g",
		PublishVerificationResults: true,
		ProviderVersion:            "1.0.0",
	})

	if err != nil {
		t.Fatal(err)
	}

	// Uncomment to verify contract locally
	/*log.Println("[debug] start verification")
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: "http://127.0.0.1:8080",
		PactURLs:        []string{"../client/pactsX/den_client-den_provider.json"},
	})*/

	if err != nil {
		t.Log(err)
	}
	log.Println("[debug] passed with no errors")
}
