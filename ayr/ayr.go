package ayr

import (
	"fmt"
	"github.com/TrizlyBear/ayr/ayr/dispatcher"
	"github.com/TrizlyBear/ayr/ayr/plugins"
	"github.com/TrizlyBear/ayr/ayr/types"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func InteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate)  {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		if c, ok := dispatcher.Ayr.Commands[i.ApplicationCommandData().Name]; ok {
			go func() {
				err := c.R(s,i)
				if err != nil {
					log.Fatalf("Failed to execute command: %s - %s", i.ApplicationCommandData().Name, err)
				}
			}()
		}
	case discordgo.InteractionMessageComponent:
		fmt.Println(i.MessageComponentData().CustomID)
		cn := strings.Split(i.MessageComponentData().CustomID, "_")
		fmt.Println(strings.Join(cn[1:],"_"))
		if c, ok := dispatcher.Ayr.Commands[cn[0]]; ok {
			if c.IR != nil {
				c.IR(s, i, cn[1:])
			}
		}
	case discordgo.InteractionApplicationCommandAutocomplete:
		data := i.ApplicationCommandData()
		cname := data.Name
		if c, ok := dispatcher.Ayr.Commands[cname]; ok {
			c.AC(s,i)
		}

	}
}

func Init()  {
	// Load environment
	err := godotenv.Load("./config/.env")
	if err != nil {
		panic(err)
	}
	// Initialize bot
	dispatcher.Ayr.Tokens["genius"] = os.Getenv("AYR_GENIUS_TOKEN")
	dispatcher.Ayr.Tokens["discogs_key"] = os.Getenv("AYR_DISCOGS_KEY")
	dispatcher.Ayr.Tokens["discogs_secret"] = os.Getenv("AYR_DISCOGS_SECRET")
	s, err := discordgo.New("Bot " + os.Getenv("AYR_TOKEN"))
	if err != nil {
		panic(err)
	}
	s.Identify.Intents = discordgo.IntentsGuildMessages

	s.AddHandler(InteractionHandler)

	err = s.Open()

	if err != nil {
		panic(err)
	}

	dispatcher.Ayr.S = s

	ps := []*types.Plugin{
		plugins.Fun,
		plugins.General,
		plugins.Leveling,
		plugins.Moderation,
		plugins.Music,
		plugins.Utility,
	}

	for _,p := range ps {
		dispatcher.Ayr.Plugins[p.Name] = p
	}


	for _,p := range ps {
		dispatcher.Ayr.AddPlugin(p)
	}



	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	s.Close()
}