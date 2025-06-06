package tg

import (
	"context"
	"fmt"
	"log"

	"github.com/ArikuWoW/telegram-bot/internal/model/messages"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TokenGetter interface {
	Token() string
}

type Client struct {
	client *tgbotapi.BotAPI
}

func New(tokenGetter TokenGetter) (*Client, error) {
	client, err := tgbotapi.NewBotAPI(tokenGetter.Token())
	if err != nil {
		return nil, fmt.Errorf("NewBotAPI: %w", err)
	}
	return &Client{client: client}, nil
}

func (c *Client) SendMessage(text string, userID int64) error {
	_, err := c.client.Send(tgbotapi.NewMessage(userID, text))
	if err != nil {
		return fmt.Errorf("client.Send: %w", err)
	}
	return nil
}

func (c *Client) ListenUpdates(ctx context.Context, msgModel *messages.Model) error {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := c.client.GetUpdatesChan(u)
	log.Println("listening for messages")

	for {
		select {
		case <-ctx.Done():
			log.Println("stopping Telegram updates")
			return ctx.Err()
		case update := <-updates:
			if update.Message != nil {
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

				err := msgModel.IncomingMessage(messages.Message{
					Text:   update.Message.Text,
					UserID: update.Message.From.ID,
				})
				if err != nil {
					log.Println("error processing message:", err)
				}
			}
		}
	}
}
