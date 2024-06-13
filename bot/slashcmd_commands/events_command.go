package slashcmd_commands

import (
	"database/sql"
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/gophers-latam/challenges/bot/helpers"
	"github.com/gophers-latam/challenges/bot/service_http"
)

func SlashEvents(s *discordgo.Session, i *discordgo.InteractionCreate) {
	msg, err := service_http.GetEvents()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.UnsuccessfulInteraction(s, i, `**Ups, sin eventos para mostrar**`)
		}
		return
	}

	for _, e := range *msg {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: e.Text,
			},
		})
	}
}
