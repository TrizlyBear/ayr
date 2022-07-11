package ayr

import (
	"github.com/bwmarrin/discordgo"
)

func AddComponent(i *discordgo.InteractionResponse, c discordgo.MessageComponent) *discordgo.InteractionResponse {
	i.Data.Components = append(i.Data.Components, c)
	return i
}

