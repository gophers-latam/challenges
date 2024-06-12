package bot

import (
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gophers-latam/challenges/bot/subcmd_commands"
	"github.com/gophers-latam/challenges/global"
)

// Function to handle the message
func HandleSubCmd(s *discordgo.Session, m *discordgo.MessageCreate) {
	if isMessageFromBot(m.Author.ID, s.State.User.ID) {
		return
	}

	args := getCommandArgs(m.Content, global.Prefix)
	if args == nil {
		return
	}

	if isOnlyPrefix(args, global.Prefix) {
		if helloCommand, ok := SubCmdCommands["hello"]; ok {
			helloCommand.Execute(s, m)
		}
		return
	}

	// Proceed with further command processing
	SubCmd(s, m, args)
}

func SubCmd(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {

	cmd := args[1]
	command, ok := SubCmdCommands[cmd]
	if !ok {
		subcmd_commands.MsgCommands(s, m)
		return
	}

	// Check if the second argument is 'help'
	if len(args) > 2 && strings.ToLower(args[2]) == "help" {
		s.ChannelMessageSend(m.ChannelID, command.Help())
		return
	}

	command.Execute(s, m)

}

// Helper function to check if the message is from the bot itself
func isMessageFromBot(authorID, botID string) bool {
	return authorID == botID
}

// Helper function to split the message into arguments and check the prefix
func getCommandArgs(content, prefix string) []string {
	args := strings.Fields(content)
	if len(args) > 0 && args[0] == prefix {
		return args
	}
	return nil
}

// Helper function to check if only the prefix is present
func isOnlyPrefix(args []string, prefix string) bool {
	return len(args) == 1 && args[0] == prefix
}

func Stat(discord *discordgo.Session, ready *discordgo.Ready) {
	err := discord.UpdateStatusComplex(discordgo.UpdateStatusData{Status: "dnd"})
	if err != nil {
		log.Println(err.Error())
	}

	_ = discord.UpdateGameStatus(1, ".go help")
	servers := discord.State.Guilds
	log.Printf("Bot running on %d servers", len(servers))
}
