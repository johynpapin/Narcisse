package main

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"time"
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