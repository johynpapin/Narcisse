package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
	"fmt"
)

func connectDiscord() (*discordgo.Session, error) {
	dg, err := discordgo.New("Bot " + viper.GetString("token"))
	if err != nil {
		return nil, fmt.Errorf("error creating Discord session: %s", err)
	}

	dg.AddHandler(handleMessageCreate)
	dg.AddHandler(handleGuildCreate)
	dg.AddHandler(handleGuildMemberAdd)

	err = dg.Open()
	if err != nil {
		return nil, fmt.Errorf("error opening connection: %s", err)
	}

	return dg, nil
}