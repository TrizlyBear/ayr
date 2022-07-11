package music

import (
	"encoding/json"
	"github.com/TrizlyBear/ayr/ayr/dispatcher"
	"github.com/TrizlyBear/ayr/ayr/embed"
	"github.com/TrizlyBear/ayr/ayr/types"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/http"
	"strconv"
)

type Hit struct {
	Type			string	`json:"type"`
	Result			struct{
		ArtistNames	string	`json:"artist_names"`
		Id			int		`json:"id"`
		Title		string	`json:"title"`
	}						`json:"result"`
}

type GeniusResponse struct {
	Response	struct{
		Hits	[]*Hit	`json:"hits"`
	}					`json:"response"`
}

type SongResponse struct {
	Response 	struct{
				Song	struct{
					FullTitle		string	`json:"full_title"`
					SongArtImageUrl	string	`json:"song_art_image_url"`
					Url				string	`json:"url"`
				}
	}
}

var Lyrics = &types.Command{
	ApplicationCommand: &discordgo.ApplicationCommand{
		Version:                  "0.0.1",
		Name:                     "lyrics",
		Description:              "Searches the lyrics of a song.",
		Options:					[]*discordgo.ApplicationCommandOption{
			{
				Type:                     discordgo.ApplicationCommandOptionString,
				Name:                     "search",
				Description:              "Search for the lyrics of a song",
				Required:                 true,
				Autocomplete:             true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{},
			},
		},
	},
	IR:                 nil,
	Init:               nil,
	AC: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		data := i.ApplicationCommandData()
		query := data.Options[0].StringValue()
		if len(query) < 3 {
			return
		}
		req, err := http.NewRequest(http.MethodGet, "https://api.genius.com/search?q="+query, nil)
		if err != nil {
			log.Println(err)
			return
		}
		req.Header.Add("Authorization","Bearer " + dispatcher.Ayr.Tokens["genius"])

		resp, err := dispatcher.Ayr.HTTPClient.Do(req)
		if err != nil {
			log.Println(err)
			return
		}

		var parsed GeniusResponse

		if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
			log.Println(err)
			return
		}
		choices := []*discordgo.ApplicationCommandOptionChoice{}

		for _,hit := range parsed.Response.Hits {

			if hit.Type == "song" {
				choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
					Name:              hit.Result.ArtistNames + " - " + hit.Result.Title,
					Value:             strconv.Itoa(hit.Result.Id),
				})
			}
		}
		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionApplicationCommandAutocompleteResult,
			Data: &discordgo.InteractionResponseData{
				Choices:         choices,
			},
		})

		if err != nil {
			log.Println(err)
			return
		}
	},
	R: func(s *discordgo.Session, m *discordgo.InteractionCreate) error {
		data := m.ApplicationCommandData()
		songid := data.Options[0].StringValue()

		req, err := http.NewRequest(http.MethodGet, "https://api.genius.com/songs/"+songid+"?text_format=plain", nil)
		if err != nil {
			log.Println(err)
			return err
		}
		req.Header.Add("Authorization","Bearer " + dispatcher.Ayr.Tokens["genius"])

		resp, err := dispatcher.Ayr.HTTPClient.Do(req)
		if err != nil {
			log.Println(err)
			return err
		}

		var parsed SongResponse

		if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
			log.Println(err)
			return err
		}

		em := embed.NewEmbed()
		em.SetColor(0xffff64)
		em.SetImage(parsed.Response.Song.SongArtImageUrl)
		em.SetTitle(parsed.Response.Song.FullTitle)
		em.SetUrl(parsed.Response.Song.Url)

		err = em.Send(m.Interaction)
		return err
	},
}