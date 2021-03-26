package self

import (
	"github.com/TrizlyBear/ayr/ayr/types"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Ayr = types.Ayr{Commands: make(map[string]types.Command),Cogs: make(map[string]types.Cog)}
	DB = &mongo.Client{}
)