package subcmd_commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gophers-latam/challenges/bot/helpers"
	"github.com/gophers-latam/challenges/bot/service_http"
)

// HoursCommand structure
type HoursCommand struct{}

// Execute method for HoursCommand
func (h *HoursCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Split the message into arguments using Fields to handle multiple spaces
	args := strings.Fields(m.Content)

	// Ensure the correct number of arguments
	if len(args) != 4 {
		helpers.UnsuccessfulMsg(s, m, `Error en subcomando, ver ayuda con: **.go help**`)
		return
	}

	hour := args[2]
	country := args[3]

	// Get the equivalent hours for the given country, handling errors
	msg, err := service_http.GetHours(hour, country)
	if err != nil {
		helpers.UnsuccessfulMsg(s, m, fmt.Sprintf("**Ups, no se puede mostrar equivalencia horaria: %s**", err))
		return
	}

	_, _ = s.ChannelMessageSend(m.ChannelID, msg)
}

// Help method for HoursCommand
func (h *HoursCommand) Help() string {
	return "Uso: .go hours {HH:MM (formato 24h)} {país/emoji de bandera}\n" +
		"Ejemplo: .go hours 14:30 argentina\n" +
		"Ejemplo: .go hours 14:30 🇦🇷"
}
