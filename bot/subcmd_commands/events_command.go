package subcmd_commands

import (
	"database/sql"
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/gophers-latam/challenges/bot/helpers"
	"github.com/gophers-latam/challenges/bot/service_http"
)

// EventsCommand structure
type EventsCommand struct{}

// Execute method for EventsCommand
func (h *EventsCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg, err := service_http.GetEvents()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			helpers.UnsuccessfulMsg(s, m, `**Ups, sin eventos para mostrar**`)
		}
		return
	}

	for _, e := range *msg {
		_, _ = s.ChannelMessageSend(m.ChannelID, e.Text)
	}
}

// Help method for EventsCommand
func (e *EventsCommand) Help() string {
	return "Uso: .go events - Listar eventos establecidos como recurrentes"
}
