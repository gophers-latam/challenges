package bot

import (
	"database/sql"
	"errors"
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
			Description: "Get hours LATAM",
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
	}

	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"bot": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: `Hola **` + i.Member.User.Username + `** ` + defaultMsg,
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

			msg, err := GetChallenge(margs[0].(string), margs[1].(string))
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					unsuccessfulInteraction(s, i, `**Ups, sin desafios que coincidan**`)
				}
				return
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msg.ChallengeFmt(),
				},
			})
		},
		"facts": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			msg, err := GetFact()
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					unsuccessfulInteraction(s, i, `**Ups, sin hechos que mencionar**`)
				}
				return
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msg.Text,
				},
			})
		},
		"events": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			msg, err := GetEvents()
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					unsuccessfulInteraction(s, i, `**Ups, sin eventos para mostrar**`)
				}
				return
			}

			for _, e := range *msg {
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: e.Text,
					},
				})
			}
		},
		"hours": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			options := i.ApplicationCommandData().Options
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			margs := make([]interface{}, 0, len(options))

			if option, ok := optionMap["hour"]; ok {
				margs = append(margs, option.StringValue())
			}

			if option, ok := optionMap["country"]; ok {
				margs = append(margs, option.StringValue())
			}

			msg := GetHours(margs[0].(string), margs[1].(string))
			if msg == "" {
				unsuccessfulInteraction(s, i, `**Ups, no se puede mostrar equivalencia horaria**`)
				return
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: msg,
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
