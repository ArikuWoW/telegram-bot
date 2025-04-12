package model

import (
	"testing"

	mock_model "github.com/ArikuWoW/telegram-bot/internal/mocks/messages"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_OnStartCommand_ShouldAnswerWithIntroMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mock_model.NewMockMessageSender(ctrl)
	model := New(sender)

	sender.EXPECT().SendMessage("Не знаю эту команду", int64(123))

	err := model.IncomingMessage(Message{
		Text:   "some text",
		UserID: 123,
	})

	assert.NoError(t, err)
}
