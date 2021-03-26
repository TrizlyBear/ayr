package types

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

type Cog struct {
	Name		string
	Commands 	map[string]Command
}

type CommandExc func(self *Ayr, s *discordgo.Session, m *discordgo.MessageCreate)

type Command struct {
	Name		string
	Owner		bool
	Cooldown	time.Duration
	Exc 		CommandExc
	Commands	map[string]Command
}

func (cog *Cog) Inject(bot Ayr)  {
	for _,e := range cog.Commands {
		bot.Commands[e.Name] = e
	}
	bot.Cogs[cog.Name] = *cog
}

func (cog Cog) Add(command Command) {
	cog.Commands[command.Name] = command
}

func (cmd Command) Add(command Command) {
	cmd.Commands[command.Name] = command
}
