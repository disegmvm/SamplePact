package client

import (
	"errors"
	"fmt"
	"github.com/pact-foundation/pact-go/types"
	"log"
	"testing"
	"transPact/pkg/server"

	"github.com/pact-foundation/pact-go/dsl"
)

func TestClientPact_Local(t *testing.T) {
	log.Println("[debug] test start")

	pact := dsl.Pact{
		Consumer: "Client",
		Provider: "Provider",
	}

	pact.Setup(true)

	// Validate that there is a user with such ID
	t.Run("get user by id", func(t *testing.T) {
		pact.
			AddInteraction().
			Given("User River Island exists").
			UponReceiving("User 'River Island' is requested").
			WithRequest(dsl.Request{
				Method: "GET",
				Path:   dsl.Term("/users/100", "/users/[0-9]+"),
			}).
			WillRespondWith(dsl.Response{
				Status: 200,
				Body: dsl.Like(server.User{
					ID:        "100",
					FirstName: "River Island",
					LastName:  "UK",
				}),
			})

		// Validate the current test case by running it against a Mock Service.
		err := pact.Verify(func() error {
			host := fmt.Sprintf("%s:%d", pact.Host, pact.Server.Port)

			user, err := GetUserByID(host, "100")
			if err != nil {
				return errors.New("error is not expected")
			}

			if user == nil || user.ID != "100" {
				return fmt.Errorf("expected user with ID %s but got %v", "100", user)
			}

			// Add interaction for below case
			/*if user == nil || user.FirstName != "River Island" {
				return fmt.Errorf("expected to have River Island, but got %s", user)
			}*/

			return err
		})

		if err != nil {
			t.Fatal(err)
		}

		// specify PACT publisher
		publisher := dsl.Publisher{}
		err = publisher.Publish(types.PublishRequest{
			PactURLs:        []string{"../client/pacts/Client-Provider.json"},
			PactBroker:      "https://pen.pactflow.io/",
			BrokerToken:     "jEQnxw7xWgYRv-3-G7Cx-g",
			ConsumerVersion: "1.0.0",
			Tags:            []string{"1.0.0", "latest"},
		})
		if err != nil {
			t.Fatal(err)
		}
	})

	// enable if no PactFlow is used
	/*if err := pact.WritePact(); err != nil {
		t.Fatal(err)
	}*/

	pact.Teardown()

	log.Println("[debug] end with no errors")
}
