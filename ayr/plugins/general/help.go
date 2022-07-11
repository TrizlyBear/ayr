package general

import (
	"github.com/TrizlyBear/ayr/ayr/dispatcher"
	"github.com/TrizlyBear/ayr/ayr/embed"
	"github.com/TrizlyBear/ayr/ayr/types"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var Help = &types.Command{
	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:          "help",
		Description:   "Returns a collection of all commands available",
		Version:       "0.0.1",
		Options:		[]*discordgo.ApplicationCommandOption{},

		NameLocalizations: &map[discordgo.Locale]string{
			discordgo.Dutch: "help",
			discordgo.Swedish: "hjälp",
		},
	},
	IR: func(s *discordgo.Session, i *discordgo.InteractionCreate, args []string){
		category := i.MessageComponentData().Values[0]
		pc := dispatcher.Ayr.Plugins[category]
		em := embed.EmbedFrom(dispatcher.Ayr)
		em.SetTitle("Help - "+strings.Title(category))
		for _,cmd := range pc.Commands {
			em.AddField(strings.Title(cmd.Name),cmd.Description, false)
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseUpdateMessage,
			Data: &discordgo.InteractionResponseData{
				Embeds:          []*discordgo.MessageEmbed{em.E()},
			},
		})
	},
	R: func(s *discordgo.Session, m *discordgo.InteractionCreate) error {
		em := embed.NewEmbed().SetTitle("Help")

		for p,plug := range dispatcher.Ayr.Plugins {
			var cmds []string
			for _,c := range plug.Commands {
				cmds = append(cmds, c.Name)
			}
			em.AddField(strings.Title(p),"_"+plug.Description+"_\n"+strings.Join(cmds, " • "),false)
		}

		r := em.Return()

		var comps []discordgo.SelectMenuOption

		for _,p := range dispatcher.Ayr.Plugins {
			comps = append(comps, discordgo.SelectMenuOption{
				Label:    		strings.Title(p.Name),
				Description: 	p.Description,
				Emoji:    		p.Emoji,
				Value: 			p.Name,
			})
		}

		r.Data.Components = append(r.Data.Components, discordgo.ActionsRow{Components: []discordgo.MessageComponent{
			discordgo.SelectMenu{
				CustomID:    "help",
				Placeholder: "Select a category",
				MinValues:   nil,
				MaxValues:   1,
				Options:     comps,
				Disabled:    false,
			},
		}})
		err := s.InteractionRespond(m.Interaction,r)
		return err
	},
}
