package handlers

import "github.com/bwmarrin/discordgo"

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate)  {
	if !m.Author.Bot && m.Content == "Ping!"{
		s.ChannelMessageSend(m.ChannelID, "Pong! :ping_pong: "+s.HeartbeatLatency().String())
	}

}
