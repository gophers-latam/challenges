package helpers

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const DefaultMsg = `envia **.go help** para usar el gopherbot...`

func UnsuccessfulMsg(s *discordgo.Session, m *discordgo.MessageCreate, t string) {
	_, _ = s.ChannelMessageSend(m.ChannelID, t)
}

func UnsuccessfulInteraction(s *discordgo.Session, i *discordgo.InteractionCreate, t string) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: t,
		},
	})
}

func MsgEmbed(s *discordgo.Session, m *discordgo.MessageCreate, e *discordgo.MessageEmbed) {
	if e.Author == nil {
		e.URL = "https://dsc.gg/gophers-latam"
	}
	e.Color = 0x78141b
	s.ChannelMessageSendEmbed(m.ChannelID, e)
}

func WordCase(s string) string {
	if len(s) == 0 {
		return s
	}

	return strings.ToUpper(s[:1]) + s[1:]
}

func IntnCrypt(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be greater than 0")
	}

	// rand number between [0, n]
	bigN := big.NewInt(int64(n))
	result, err := rand.Int(rand.Reader, bigN)
	if err != nil {
		return 0, err
	}

	return int(result.Int64()), nil
}
