package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gophers-latam/challenges/bot/subcmd_commands"
)

// Define the Command interface
type Command interface {
	Execute(s *discordgo.Session, m *discordgo.MessageCreate)
	Help() string
}

// Initialize the SubCmdCommands map
var SubCmdCommands = map[string]Command{}

func SubCmdRegisterCommands() {
	SubCmdCommands["hours"] = &subcmd_commands.HoursCommand{}
	SubCmdCommands["challenge"] = &subcmd_commands.ChallengeCommand{}
	SubCmdCommands["events"] = &subcmd_commands.EventsCommand{}
	SubCmdCommands["facts"] = &subcmd_commands.FactsCommand{}
	SubCmdCommands["hello"] = &subcmd_commands.HelloCommand{}
}
