package main

import (
	"github.com/bwmarrin/discordgo"
	messages "github.com/tomiok/challenge-svc/discordbot"
	"github.com/tomiok/challenge-svc/global"
	"github.com/tomiok/challenge-svc/logs"
	"github.com/tomiok/challenge-svc/storage"
	"go.uber.org/zap"
	"log"
)

func main() {
	logs.InitLogs()

	dg, err := discordgo.New("Bot " + global.Token)
	if err != nil {
		log.Fatal("session error:", err.Error())
	}
	// dg.Debug = true

	// Register bot handlers.
	dg.AddHandler(messages.MessageCmd)
	dg.AddHandler(messages.SetStatus)

	err = dg.Open()
	if err != nil {
		log.Fatal("websocket error,", err.Error())
	}

	storage.Migrate()
	port := global.Port
	if port == "" {
		port = "5000"
	}

	zap.L().Fatal("cannot init server", zap.Error(start(port)))
}
