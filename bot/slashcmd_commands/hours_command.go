package slashcmd_commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/gophers-latam/challenges/bot/helpers"
	"github.com/gophers-latam/challenges/bot/service_http"
)

func SlashHours(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	var hour, country string

	if option, ok := optionMap["hour"]; ok {
		hour = option.StringValue()
	}

	if option, ok := optionMap["country"]; ok {
		country = option.StringValue()
	}

	if hour == "" || country == "" {
		helpers.UnsuccessfulInteraction(s, i, `**Ups, opciones de comando faltantes**`)
		return
	}

	msg, err := service_http.GetHours(hour, country)
	if err != nil {
		helpers.UnsuccessfulInteraction(s, i, fmt.Sprintf("**Ups, no se puede mostrar equivalencia horaria: %s**", err))
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg,
		},
	})
}
