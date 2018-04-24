package main

import (
	"github.com/bwmarrin/discordgo"
	"fmt"
	"strings"
)

func init() {
	newCommand("help", helpHelp(), commandHelp)
}

func helpHelp() string {
	return "== Aide de la commande `help` ==" +
		"`help` affiche l’aide de toutes les commandes" +
		"`help <commande>` affiche l’aide de la commande <commande>"
}

func commandHelp(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	if len(args) == 0 {
		fmt.Println("== Aide ==")
		for _, command := range commands {
			s.ChannelMessageSend(m.ChannelID, command.Help)
		}
	} else {
		if command, exist := commands[strings.ToLower(args[0])]; exist {
			s.ChannelMessageSend(m.ChannelID, command.Help)
		} else {
			s.ChannelMessageSend(m.ChannelID, "Je ne connais pas cette commande. Veuillez utiliser la commande `help` pour afficher la liste de toutes les commandes.")
		}
	}

	return nil
}
