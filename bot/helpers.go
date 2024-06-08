package bot

import (
	"log"
	"unicode"

	"github.com/bwmarrin/discordgo"
)

const defaultMsg = `envia **.go help** para usar el gopherbot...`

func Stat(discord *discordgo.Session, ready *discordgo.Ready) {
	err := discord.UpdateStatusComplex(discordgo.UpdateStatusData{Status: "dnd"})
	if err != nil {
		log.Println(err.Error())
	}

	_ = discord.UpdateGameStatus(1, ".go help")
	servers := discord.State.Guilds
	log.Printf("Bot running on %d servers", len(servers))
}

func unsuccessfulMsg(s *discordgo.Session, m *discordgo.MessageCreate, t string) {
	_, _ = s.ChannelMessageSend(m.ChannelID, t)
}

func unsuccessfulInteraction(s *discordgo.Session, i *discordgo.InteractionCreate, t string) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: t,
		},
	})
}

func msgEmbed(s *discordgo.Session, m *discordgo.MessageCreate, e *discordgo.MessageEmbed) {
	if e.Author == nil {
		e.URL = "https://dsc.gg/gophers-latam"
	}
	e.Color = 0x78141b
	s.ChannelMessageSendEmbed(m.ChannelID, e)
}

func wordCase(s string) string {
	if len(s) == 0 {
		return s
	}

	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}
