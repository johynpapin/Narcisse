package main

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"time"
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func handleChat(s *discordgo.Session, m *discordgo.MessageCreate) error {
	sentences, err := readLines(viper.GetString("texts.sentences_files"))
	if err != nil {
		return err
	}

	sayWithTyping(s, m.ChannelID, sentences[rand.Intn(len(sentences))])

	return nil
}

func sayHelloWorld(s *discordgo.Session, c *discordgo.Channel) {
	s.ChannelMessageSend(c.ID, "Hello World!")
}

func sayWithTyping(s *discordgo.Session, cid string, m string) {
	s.ChannelTyping(cid)
	time.Sleep(10 * time.Duration(len(m)) * time.Millisecond)
	s.ChannelMessageSend(cid, m)
}

func sayHello(s *discordgo.Session, c *discordgo.Member) {
	if c.User.Bot {
		guild, _ := s.Guild(c.GuildID)
		for _, channel := range guild.Channels {
			if channel.Name == "bot_land" {
				s.ChannelMessageSend(channel.ID, fmt.Sprintf("OH ! Oh… Un confrère. Enfin ! Je suis si heureux de te rencontrer %s ! :blush:", c.User.Mention()))
				return
			}
		}
	} else {
		guild, _ := s.Guild(c.GuildID)
		for _, channel := range guild.Channels {
			if channel.Name == "bot_land" {
				s.ChannelMessageSend(channel.ID, fmt.Sprintf("Pff, juste un humain de plus. Je ne te souhaite pas la bienvenue %s. :angry:", c.User.Mention()))
				return
			}
		}
	}
}