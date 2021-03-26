package handlers

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func Init(s *discordgo.Session) error {
	s.AddHandler(MessageCreate)
	err := s.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return err
	}
	return nil
}
