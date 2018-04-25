package main

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"time"
	"errors"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func handleChat(s *discordgo.Session, m *discordgo.MessageCreate) error {
	sentences, err := readLines("sentences.txt")
	if err != nil {
		return err
	}

	s.ChannelMessageSend(m.ChannelID, sentences[rand.Intn(len(sentences))])

	return nil
}

func sayHello(s *discordgo.Session) error {
	chs, err := s.UserChannels()
	if err != nil {
		return err
	}

	for _, ch := range chs {
		if ch.Name == "bot_land" {
			s.ChannelMessageSend(ch.ID, "Hello World!")
			return nil
		}
	}

	return errors.New("channel bot_land not found")
}