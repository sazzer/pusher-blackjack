package chatter

import (
	"context"

	"github.com/pusher/chatkit-server-go"
	"github.com/pusher/pusher-platform-go/auth"
)

const CROUPIER = "3893095E-0365-4EC4-993C-7436D57246D1"

type Chatter struct {
	client *chatkit.Client
}

func New() Chatter {
	client, _ := chatkit.NewClient(
		"CHATKIT_INSTANCE_LOCATOR",
		"CHATKIT_KEY",
	)
	return Chatter{client}
}

func (chatter *Chatter) Authenticate(user string) (*auth.Response, error) {
	chatter.client.CreateUser(context.Background(), chatkit.CreateUserOptions{
		ID:   user,
		Name: user,
	})

	return chatter.client.Authenticate(chatkit.AuthenticatePayload{
		GrantType: "client_credentials",
	}, chatkit.AuthenticateOptions{
		UserID: &user,
	})
}

func (chatter *Chatter) Message(table string, message string) {
	chatter.client.SendMessage(context.Background(), chatkit.SendMessageOptions{
		RoomID:   table,
		SenderID: CROUPIER,
		Text:     message,
	})
}
