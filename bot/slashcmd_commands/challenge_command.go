package slashcmd_commands

import (
	"database/sql"
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/gophers-latam/challenges/bot/helpers"
	"github.com/gophers-latam/challenges/bot/service_http"
)

func SlashChallenge(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// options in the order provided by the user.
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	var level, challengeType string

	if option, ok := optionMap["level"]; ok {
		level = option.StringValue()
	}

	if option, ok := optionMap["type"]; ok {
		challengeType = option.StringValue()
	}

	msg, err := service_http.GetChallenge(level, challengeType)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.UnsuccessfulInteraction(s, i, `**Ups, sin desafios que coincidan**`)
		}
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg.ChallengeFmt(),
		},
	})
}

func SlashChallengeHelp(s *discordgo.Session, i *discordgo.InteractionCreate) {
	msg, _ := service_http.GetCommand(".go challenge help")
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg.Text,
		},
	})
}
