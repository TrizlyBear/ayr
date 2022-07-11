package music

import (
	"encoding/json"
	"fmt"
	"github.com/TrizlyBear/ayr/ayr/dispatcher"
	"github.com/TrizlyBear/ayr/ayr/embed"
	"github.com/TrizlyBear/ayr/ayr/types"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type DiscogsQueryRes struct {
	Results	[]struct {
		Title	string		`json:"title"`
		Id		int			`json:"id"`
		Format	[]string	`json:"format"`
		Year	string		`json:"year"`
	} `json:"results"`


}

type DiscogsRelease struct {
	Title			string		`json:"title"`
	Thumb			string		`json:"thumb"`
	Artists	[]struct{
		Name		string		`json:"name"`
		ResourceUrl	string		`json:"resource_url"`
	}							`json:"artists"`
	Country			string		`json:"country"`
	Genres			[]string 	`json:"genres"`
	Formats []struct{
		Descriptions	[]string	`json:"descriptions"`
		Name		string		`json:"name"`
		Qty			string		`json:"qty"`
	}							`json:"formats"`
	Tracklist	[]struct{
		Duration	string		`json:"duration"`
		Position	string		`json:"position"`
		Title		string		`json:"title"`
	}	`json:"tracklist"`
	LowestPrice		float64		`json:"lowest_price"`
	Uri				string		`json:"uri"`
	Year			int			`json:"year"`
	NumForSale		int			`json:"num_for_sale"`
}

var Discogs = &types.Command{
	ApplicationCommand: &discordgo.ApplicationCommand{
		Version:     "0.0.1",
		Name:        "discogs",
		Description: "Searches the a release on Discogs.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:         discordgo.ApplicationCommandOptionString,
				Name:         "search",
				Description:  "Search for the lyrics of a song",
				Required:     true,
				Autocomplete: true,
				Choices:      []*discordgo.ApplicationCommandOptionChoice{},
			},
			{
				Type: 	discordgo.ApplicationCommandOptionString,
				Name:	"currency",
				Description: "Select the currency for the pricing information.",
				Required: false,
				Autocomplete: false,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{"United States Dollar",nil,"USD"},
					{"British Pound",nil,"GBP"},
					{"Euro",nil,"EUR"},
					{"Canadian Dollar",nil,"CAD"},
					{"Australian Dollar",nil,"AUD"},
					{"Japanese Yen",nil,"JPY"},
					{"Swiss Franc",nil,"CHF"},
					{"Mexican Peso",nil,"MXN"},
					{"Brazilian Real",nil,"BRL"},
					{"New Zealand Dollar",nil,"NZD"},
					{"Swedish Krona",nil,"SEK"},
					{"South African Rand",nil,"ZAR"},
				},
				MaxValue: 1,
			},
		},
	},
	AC: func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		data := i.ApplicationCommandData()
		query := data.Options[0].StringValue()
		if len(query) < 3 {
			return
		}
		req, _ := http.NewRequest(http.MethodGet, "https://api.discogs.com/database/search?q="+query+"&type=release", nil)
		req.Header.Add("Authorization","Discogs key="+dispatcher.Ayr.Tokens["discogs_key"]+", secret="+dispatcher.Ayr.Tokens["discogs_secret"])
		res, err := dispatcher.Ayr.HTTPClient.Do(req)
		if err != nil {
			log.Println(err)
			return
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println(err)
			return
		}
		var result DiscogsQueryRes
		err = json.Unmarshal(body, &result)
		if err != nil {
			log.Println(err)
			return
		}

		if len(result.Results) > 10 {
			result.Results = result.Results[:10]
		}

		choices := []*discordgo.ApplicationCommandOptionChoice{}

		for _, r := range result.Results {
			name :=  r.Title + " ("+strings.Join(r.Format,", ")+", "+r.Year+")"
			if len(name) > 100 {
				name = name[:96]+"..."
			}
			choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
				Name:             name,
				NameLocalizations: nil,
				Value:             strconv.Itoa(r.Id),
			})
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
		releaseid := data.Options[0].StringValue()
		currency := "EUR"
		if len(data.Options) > 1 {
			currency = data.Options[1].StringValue()
		}
		req, _ := http.NewRequest(http.MethodGet, "https://api.discogs.com/releases/"+releaseid+"?curr_abbr="+currency, nil)
		req.Header.Add("Authorization","Discogs key="+dispatcher.Ayr.Tokens["discogs_key"]+", secret="+dispatcher.Ayr.Tokens["discogs_secret"])

		res, err := dispatcher.Ayr.HTTPClient.Do(req)
		if err != nil {
			log.Println(err)
			return err
		}

		body, err := ioutil.ReadAll(res.Body)

		var release DiscogsRelease

		err = json.Unmarshal(body, &release)
		if err != nil {
			log.Println(err)
			return err
		}

		em := embed.EmbedFrom(dispatcher.Ayr)
		em.SetTitle(release.Title)
		em.SetUrl(release.Uri)
		em.Thumbnail = &discordgo.MessageEmbedThumbnail{
			URL:      release.Thumb,
		}
		em.AddField("Artist", "["+release.Artists[0].Name+"]("+release.Artists[0].ResourceUrl+")",false)
		em.AddField("Country / Region", release.Country, true)
		em.AddField("Year", strconv.Itoa(release.Year),true)

		var formats []string

		for _, f := range release.Formats {
			formats = append(formats, f.Qty+"x"+f.Name+ " ("+strings.Join(f.Descriptions,", ")+")")
		}

		em.AddField("Format(s)", strings.Join(formats,", "),false)

		var tracks []string

		for _, t := range release.Tracklist {
			tracks = append(tracks, t.Position+" \t\t"+t.Duration+" \t\t"+t.Title)
		}

		em.AddField("Tracklist",strings.Join(tracks,"\n"),false)

		em.AddField("Lowest price", fmt.Sprintf("%.2f", release.LowestPrice) + " " + currency, true)
		em.AddField("For sale",strconv.Itoa(release.NumForSale), true)


		err = em.Send(m.Interaction)
		return err
	},
}
