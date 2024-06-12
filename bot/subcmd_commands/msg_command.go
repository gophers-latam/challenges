package subcmd_commands

import (
	"database/sql"
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/gophers-latam/challenges/bot/helpers"
	"github.com/gophers-latam/challenges/bot/service_http"
	"github.com/gophers-latam/challenges/global"
)

func MsgCommands(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := m.Content

	msg, err := service_http.GetCommand(cmd)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.UnsuccessfulMsg(s, m, `**Ups, intenta de nuevo sin espacios extras**`)
		}
		return
	}

	if cmd == global.Prefix+" help" {
		embed := &discordgo.MessageEmbed{
			Title:       cmd,
			Description: msg.Text,
		}
		helpers.MsgEmbed(s, m, embed)
	} else {
		_, _ = s.ChannelMessageSend(m.ChannelID, msg.Text)
	}
}
