package client

import (
	"testing"
	"transPact/pkg/client"
)

func TestClientPact_Local(t *testing.T) {

	user, err := client.GetUserByID("localhost:8080", "100")
	if user == nil {
		println("user is NIL")
	}
	if err != nil {
		panic(err)
	}

}
