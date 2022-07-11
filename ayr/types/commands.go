package types

import "github.com/bwmarrin/discordgo"

type Plugin struct {
	Name		string
	Emoji		discordgo.ComponentEmoji
	Description	string
	Commands	[]*Command
}

// Commands structure:
// 	IR: Runs on interaction response
//	Init: Runs when initializing the application command
//	R: Runs when command is executed
type Command struct {
	*discordgo.ApplicationCommand
	Plugin		*Plugin
	Alias		[]string
	IR			func(s *discordgo.Session, i *discordgo.InteractionCreate, args []string)
	Init		func(c *Command)
	AC			func(s *discordgo.Session, i *discordgo.InteractionCreate)
	R			func(s *discordgo.Session, m *discordgo.InteractionCreate)error
}
