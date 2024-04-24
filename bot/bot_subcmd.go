package bot

import (
	"database/sql"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gophers-latam/challenges/global"
)

func SubCmd(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore msg by itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// bot word mention
	words := strings.Fields(m.Content)

	for _, word := range words {
		if word == "bot" {
			_, _ = s.ChannelMessageSend(m.ChannelID, `envia: .go (para usar el gopherbottttt)`)
			return
		}
	}

	// stop if not use subcommand prefix
	args := strings.Split(m.Content, " ")
	if args[0] != global.Prefix {
		return
	}

	// go to hello subcommand
	if len(args) == 1 && args[0] == ".go" {
		msgHello(s, m)
		return
	}

	// go to challenges subcommands
	if args[1] == "challenge" {
		msgChallenges(s, m)
		// more subcommands in database
	} else {
		msgCommands(s, m)
	}
}

func msgHello(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, _ = s.ChannelMessageSend(m.ChannelID,
		`Hola **`+m.Author.Username+`**`)
}

func msgChallenges(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := m.Content

	values := strings.Split(cmd, " ")
	l := len(values)

	if l == 3 && values[2] == "help" {
		msgCommands(s, m)
		return
	}

	if l == 4 {
		level := values[2]
		topic := values[3]

		msg, err := GetChallenge(level, topic)
		if err != nil {
			if err != sql.ErrNoRows {
				unsuccessfulMsg(s, m, `**Ups, sin desafios que coincidan**`)
				return
			}
		}

		if msg.Description == "" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "no challenge with **Description** found")
			return
		}

		message := `[*challenge*]⤵️
		-**Level:** ` + string(msg.Level) + ` -**Type:** ` + string(msg.ChallengeType) + `
		-**Description: ** ` + msg.Description

		_, _ = s.ChannelMessageSend(m.ChannelID, message)
		return
	}

	unsuccessfulMsg(s, m, `Error en subcomando, ver ayuda con: **.go challenge help**`)
}

func msgCommands(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := m.Content

	msg, err := GetCommand(cmd)
	if err != nil {
		if err != sql.ErrNoRows {
			unsuccessfulMsg(s, m, `**Ups, intenta de nuevo, sin espacios extras**`)
			return
		}
	}

	if cmd == global.Prefix+" help" {
		msgEmbed(s, m, cmd, msg.Text)
	} else {
		_, _ = s.ChannelMessageSend(m.ChannelID, msg.Text)
	}

	return
}
