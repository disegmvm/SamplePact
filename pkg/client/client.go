package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"transPact/pkg/server"
)

func GetUserByID(host string, id string) (*server.User2, error) {
	uri := fmt.Sprintf("http://%s/users/%s", host, id)
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var user server.User2
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
