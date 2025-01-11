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

	// Init type slash InteractionRespond
	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: (*msg)[0].Text,
		},
	})
	if err != nil {
		helpers.UnsuccessfulInteraction(s, i, `**Ups, error enviando evento inicial**`)
		return
	}

	// Send remaining messages with ChannelMessageSend
	for _, e := range (*msg)[1:] {
		_, err := s.ChannelMessageSend(i.ChannelID, e.Text)
		if err != nil {
			_, _ = s.ChannelMessageSend(i.ChannelID, `**Ups, error enviando resto de eventos**`)
		}
	}
}
