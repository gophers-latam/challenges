package subcmd_commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gophers-latam/challenges/bot/service_http"
)

// HelloCommand structure
type HelloCommand struct{}

// Execute method for HelloCommand
func (h *HelloCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, _ = s.ChannelMessageSend(m.ChannelID, `Hola **`+m.Author.Username+`**`)
}

// Help method for HelloCommand
func (h *HelloCommand) Help(cmd string) string {
	msg, _ := service_http.GetCommand(cmd + " help")
	return msg.Text
}
