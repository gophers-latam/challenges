package bot

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gophers-latam/challenges/global"
)

func SubCmd(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore msg by itself
	if m.Author.ID == s.State.User.ID {
		return
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

	switch {
	case args[1] == "facts":
		msgFacts(s, m)
	case args[1] == "events":
		msgEvents(s, m)
	case args[1] == "hours":
		msgHours(s, m)
	case args[1] == "challenge":
		msgChallenges(s, m)
	default:
		// more subcommands in database
		msgCommands(s, m)
	}
}

func msgHello(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, _ = s.ChannelMessageSend(m.ChannelID,
		`Hola **`+m.Author.Username+`**`)
}

func msgFacts(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg, err := GetFact()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			unsuccessfulMsg(s, m, `**Ups, sin hechos que mencionar**`)
		}
		return
	}

	embed := &discordgo.MessageEmbed{
		Title: msg.Text,
		Author: &discordgo.MessageEmbedAuthor{
			Name: "El Programador Pobre",
		},
	}
	msgEmbed(s, m, embed)

	return
}

func msgEvents(s *discordgo.Session, m *discordgo.MessageCreate) {
	msg, err := GetEvents()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			unsuccessfulMsg(s, m, `**Ups, sin eventos para mostrar**`)
		}
		return
	}

	for _, e := range *msg {
		_, _ = s.ChannelMessageSend(m.ChannelID, e.Text)
	}

	return
}

func msgHours(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := m.Content

	values := strings.Split(cmd, " ")
	l := len(values)

	if l == 4 {
		hour := values[2]
		country := values[3]

		msg := GetHours(hour, country)
		if msg == "" {
			unsuccessfulMsg(s, m, `**Ups, no se puede mostrar equivalencia horaria**`)
			return
		}

		_, _ = s.ChannelMessageSend(m.ChannelID, msg)
		return
	}

	unsuccessfulMsg(s, m, `Error en subcomando, ver ayuda con: **.go help**`)
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
			if errors.Is(err, sql.ErrNoRows) {
				unsuccessfulMsg(s, m, `**Ups, sin desafios que coincidan**`)
			}
			return
		}

		if msg.Description == "" {
			_, _ = s.ChannelMessageSend(m.ChannelID, "Ups, desafío sin **Descripción**")
			return
		}

		_, _ = s.ChannelMessageSend(m.ChannelID, msg.ChallengeFmt())
		return
	}

	unsuccessfulMsg(s, m, `Error en subcomando, ver ayuda con: **.go challenge help**`)
}

func msgCommands(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := m.Content

	msg, err := GetCommand(cmd)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			unsuccessfulMsg(s, m, `**Ups, intenta de nuevo sin espacios extras**`)
		}
		return
	}

	if cmd == global.Prefix+" help" {
		embed := &discordgo.MessageEmbed{
			Title:       cmd,
			Description: msg.Text,
		}
		msgEmbed(s, m, embed)
	} else {
		_, _ = s.ChannelMessageSend(m.ChannelID, msg.Text)
	}

	return
}
