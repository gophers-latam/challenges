package slashcmd_commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gophers-latam/challenges/bot/helpers"
	"github.com/gophers-latam/challenges/bot/service_http"
)

func SlashHello(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: `Hola **` + i.Member.User.Username + `** ` + helpers.DefaultMsg,
		},
	})
}

func SlashHelp(s *discordgo.Session, i *discordgo.InteractionCreate) {
	msg, _ := service_http.GetCommand(".go help")
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: msg.Text,
		},
	})
}
