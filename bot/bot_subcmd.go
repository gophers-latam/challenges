package bot

import (
	"database/sql"
	"strings"
	"math/rand"
	"regexp"

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
		_, _ = s.ChannelMessageSend(m.ChannelID, `envia: .go (para usar el gopherbot)`)
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

	if args[1] == "facts" {
		facts := []string{
			"1. Los miembros de Gophers LATAM pueden escribir una api Go en microsoft paint y compilarla en excel escrita en Ruby",
			"2. Cuando el compilador encuentra un error en el código de Gophers LATAM, el compilador se disculpa",
			"3. Gophers LATAM puede dividir por cero, pero el compilador asustado intenta multiplicar por el infinito",
			"4. Gophers LATAM puede arrojar una excepción mas lejos que nadie y en menor tiempo",
			"5. Cuando Gophers LATAM presiona [ctr-alt-del] es el resto mundo el que se reinicia",
			"6. Gophers LATAM no necesita recolector de basura. Solo mira a los objetos fijamente y los objetos se destruyen a si mismos muertos de miedo",
			"7. Gophers LATAM no necesita compilar su código Go. Le basta escribir en Javascript y se traduce el mismo a binario",
			"8. Los miembros de Gophers LATAM no tienen tecla [CONTROL] en su teclado, ellos siempre están en control",
			"9. Gophers LATAM puede detectar el siguiente número en una secuencia aleatoria",
			"10. Gophers LATAM puede ejecutar un loop infinito en 3 segundos",
			"11. Gophers LATAM no puede producir un Null Pointer Exception, si Gophers LATAM apunta a Null, un objeto se materializa instantaneamente",
			"12. Gophers LATAM puede hacer control-z con lápiz y papel",
			"14. Gophers LATAM te hace updates a la base de datos con el buscaminas",
			"15. Los arrays de Gophers LATAM son de tamaño infinito porque Gophers LATAM no tiene límites.",
			"16. Gophers LATAM terminó World of Warcraft",
			"17. Solo hay 10 clases de personas, los que son parte de Gophers LATAM y los que no",
		}

		selection := rand.Intn(len(facts))

		author := discordgo.MessageEmbedAuthor{
			Name: "El Programador Pobre",
		}			

		embed := discordgo.MessageEmbed{
			Title: facts[selection],
			Author: &author,
		}

		s.ChannelMessageSendEmbed(m.ChannelID, &embed)
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
