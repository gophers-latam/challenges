package bot

import (
	"database/sql"
	"regexp"
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
	if matched, err := regexp.MatchString("\\bbot\\b", m.Content); err == nil && matched {
		_, _ = s.ChannelMessageSend(m.ChannelID, defaultMsg)
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

	if args[1] == "facts" { // go to facts subcommands
		msgFacts(s, m)
	} else if args[1] == "challenge" { // go to challenges subcommands
		msgChallenges(s, m)
	} else { // more subcommands in database
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
		if err != sql.ErrNoRows {
			unsuccessfulMsg(s, m, `**Ups, algo anda mal**`)
			return
		}
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

		_, _ = s.ChannelMessageSend(m.ChannelID, msg.ChallengeFmt())
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
