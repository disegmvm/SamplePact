package client

import (
	"errors"
	"fmt"
	"log"
	"testing"
	"transPact/pkg/server"

	"github.com/pact-foundation/pact-go/dsl"
)

func TestClientPact_Local(t *testing.T) {
	// initialize PACT DSL
	pact := dsl.Pact{
		Consumer:                 "example-client",
		Provider:                 "example-server",
		DisableToolValidityCheck: false,
	}

	// setup a PACT Mock Server
	pact.Setup(true)

	t.Run("get user by id", func(t *testing.T) {
		id := "1"

		log.Println("[XXX DEBUG XXX] i'm in test ...")

		pact.
			AddInteraction().                           // specify PACT interaction
			Given("User Alice exists").                 // specify Provider state
			UponReceiving("User 'Alice' is requested"). // specify test case name
			WithRequest(dsl.Request{                    // specify expected request
				Method: "GET",
				Path:   dsl.Term("/users/1", "/users/[0-9]+"), // specify matching for endpoint
			}).
			WillRespondWith(dsl.Response{ // specify minimal expected response
				Status: 200,
				Body: dsl.Like(server.User{ // pecify matching for response body
					ID:        id,
					FirstName: "Alice",
					LastName:  "Doe",
				}),
			})

		// verify interaction on client side
		err := pact.Verify(func() error {
			// specify host anf post of PACT Mock Server as actual server
			host := fmt.Sprintf("%s:%d", pact.Host, pact.Server.Port)

			// execute function
			user, err := GetUserByID(host, id)
			if err != nil {
				return errors.New("error is not expected")
			}

			// check if actual user is equal to expected
			if user == nil || user.ID != id {
				return fmt.Errorf("expected user with ID %s but got %v", id, user)
			} else {
				log.Println("[XXX DEBUG XXX] ONE TEST PASSED ...")
			}

			return err
		})

		if err != nil {
			t.Fatal(err)
		}

		log.Println("[XXX DEBUG XXX] TESTI PROSHLI &&&??? ...")
	})

	// write Contract into file
	if err := pact.WritePact(); err != nil {
		t.Fatal(err)
	}

	// stop PACT mock server
	pact.Teardown()
}
