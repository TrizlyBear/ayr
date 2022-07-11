package dispatcher

import (
	"github.com/TrizlyBear/ayr/ayr/types"
	"net/http"
)

const (
	AppId 	= "590546345979412482"
	GuildId = "400653805013696512"
)

var (
	Ayr = &types.Ayr{
		Descriptions: map[string]string{},
		Commands: map[string]*types.Command{},
		Color:		0xFFD700 ,
		HTTPClient: &http.Client{},
		Plugins: map[string]*types.Plugin{},
		Tokens: map[string]string{},
	}
)
