package subcmd_commands

import (
	"database/sql"
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/gophers-latam/challenges/bot/helpers"
	"github.com/gophers-latam/challenges/bot/service_http"
)

// FactsCommand structure
type FactsCommand struct{}

// Execute method for FactsCommand
func (h *FactsCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg, err := service_http.GetFact()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.UnsuccessfulMsg(s, m, `**Ups, sin hechos que mencionar**`)
		}
		return
	}

	embed := &discordgo.MessageEmbed{
		Title: msg.Text,
		Author: &discordgo.MessageEmbedAuthor{
			Name: "El Programador Pobre",
		},
	}

	helpers.MsgEmbed(s, m, embed)
}

// Help method for FactsCommand
func (h *FactsCommand) Help(cmd string) string {
	msg, _ := service_http.GetCommand(cmd + " help")
	return "Uso: " + msg.Text
}
