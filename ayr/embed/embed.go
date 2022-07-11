package embed

import (
	"github.com/TrizlyBear/ayr/ayr/dispatcher"
	"github.com/TrizlyBear/ayr/ayr/types"
	"github.com/bwmarrin/discordgo"
)

type Embed struct {
	*discordgo.MessageEmbed
}

func (e *Embed) E() *discordgo.MessageEmbed {
	return e.MessageEmbed
}

func NewEmbed() *Embed {
	if dispatcher.Ayr != nil {
		return &Embed{dispatcher.Ayr.Embed()}
	}
	return &Embed{&discordgo.MessageEmbed{}}
}

func EmbedFrom(a *types.Ayr) *Embed {
	return &Embed{a.Embed()}
}

func (e *Embed) SetTitle(title string) *Embed {
	e.Title = title
	return e
}

func (e *Embed) SetDescription(description string) *Embed {
	e.Description = description
	return e
}

func (e *Embed) SetTimestamp(timestamp string) *Embed {
	e.Timestamp = timestamp
	return e
}

func (e *Embed) AddField(title string, value string, inline bool) *Embed {
	if e.Fields == nil {
		e.Fields = []*discordgo.MessageEmbedField{}
	}
	e.Fields = append(e.Fields, &discordgo.MessageEmbedField{
		Name:  	title,
		Value:  value,
		Inline: inline,
	})

	return e
}

func (e *Embed) SetColor(color int) *Embed {
	e.Color = color
	return e
}

func (e *Embed) SetImage(url string) *Embed  {
	e.Image = &discordgo.MessageEmbedImage{
		URL:	url,
	}

	return e
}

func (e *Embed) SetUrl(url string) *Embed {
	e.URL = url
	return e
}

func (e *Embed) Return() *discordgo.InteractionResponse {
	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds:	[]*discordgo.MessageEmbed{e.E()},
		},
	}
}

func NewError(title string, description string) *Embed {
	e := NewEmbed()
	e.SetTitle("Error - " + title)
	e.SetDescription(description)
	e.SetColor(0xED4337)

	return e
}

func (e *Embed) Send(i *discordgo.Interaction) error {
	resp := e.Return()
	err := dispatcher.Ayr.S.InteractionRespond(i,resp)
	if err != nil {
		return err
	}
	return nil
}