package subcmd_commands

import (
	"github.com/bwmarrin/discordgo"
)

// HelloCommand structure
type HelloCommand struct{}

// Execute method for HelloCommand
func (h *HelloCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, _ = s.ChannelMessageSend(m.ChannelID, `Hola **`+m.Author.Username+`**`)
}

// Help method for HelloCommand
func (h *HelloCommand) Help() string {
	return "Uso: **.go** - Saludo del bot"
}
