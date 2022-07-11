package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/TrizlyBear/ayr/ayr/database"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type ReqRes []database.PlayerPlain

func main()  {
	err := godotenv.Load("../../.././config/.env")

	if err != nil {
		panic(err)
	}

	uri := os.Getenv("AYR_DBURI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("ayr").Collection("players22")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}
		var players ReqRes
		err = json.Unmarshal(body, &players)
		if err != nil {
			fmt.Println(err)
			return
		}
		for _,p := range players {
			r,err := strconv.Atoi(p.Rating)
			if err != nil {
				panic(err)
			}
			rt,err := strconv.Atoi(p.RareType)
			if err != nil {
				panic(err)
			}
			id,err := strconv.Atoi(p.Id)
			if err != nil {
				panic(err)
			}
			rarev,err := strconv.Atoi(p.Rare)
			if err != nil {
				panic(err)
			}
			rare := false
			if rarev > 0 {
				rare = true
			}
			player := database.Player{
				Rating:      r,
				Position:    p.Position,
				ClubImage:   p.ClubImage,
				Image:       p.Image,
				RareType:    rt,
				FullName:    p.FullName,
				UrlName:     p.UrlName,
				Id:          id,
				NationImage: p.NationImage,
				Rare:        rare,
				Version:     p.Version,
			}
			var res database.Player
			err = coll.FindOne(context.TODO(),bson.D{{"id",player.Id}}).Decode(&res)
			if err == mongo.ErrNoDocuments {
				_, err = coll.InsertOne(context.TODO(), player)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				if err != nil {
					fmt.Println(err)
				}
			}
			fmt.Println(req.URL.Query())
		}
	})



	http.ListenAndServe(":8090",nil)

	//coll := client.Database("ayr").Collection("players")

	//fmt.Println(coll.FindOne(context.TODO(), bson.D{}))
}

var _=`
var alp = 'abcdefghijklmnopqrstuvwxyz'.split("")
//

await fetch("http://127.0.0.1:8090",{
  "mode":"cors"
})

 alp.forEach(a => {
  	 alp.forEach(b => {
    	alp.forEach(async (c) => {
      await fetch("https://www.futbin.com/search?year=22&extra=1&v=1&term="+a+b+c, {
    "credentials": "include",
    "headers": {
        "User-Agent": "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:91.0) Gecko/20100101 Firefox/91.0",
        "Accept": "application/json, text/javascript, */*; q=0.01",
        "Accept-Language": "en-US,en;q=0.5",
        "X-Requested-With": "XMLHttpRequest",
        "Sec-Fetch-Dest": "empty",
        "Sec-Fetch-Mode": "cors",
        "Sec-Fetch-Site": "same-origin"
    },
    "referrer": "https://www.futbin.com/",
    "method": "GET",
    "mode": "cors"
}).then(async(content)=>{
        var co = await content.json()
        console.log(co)
       	await fetch("http://127.0.0.1:8090/?query="+a+b+c,{
  				"mode":"cors",
          "method":"POST",
          "body":JSON.stringify(co)
})
      } );
    	
      await new Promise(r => setTimeout(r, 2000));  
    })
  })
})


`
