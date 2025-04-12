package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ArikuWoW/telegram-bot/internal/clients/tg"
	"github.com/ArikuWoW/telegram-bot/internal/config"
	"github.com/ArikuWoW/telegram-bot/internal/model/messages"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		sig := <-sigCh
		log.Printf("signal received: %v, shutting down", sig)
		cancel()
	}()

	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed:", err)
	}

	tgClient, err := tg.New(config)
	if err != nil {
		log.Fatal("tg client init failed:", err)
	}

	msgModel := messages.New(tgClient)

	if err := tgClient.ListenUpdates(ctx, msgModel); err != nil {
		log.Println("bot stopped with error:", err)
	} else {
		log.Println("boy stopped gracefully")
	}
}
