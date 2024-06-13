package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/gophers-latam/challenges/bot/slashcmd_commands"
)

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	// TODO: refactor to strategy
	"bot":            slashcmd_commands.SlashHello,
	"help":           slashcmd_commands.SlashHelp,
	"challenge_help": slashcmd_commands.SlashChallengeHelp,
	"challenge":      slashcmd_commands.SlashChallenge,
	"facts":          slashcmd_commands.SlashFacts,
	"events":         slashcmd_commands.SlashEvents,
	"hours":          slashcmd_commands.SlashHours,
}

func HandleSlhCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if h, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
		h(s, i)
	}
}

func RemoveSlhCmd(s *discordgo.Session, cmd []*discordgo.ApplicationCommand) {
	log.Println("Removing slash commands...")
	// deleting only the commands that we added to recreate in next start
	// registeredCommands, err := s.ApplicationCommands(s.State.User.ID, *GuildID)

	for _, v := range cmd {
		err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID)
		if err != nil {
			log.Panicf("Cannot delete '%v' command: %v", v.Name, err)
		}
	}
}
