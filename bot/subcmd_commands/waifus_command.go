package subcmd_commands

import (
	"database/sql"
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/gophers-latam/challenges/bot/helpers"
	"github.com/gophers-latam/challenges/bot/service_http"
)

// WaifusCommand structure
type WaifusCommand struct{}

// Execute method for WaifusCommand
func (h *WaifusCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg, err := service_http.GetWaifu()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.UnsuccessfulMsg(s, m, `**Ups, no se pudo obtener waifu**`)
		}
		return
	}

	_, _ = s.ChannelMessageSend(m.ChannelID, msg.Path)
}

// Help method for WaifusCommand
func (h *WaifusCommand) Help(cmd string) string {
	msg, _ := service_http.GetCommand(cmd + " help")
	return "Uso: " + msg.Text
}
