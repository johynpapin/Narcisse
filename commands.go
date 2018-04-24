package main

import (
	"github.com/bwmarrin/discordgo"
	"strings"
	"fmt"
)

var commands = make(map[string]command)

func parseCommand(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	if len(args) == 0 {
		s.ChannelMessageSend(m.ChannelID, "Que dois-je faire ? Pour avoir la liste des commandes, utilisez la commande `help`.")
		return nil
	}

	if command, exist := commands[strings.ToLower(args[0])]; exist {
		command.Exec(s, m, args[1:])
	}

	s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Je n’ai pas compris la commande `%s`, désolé. Pour avoir la liste des commandes, utilisez la commande `help`.", args[0]))

	return nil
}

func newCommand(name string, help string, exec func(*discordgo.Session, *discordgo.MessageCreate, []string) error) command {
	commands[name] = command{
		Name: name,
		Help: help,
		Exec: exec,
	}

	return commands[name]
}
