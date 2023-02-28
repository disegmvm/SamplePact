package client

import (
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"net/http"
	"testing"
	"transPact/pkg/server"
)

var commonHeaders = dsl.MapMatcher{
	"Content-Type": dsl.Term("application/json; charset=utf-8", `application\/json`),
}

func Test_RiExistTest(t *testing.T) {

	// Create Pact connecting to local Daemon
	pact := &dsl.Pact{
		Consumer: "PBB_SQS",
		Provider: "SKU_Extract_handler",
		//Host:     "localhost",
	}

	defer pact.Teardown()

	var test = func() (err error) {
		url := fmt.Sprintf("http://localhost:%d/sample/endpoint/to/get/transformed/sku", pact.Server.Port)
		//req, err := http.NewRequest("GET", url, strings.NewReader(`{"name":"River Island"}`))
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return
		}

		// NOTE: by default, request bodies are expected to be sent with a Content-Type
		// of application/json. If you don't explicitly set the content-type, you
		// will get a mismatch during Verification.
		req.Header.Set("Content-Type", "application/json")
		//req.Header.Set("Authorization", "Bearer 1234")

		_, err = http.DefaultClient.Do(req)
		return
	}

	pact.
		AddInteraction().
		Given("RI Exist").
		UponReceiving("A request to retrieve RI").
		WithRequest(dsl.Request{
			Method: "GET",
			//Path:   dsl.Like(server.TransformedSku{}),
			//Path:    dsl.Term("/sample/endpoint/to/get/transformed/sku", "/users/endpoint/to/get/transformed/sku"),
			Path:    dsl.Term("/sample/endpoint/to/get/transformed/sku", "/sample\\/endpoint\\/to\\/get\\/transformed\\/sku"),
			Headers: commonHeaders,
		}).
		WillRespondWith(dsl.Response{
			Status: 200,
			Body: dsl.Match(
				//server.User2{},
				server.TransformedSku{},
			),
			Headers: dsl.MapMatcher{
				//"X-Api-Correlation-Id": dsl.Like("100"),
				"Content-Type": dsl.Term("application/json; charset=utf-8", `application\/json`),
				//"X-Auth-Token":         dsl.Like("1234"),
			},

			/*Body: dsl.Like(server.User2{
				ID:        "1",
				FirstName: "Default first name",
				LastName:  "Default last name",
				Title:     "River Island",
			}),*/
		})

	err := pact.Verify(test)
	if err != nil {
		t.Fatalf("Error on Verify: %v", err)
	}

	// specify PACT publisher
	publisher := dsl.Publisher{}
	err = publisher.Publish(types.PublishRequest{
		PactURLs:        []string{"../client/pacts/PBB_SQS-SKU_Extract_handler.json"},
		PactBroker:      "https://pen.pactflow.io/",
		BrokerToken:     "jEQnxw7xWgYRv-3-G7Cx-g",
		ConsumerVersion: "1.0.0",
		Tags:            []string{"1.0.0", "latest"},
	})
	if err != nil {
		t.Fatal(err)
	}
}

// LoginResponse is the login response API struct.
type Response struct {
	User *server.User2 `json:"user"`
}

/*func TestExampleConsumerLoginHandler_UserDoesNotExist(t *testing.T) {
	var testJmarieDoesNotExists = func() error {
		client := Client{
			Host: fmt.Sprintf("http://localhost:%d", pact.Server.Port),
			// This header will be dynamically replaced at verification time
			token: "Bearer 1234",
		}
		client.loginHandler(rr, req)

		if client.user != nil {
			return fmt.Errorf("Expected user to be nil but in stead got: %v", client.user)
		}

		return nil
	}

	pact.
		AddInteraction().
		Given("User jmarie does not exist").
		UponReceiving("A request to login with user 'jmarie'").
		WithRequest(request{
			Method:  "POST",
			Path:    s("/login/10"),
			Body:    loginRequest,
			Headers: commonHeaders,
			Query: dsl.MapMatcher{
				"foo": s("anything"),
			},
		}).
		WillRespondWith(dsl.Response{
			Status:  404,
			Headers: commonHeaders,
		})

	err := pact.Verify(testJmarieDoesNotExists)
	if err != nil {
		t.Fatalf("Error on Verify: %v", err)
	}
}*/
