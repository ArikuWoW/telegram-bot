package messages

import (
	"strconv"
	"strings"
)

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

type Expense struct {
	Amount int64
	Group  string
}

var Expenses = []Expense{}

func (s *Model) IncomingMessage(msg Message) error {
	parts := strings.Fields(msg.Text)
	if len(parts) == 0 {
		return s.tgClient.SendMessage("Пустое сообщение", msg.UserID)
	}

	cmd := parts[0]

	switch cmd {
	case "/start":
		return s.tgClient.SendMessage("Привет", msg.UserID)
	case "/add":
		if len(parts) < 3 {
			return s.tgClient.SendMessage("Формат добавления расхода: /add <сумма> <товар>", msg.UserID)
		}

		amountStr := parts[1]
		category := parts[2]
		amount, err := strconv.Atoi(amountStr)
		if err != nil {
			return s.tgClient.SendMessage("Сумма должна быть числом", msg.UserID)
		}

		Expenses = append(Expenses, Expense{
			Amount: int64(amount),
			Group:  category,
		})

		return s.tgClient.SendMessage("Расход добавлен", msg.UserID)
	default:
		return s.tgClient.SendMessage("Не знаю эту команду", msg.UserID)
	}

}
