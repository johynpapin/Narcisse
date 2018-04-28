package main

import (
	"github.com/turnage/graw/reddit"
	"fmt"
)

func connectReddit() (reddit.Bot, error){
	bot, err := reddit.NewBotFromAgentFile("reminderbot.agent", 0)
	if err != nil {
		return nil, fmt.Errorf("error creating Reddit session: %s", err)
	}

	return bot, nil
}