package main

import "github.com/bwmarrin/discordgo"

type command struct {
	Name string
	Help string

	Exec func(*discordgo.Session, *discordgo.MessageCreate, []string) error
}
