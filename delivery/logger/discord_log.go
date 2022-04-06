package logger

import (
	"fmt"
	"github.com/gtuk/discordwebhook"
	"github.com/rs/zerolog"
	"gokost.com/m/config"
	"os"
	"time"
)

var (
	Log *zerolog.Logger
)

func NewLoger(logLevel string) {
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	if logLevel == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	loger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	Log = &loger
}

func SendLogToDiscord(serve string, errSend error) {
	var username = config.GetConfigValue("DISCORD_USERNAME")
	var url = config.GetConfigValue("DISCORD_WEBHOOK_URL")

	currentTime := time.Now().Format("2006-01-02 15:04:05")
	content := fmt.Sprintf("(%s) Serve: %s, Error: %s", currentTime, serve, errSend.Error())
	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
	}

	discordwebhook.SendMessage(url, message)
	//if err != nil {
	//	log.Fatal(err)
	//}
	Log.Err(errSend).Msg(serve)
}
