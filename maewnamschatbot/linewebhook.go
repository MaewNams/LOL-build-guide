package maewnamschatbot

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/line/line-bot-sdk-go/linebot"
)

// Claims struct
type Claims struct {
	Sub string `json:"sub"`
	Iss string `json:"iss"`
	Aud string `json:"aud"`
	jwt.StandardClaims
}

var (
	channelSecret = os.Getenv("CHANNEL_SECRET")
	channelToken  = os.Getenv("CHANNEL_TOKEN")
	jwtKey        = []byte(userSecret)
	sKey          = []byte(mumuSecret)
	httpClient    = &http.Client{
		Timeout: 10 * time.Second,
	}

	config *Config
)

func init() {
	config, _ = ReadConfig("config.yml")
}

// LineWebhook will be called by Line when Line user interact with the bot.
func LineWebhook(w http.ResponseWriter, r *http.Request) {
	// TODO: use conf path as ENV
	bot, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		log.Panicln("Error creating Line bot:", err.Error())
		w.WriteHeader(200)
		return
	}

	events, err := bot.ParseRequest(r)
	if err != nil {
		log.Panicln("Error receiving webhook request:", err.Error())
	} else {
		for _, event := range events {
			err = ProcessEvent(config, bot, event)
			if err != nil {
				log.Panicln("Error processing webhook event:", err.Error())
			}
		}
	}
	// always returns 200
	w.WriteHeader(200)
}

// ProcessEvent processes individual Line webhook event
func ProcessEvent(config *Config, bot *linebot.Client, event *linebot.Event) error {

	if event.Type == linebot.EventTypeMessage {
		userID := event.Source.UserID
		switch message := event.Message.(type) {
		case *linebot.TextMessage:
				textMsg := linebot.NewTextMessage("Hello world")
				_, err := bot.ReplyMessage(event.ReplyToken, textMsg).Do()
				if err != nil {
					log.Println("Error replying message")
					return err
				}
			}

		case *linebot.StickerMessage:
			log.Println("Received sticker message")

	}
	return nil

}
