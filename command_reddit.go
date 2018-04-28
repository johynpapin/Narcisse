package main

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
)

func init() {
	newCommand("reddit", helpReddit(), commandReddit)
}

func helpReddit() string {
	return "== Aide de la commande `insult` ==\n" +
		"`insult` affiche une insulte"
}

func commandReddit(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	harvest, err := rb.Listing("/r/programmerHumor", "")
	if err != nil {
		return fmt.Errorf("failed to fetch /r/programmerHumor: %s", err)
	}

	for _, post := range harvest.Posts[:5] {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("[%s] posted [%s]\n", post.Author, post.Title))
	}

	return nil
}
