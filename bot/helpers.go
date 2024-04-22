package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

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

func msgEmbed(s *discordgo.Session, m *discordgo.MessageCreate, t, d string) {
	embed := &discordgo.MessageEmbed{
		Title:       t,
		URL:         "https://discord.gg/AEarh2kSvn",
		Description: d,
		Color:       0x78141b,
	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}
