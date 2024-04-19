package bot

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

var (
	// Commands/options without description will fail the registration
	// of the command.
	commands = []*discordgo.ApplicationCommand{
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
			Description: "Show challenges help",
		},
		{
			Name:        "challenge",
			Description: "Get challenge",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "level",
					Description: "chanllenge {level}",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "type",
					Description: "chanllenge {type}",
					Required:    true,
				},
			},
		},
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"bot": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: `Hola **` + i.Member.User.Username + `** - envia: .go (para usar el gopherbot)`,
				},
			})
		},
		"help": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			msg, _ := GetCommand(".go help")
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msg.Text,
				},
			})
		},
		"challenge_help": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			msg, _ := GetCommand(".go challenge help")
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msg.Text,
				},
			})
		},
		"challenge": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// options in the order provided by the user.
			options := i.ApplicationCommandData().Options

			// convert the slice options into a map
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			margs := make([]interface{}, 0, len(options))

			if option, ok := optionMap["level"]; ok {
				margs = append(margs, option.StringValue())
			}

			if option, ok := optionMap["type"]; ok {
				margs = append(margs, option.StringValue())
			}

			msg, _ := GetChallenge(margs[0].(string), margs[1].(string))
			text := `[*Challenge*] 
			-**Level:** ` + string(msg.Level) + ` -**Type:** ` + string(msg.ChallengeType) + `
			-**Description: ** ` + msg.Description

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: text,
				},
			})
		},
	}
)

func InitSlhCmd(s *discordgo.Session) []*discordgo.ApplicationCommand {
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

func SlhCmd(s *discordgo.Session, i *discordgo.InteractionCreate) {
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
