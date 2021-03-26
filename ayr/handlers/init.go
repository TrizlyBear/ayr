package handlers

import (
	"fmt"
	"github.com/TrizlyBear/ayr/ayr/handlers/events"
	"github.com/TrizlyBear/ayr/plugins/test"
	"github.com/bwmarrin/discordgo"
)



func Init(s *discordgo.Session) error {
	s.AddHandler(events.MessageCreate)
	err := s.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return err
	}

	test.Init()

	return nil
}
