package ayr

import (
	"context"
	"fmt"
	"github.com/TrizlyBear/ayr/ayr/dispatcher"
	"github.com/TrizlyBear/ayr/ayr/handlers"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"
)

func Init()  {
	_, b, _, _ := runtime.Caller(0)
	ProjectRootPath := filepath.Join(filepath.Dir(b), "../")
	err := godotenv.Load(ProjectRootPath+"/config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	fmt.Println("Connecting bot...")
	bot, err := discordgo.New("Bot "+ os.Getenv("AYR_TOKEN"))
	self.Ayr.Bot = bot
	if err != nil {
		log.Fatal(err)
		return
	} else {
		fmt.Println("Connected bot")
	}

	fmt.Println("Initializing Database client")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://" + os.Getenv("AYR_DBIP") + ":" + os.Getenv("AYR_DBPORT")))
	if err != nil {
		fmt.Println("Couldn't connect to Database",err)
	}
	self.DB = client

	err = handlers.Init(bot)
	if err != nil {
		return
	}

	fmt.Println("Serving")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	bot.Close()
}
