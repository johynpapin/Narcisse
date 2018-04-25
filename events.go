package main

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	msg := strings.Split(m.Content, " ")

	if msg[0] == "!n" || msg[0] == "!narcisse" {
		err := parseCommand(s, m, msg[1:])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Désolé, j’ai rencontré un problème interne.")
		}
	} else if ch, _ := s.Channel(m.ChannelID); ch.Name == "bot_land" && strings.Contains(strings.ToLower(m.ContentWithMentionsReplaced()), "narcisse") {
		handleChat(s, m)
	}
}

func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.Name == "bot_land" {
			sayHello(s, channel)
			return
		}
	}
}
