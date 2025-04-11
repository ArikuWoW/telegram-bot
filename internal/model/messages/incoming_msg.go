package model

type MessageSender interface {
	SendMessage(text string, userID int64) error
}

type Model struct {
	tgClient MessageSender
}

func New(tgClient MessageSender) *Model {
	return &Model{tgClient: tgClient}
}

type Message struct {
	Text   string
	UserID int64
}

func (m *Model) IncomingMessage(msg Message) error {
	return nil
}
