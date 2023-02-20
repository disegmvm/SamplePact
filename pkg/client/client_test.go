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
	log.Println("OOOOOOOOOOOOOOOOOOOOOOOO nachalo OOOOOOOOOOOOOOOOOOOO")

	// инициализация PACT DSL
	pact := dsl.Pact{
		Consumer: "klient_ebanii",
		Provider: "provaider_ebanii",
	}

	// установка PACT Mock Server
	pact.Setup(true)

	t.Run("get user by id", func(t *testing.T) {
		id := "1"

		pact.
			AddInteraction().                           // задается PACT-взаимодействие
			Given("User Alice exists").                 // задается состояние Поставщика
			UponReceiving("User 'Alice' is requested"). // задается название тест-кейса
			WithRequest(dsl.Request{                    // задается ожидаемый запрос
				Method: "GET",
				Path:   dsl.Term("/users/1", "/users/[0-9]+"), // задается соответствие для конечной точки
			}).
			WillRespondWith(dsl.Response{ // задается минимальный ожидаемый ответ
				Status: 200,
				Body: dsl.Like(server.User{ // задается соответствие для тела ответа
					ID:        id,
					FirstName: "Alice",
					LastName:  "Doe",
				}),
			})

		// верификация взаимодействия на стороне клиента
		err := pact.Verify(func() error {
			// задается хост и публикуется PACT Mock Server в качестве актуального сервера
			host := fmt.Sprintf("%s:%d", pact.Host, pact.Server.Port)

			// выполнение функции
			user, err := GetUserByID(host, id)
			if err != nil {
				return errors.New("error is not expected")
			}

			// проверка, что полученный пользователь соответствует ожидаемому
			if user == nil || user.ID != id {
				return fmt.Errorf("expected user with ID %s but got %v", id, user)
			}

			return err
		})

		if err != nil {
			t.Fatal(err)
		}
	})

	// Контракт пишется в файл
	if err := pact.WritePact(); err != nil {
		t.Fatal(err)
	}

	// остановка PACT Mock Server
	pact.Teardown()

	log.Println("OOOOOOOOOOOOOOOOOOOOOOOO konec OOOOOOOOOOOOOOOOOOOO")
}
