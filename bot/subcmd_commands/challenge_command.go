package subcmd_commands

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gophers-latam/challenges/bot/helpers"
	"github.com/gophers-latam/challenges/bot/service_http"
)

// ChallengeCommand structure
type ChallengeCommand struct{}

// Execute method for ChallengeCommand
func (h *ChallengeCommand) Execute(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := m.Content
	values := strings.Fields(cmd)
	valuesLen := len(values)

	switch valuesLen {
	case 4:
		level := values[2]
		topic := values[3]

		msg, err := service_http.GetChallenge(level, topic)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				helpers.UnsuccessfulMsg(s, m, `**Ups, sin desafíos que coincidan**`)
			}
			return
		}

		if msg.Description == "" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Ups, desafío sin **Descripción**")
			return
		}

		_, _ = s.ChannelMessageSend(m.ChannelID, msg.ChallengeFmt())
		return
	default:
		helpers.UnsuccessfulMsg(s, m, `Error en subcomando, ver ayuda con: **.go challenge help**`)
	}
}

// Help method for ChallengeCommand
func (h *ChallengeCommand) Help(cmd string) string {
	msg, _ := service_http.GetCommand(cmd + " help")
	return "Uso: " + msg.Text
}
