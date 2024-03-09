package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/miem-refugees/wish-wish/internal/telegram"
	"github.com/miem-refugees/wish-wish/internal/telegram/handler"
	"github.com/miem-refugees/wish-wish/internal/telegram/handler/middleware"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	bot, err := telegram.NewBot(
		os.Getenv("TOKEN"),
		logger,
	)
	if err != nil {
		log.Fatal(err)
	}

	bot.RegisterHandler(
		handler.CommonHandlers(logger)...,
	)
	bot.Use(
		middleware.Logger(logger),
	)

	logger.Info("Starting bot")
	go bot.Start()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	logger.Info("Got interrupt signal, stopping bot")
	bot.Stop()
	logger.Info("Bot stopped")
}
