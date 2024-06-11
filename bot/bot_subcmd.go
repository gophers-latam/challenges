package bot

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gophers-latam/challenges/global"
)

func SubCmd(s *discordgo.Session, m *discordgo.MessageCreate) {
	// ignore msg by itself and split the message into arguments
	if m.Author.ID == s.State.User.ID {
		return
	}

	args := strings.Fields(m.Content)
	if len(args) == 0 || args[0] != global.Prefix {
		return
	}

	// go to hello subcommand if only prefix is present
	if len(args) == 1 && args[0] == ".go" {
		msgHello(s, m)
		return
	}

	/*
			// OPTIONAL but NTH
			// ensure there are at least two arguments to process subcommands
		    if len(args) < 2 {
		        // respond with a message indicating incomplete command
		        s.ChannelMessageSend(m.ChannelID, "Incomplete command. Please provide a subcommand.")
		        return
		    }
	*/

	// process subcommands
	switch args[1] {
	case "facts":
		msgFacts(s, m)
	case "events":
		msgEvents(s, m)
	case "hours":
		msgHours(s, m)
	case "challenge":
		msgChallenges(s, m)
	default:
		// handle more subcommands from the database
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
}

func msgHours(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Split the message into arguments using Fields to handle multiple spaces
	args := strings.Fields(m.Content)

	// Ensure the correct number of arguments
	if len(args) != 4 {
		unsuccessfulMsg(s, m, `Error en subcomando, ver ayuda con: **.go help**`)
		return
	}

	hour := args[2]
	country := args[3]

	// Get the equivalent hours for the given country, handling errors
	msg, err := GetHours(hour, country)
	if err != nil {
		unsuccessfulMsg(s, m, fmt.Sprintf("**Ups, no se puede mostrar equivalencia horaria: %s**", err))
		return
	}

	_, _ = s.ChannelMessageSend(m.ChannelID, msg)
}

func msgChallenges(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := m.Content
	values := strings.Fields(cmd)
	valuesLen := len(values)

	switch valuesLen {
	case 3:
		if values[2] == "help" {
			msgCommands(s, m)
			return
		}
	case 4:
		level := values[2]
		topic := values[3]

		msg, err := GetChallenge(level, topic)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				unsuccessfulMsg(s, m, `**Ups, sin desafíos que coincidan**`)
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
		unsuccessfulMsg(s, m, `Error en subcomando, ver ayuda con: **.go challenge help**`)
	}
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
}
