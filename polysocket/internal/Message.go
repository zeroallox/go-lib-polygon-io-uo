package internal

import (
	"errors"
	"fmt"
)

type Message struct {
	Ev      string `json:"ev,omitempty"`
	Status  string `json:"status,omitempty"`
	Message string `json:"Message,omitempty"`
	Action  string `json:"action,omitempty"`
	Params  string `json:"params,omitempty"`
}

func MakeAuthMessage(apiKey string) Message {
	return Message{
		Action: "auth",
		Params: apiKey,
	}
}

func ConfigureModSubMessage(msg *Message, action SubAction, topic string, symbols []string) error {

	if action == SAInvalid {
		return errors.New("invalid action")
	}

	if len(topic) == 0 {
		return errors.New("invalid topic")
	}

	msg.Action = action.string()
	msg.Params = generateSubListString(topic, symbols)

	return nil
}

func generateSubListString(topic string, symbols []string) string {

	var str string
	str = str + fmt.Sprintf("%v.%v", topic, symbols[0])

	for _, cSymbol := range symbols[1:] {
		str = str + fmt.Sprintf(",%v.%v", topic, cSymbol)
	}

	return str
}
