package main

import (
	"github.com/bwmarrin/discordgo"
)

func init() {
	newCommand("insult", helpInsult(), commandInsult)
}

func helpInsult() string {
	return "== Aide de la commande `insult` ==\n" +
		"`insult` affiche une insulte"
}

func commandInsult(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	insult := new(insult)

	err := getJson("https://www.insult.es/api/random", insult)
	if err != nil {
		return err
	}

	s.ChannelMessageSend(m.ChannelID, insult.Insult.Value)

	return nil
}
