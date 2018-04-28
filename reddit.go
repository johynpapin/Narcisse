package main

import (
	"github.com/turnage/graw/reddit"
	"fmt"
	"github.com/spf13/viper"
)

func connectReddit() (reddit.Bot, error){
	bot, err := reddit.NewBotFromAgentFile(viper.GetString("reddit.agent_file"), 0)
	if err != nil {
		return nil, fmt.Errorf("error creating Reddit session: %s", err)
	}

	return bot, nil
}