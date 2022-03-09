package messages

import (
	"database/sql"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func MessageCmd(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore msg by itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	log.Println(m.Content)
	if strings.Contains(m.Content, ".go challenge") {
		printMessages(s, m)
	}
}

func printHelpMessage(s *discordgo.Session, m *discordgo.MessageCreate, strs []string) {
	if strs[2] == "help" {

		_, _ = s.ChannelMessageSend(m.ChannelID,
			`Para usar el bot de challenge, debes seguir esta sintaxis:
	** .go challenge {nivel} {tipo_de_chellenge} **
	
	Los niveles disponibles son: easy, medium, hard
	Los tipos disponibles son: algorithm, concurrency, database, web, cli, core, frontend

	Podes subir el texto del challenge aca: https://challenbot.herokuapp.com`)
	}
}

func printUnsuccessfulMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, _ = s.ChannelMessageSend(m.ChannelID, `**No sé, error en comando, prueba .go challenge help**`)
}

func printMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := m.Content

	values := strings.Split(cmd, " ")
	l := len(values)

	if l == 3 {
		printHelpMessage(s, m, values)
		return
	}

	if l == 4 {
		level := values[2]
		topic := values[3]

		msg, err := GetChallenge(level, topic)
		if err != nil {
			if err != sql.ErrNoRows {
				_, _ = s.ChannelMessageSend(m.ChannelID, `**Ups, intenta de nuevo, sin espacios extras**`)
				return
			}
		}

		if msg.Description == "" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "no encontramos ningun challenge")
			return
		}

		message := `[*Challenge*] 
			-**Level:** ` + string(msg.Level) + ` -**Type:** ` + string(msg.ChallengeType) + `
			-**Description:**` + msg.Description

		_, _ = s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	printUnsuccessfulMessage(s, m)
}
