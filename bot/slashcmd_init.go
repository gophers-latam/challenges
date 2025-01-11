package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// Commands/options without description will fail the registration
// of the command.
var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "bot",
		Description: "Call bot",
	},
	{
		Name:        "help",
		Description: "Show .go help",
	},
	{
		Name:        "challenge_help",
		Description: "Show .go challenges help",
	},
	{
		Name:        "challenge",
		Description: "Show .go challenge",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "level",
				Description: "challenge {level}",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "type",
				Description: "challenge {type}",
				Required:    true,
			},
		},
	},
	{
		Name:        "facts",
		Description: "Show .go facts",
	},
	{
		Name:        "events",
		Description: "Show .go events",
	},
	{
		Name:        "hours",
		Description: "Show .go hours (LATAM)",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "hour",
				Description: "24h format {HH:MM}",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "country",
				Description: "country {name}",
				Required:    true,
			},
		},
	},
	{
		Name:        "waifu",
		Description: "Show .go waifu",
	},
	{
		Name:        "media",
		Description: "Show .go media",
	},
}

func RegisterSlhCmds(s *discordgo.Session) []*discordgo.ApplicationCommand {
	log.Println("Adding slash commands...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for i, v := range commands {
		cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
		if err != nil {
			log.Panicf("Cannot create '%v' command: %v", v.Name, err)
		}
		registeredCommands[i] = cmd
	}

	return registeredCommands
}
