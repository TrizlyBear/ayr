package ayr

import (
	"fmt"
	"github.com/TrizlyBear/ayr/ayr/handlers"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
)

func Init()  {
	_, b, _, _ := runtime.Caller(0)
	ProjectRootPath := filepath.Join(filepath.Dir(b), "../")
	err := godotenv.Load(ProjectRootPath+"/config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	bot, err := discordgo.New("Bot "+ os.Getenv("AYR_TOKEN"))
	if err != nil {
		log.Fatal(err)
		return
	}

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
