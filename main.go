package main

import (
	"fmt"
	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
	"log"
	"os"
)

type subscriberBot struct {
	bot reddit.Bot
}

func (s *subscriberBot) Mention(mention *reddit.Message) error {
	 fmt.Println("Detected a mention!")
	 return nil
}

func main() {
	app := reddit.App{
		ID:       os.Getenv("REDDIT_BOT_ID"),
		Secret:   os.Getenv("REDDIT_BOT_SECRET"),
		Username: os.Getenv("REDDIT_BOT_USERNAME"),
		Password: os.Getenv("REDDIT_BOT_PASSWORD"),
	}

	agentString := os.Getenv("AGENT_STRING_PLATFORM") + ":" +
		os.Getenv("AGENT_STRING_APP_ID") + ":" +
		os.Getenv("AGENT_STRING_VERSION")

	botConfig := reddit.BotConfig{
		App:    app,
		Agent: agentString,
	}

	bot, err := reddit.NewBot(botConfig)
	if err != nil {
		log.Fatal(err)
		return
	}

	cfg := graw.Config{Mentions: true}
	handler := &subscriberBot{bot: bot}

	if _, wait, err := graw.Run(handler, bot, cfg); err != nil {
		fmt.Println("Failed to start graw run: ", err)
	} else {
		fmt.Println("graw run failed: ", wait())
	}
}
