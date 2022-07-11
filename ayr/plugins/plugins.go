package plugins

import (
	"github.com/TrizlyBear/ayr/ayr/plugins/fun"
	"github.com/TrizlyBear/ayr/ayr/plugins/general"
	"github.com/TrizlyBear/ayr/ayr/plugins/music"
	"github.com/TrizlyBear/ayr/ayr/types"
	"github.com/bwmarrin/discordgo"
)

var (
	Fun = &types.Plugin{
		Name:        	"fun",
		Emoji: 			discordgo.ComponentEmoji{Name: "🎮"},
		Description: 	"Entertaining commands.",
		Commands:		[]*types.Command{fun.RPC},
	}

	Utility = &types.Plugin{
		Name:        	"packbot",
		Emoji: 			discordgo.ComponentEmoji{Name: "📦"},
		Description: 	"Commands regarding a card collection game based on FiFa players.",
		Commands:    	[]*types.Command{},
	}

	Moderation = &types.Plugin{
		Name:       	"moderation",
		Emoji: 			discordgo.ComponentEmoji{Name: "⚒"},
		Description:	"Commands for moderation.",
		Commands:    	[]*types.Command{},
	}

	Music = &types.Plugin{
		Name:        	"music",
		Emoji: 			discordgo.ComponentEmoji{Name:"🎧"},
		Description: 	"Commands for searching lyrics, checking last.fm statistics, etc.",
		Commands:    	[]*types.Command{music.Lyrics, music.Discogs},
	}

	General = &types.Plugin{
		Name:        	"general",
		Emoji: 			discordgo.ComponentEmoji{Name: "🗃"},
		Description: 	"Informative and general purpose commands.",
		Commands:    	[]*types.Command{general.Help},
	}

	Leveling = &types.Plugin{
		Name:        	"leveling",
		Emoji: 			discordgo.ComponentEmoji{Name: "🥇"},
		Description: 	"A plugin to implement a custom leveling system.",
		Commands:    	[]*types.Command{},
	}
)

