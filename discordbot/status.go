package messages

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func SetStatus(discord *discordgo.Session, ready *discordgo.Ready) {
	// TODO: update here when fixed by Discord
	//_, err := discord.UserUpdateStatus(discordgo.StatusDoNotDisturb)
	err := discord.UpdateStatusComplex(discordgo.UpdateStatusData{Status: "dnd"})
	if err != nil {
		log.Println(err.Error())
	}

	_ = discord.UpdateGameStatus(1, ".go Googleando")
	servers := discord.State.Guilds
	log.Printf("Bot iniciado en %d servers", len(servers))
}
