package test

import (
	self "github.com/TrizlyBear/ayr/ayr/dispatcher"
	"github.com/TrizlyBear/ayr/ayr/types"
	"github.com/bwmarrin/discordgo"
)

var Testcog = &types.Cog{
	Name:     "Test",
	Commands: make(map[string]types.Command),
}

var test = types.Command{
	Name:     	"lesgo",
	Owner:    	false,
	Cooldown: 	0,
	Exc: 		func(self *types.Ayr, s *discordgo.Session, m *discordgo.MessageCreate) {
		s.ChannelMessageSend(m.ChannelID,"https://tenor.com/view/dababy-convertable-gif-20206040")
	},
	Commands: nil,
}

func Init()  {
	Testcog.Add(test)
	Testcog.Inject(self.Ayr)
}


