package types

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/http"
	"time"
)

type Ayr struct {
	S				*discordgo.Session
	Commands 		map[string]*Command
	Descriptions	map[string]string
	Color			int
	Plugins			map[string]*Plugin
	Tokens			map[string]string
	HTTPClient		*http.Client
}

func (e *Ayr) AddPlugin(p *Plugin) {
	e.Descriptions[p.Name] = p.Description

	for _,c := range p.Commands {
		e.Commands[c.Name] = c
		c.Plugin = p
		fmt.Println(e.S.State.User.ID)
		_, err := e.S.ApplicationCommandCreate(e.S.State.User.ID,"",c.ApplicationCommand)
		if c.Init != nil {
			c.Init(c)
		}
		if err != nil {
			log.Fatalf("Error resgistring command: %s", err)
		}

	}
}

func (e *Ayr) Embed() *discordgo.MessageEmbed {
	em := &discordgo.MessageEmbed{
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       e.Color,
		Fields:      []*discordgo.MessageEmbedField{},
	}
	return em
}