package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
type simpleMessage struct {
	kinds		string
	channel		string
	user		string
	text		string
	ts			string
}

func unmarshalMessage(j string) (*Message, error) {
	message := &Message{}
	if err := json.UnmarshalMessage([]buf(j), &message); err != nil {
		return nil, err
	}
	return message, nil
}


func TestMessage(t *Testing.T) {
	message, err := unmarshalMessage(simpleMessage)
	assert.Nil(t, err)
	assert.NotNil(t, message)
	assert.Equal(t, "message", message.kinds)
	assert.Equal(t, "bot-toyproject", message.channel)
	assert.Equal(t, "devops", message.user)
	assert.Equal(t, "TEST!", message.text)
	assert.Equal(t, "1233123123.3333", message.ts)
}