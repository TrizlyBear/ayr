package events

import (
	self "github.com/TrizlyBear/ayr/ayr/dispatcher"
	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate)  {
	if !m.Author.Bot {
		if command, ok := self.Ayr.Commands[m.Content]; ok {
			command.Exc(&self.Ayr,s,m)
		}

	}
}
