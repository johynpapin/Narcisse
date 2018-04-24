package main

import "github.com/bwmarrin/discordgo"

func handleChat(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Oui ? On me demande ?")
}