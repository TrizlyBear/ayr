package fun

import (
	"github.com/TrizlyBear/ayr/ayr/dispatcher"
	"github.com/TrizlyBear/ayr/ayr/embed"
	"github.com/TrizlyBear/ayr/ayr/types"
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strings"
)

var RPC = &types.Command{
	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:          "rpc",
		Description:   "Play rock paper scissors",
		Version:       "0.0.1",
		Options:       []*discordgo.ApplicationCommandOption{
			&discordgo.ApplicationCommandOption{
				Type:         discordgo.ApplicationCommandOptionString,
				Name:         "choice",
				Description:  "Choice: either rock, paper or scissors",
				ChannelTypes: nil,
				Required:     true,
				Options:      nil,
				Autocomplete: false,
				Choices:      []*discordgo.ApplicationCommandOptionChoice{
					{
						Name:  "Rock",
						Value: "rock",
					},
					{
						Name:  "Paper",
						Value: "paper",
					},
					{
						Name:  "Scissors",
						Value: "scissors",
					},
				},
			},
		},
	},
	Alias:              nil,
	R: func(s *discordgo.Session, m *discordgo.InteractionCreate) error {
		choices := []string{"rock","paper","scissors"}
		choice := m.ApplicationCommandData().Options[0].StringValue()
		bot := choices[rand.Intn(3)]
		var result string
		if bot == choice {
			result = "It's a draw!"
		} else if (bot == "rock" && choice == "paper") || (bot == "scissors" && choice == "rock") || (bot == "paper" && choice == "scissors") {
			result = "Congratulations! You won."
		} else {
			result = "Oh no, You lost!"
		}

		e := embed.EmbedFrom(dispatcher.Ayr).SetTitle("Rock Paper Scissors")
		e.SetDescription(result)
		e.AddField("Your choice",strings.Title(choice),true)
		e.AddField("Ayr's choice",strings.Title(bot),true)

		e.Send(m.Interaction)
		return nil
	},
}
