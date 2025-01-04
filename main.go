package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/gophers-latam/challenges/bot"
	"github.com/gophers-latam/challenges/global"
	chg "github.com/gophers-latam/challenges/http"
	"github.com/gophers-latam/challenges/storage"
	"go.uber.org/zap"
)

func main() {
	config := global.GetConfig()
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatal("session error:", err.Error())
	}

	// Register SubCmds
	bot.RegisterSubCmds()

	// bot handlers
	dg.AddHandler(bot.Stat)
	dg.AddHandler(bot.HandleSubCmd)
	dg.AddHandler(bot.HandleSlhCmd)
	dg.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	if err = dg.Open(); err != nil {
		log.Fatal("bot error:", err.Error())
	}

	cmd := bot.RegisterSlhCmds(dg)

	defer dg.Close()
	// RemoveSlhCmd: please in development environment turn off before last run
	defer bot.RemoveSlhCmd(dg, cmd)

	// web handlers
	storage.Migrate()
	wa := chg.WebApp{DB: storage.Get(), Port: config.Port}
	go func() {
		zap.L().Fatal("web error:",
			zap.Error(wa.App().Listen(":"+wa.Port)),
		)
	}()

	// wait for exit signal to make defer funcs
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit
	if err := wa.App().Shutdown(); err != nil {
		log.Fatal("forced to shutdown:", err)
	}
}
