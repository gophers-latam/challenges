package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gophers-latam/challenges/bot/subcmd_commands"
)

// Define the Command interface
type Command interface {
	Execute(s *discordgo.Session, m *discordgo.MessageCreate)
	Help(string) string
}

// Initialize the SubCmds map
var SubCmds = map[string]Command{}

func RegisterSubCmds() {
	SubCmds["hours"] = &subcmd_commands.HoursCommand{}
	SubCmds["challenge"] = &subcmd_commands.ChallengeCommand{}
	SubCmds["events"] = &subcmd_commands.EventsCommand{}
	SubCmds["facts"] = &subcmd_commands.FactsCommand{}
	SubCmds["hello"] = &subcmd_commands.HelloCommand{}
}
