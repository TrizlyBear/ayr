package types

import (
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/mongo"
)

type Ayr struct {
	Bot 		*discordgo.Session
	Commands 	map[string]Command
	Cogs		map[string]Cog
}

type DataBase *mongo.Client

