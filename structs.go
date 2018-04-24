package main

import "github.com/bwmarrin/discordgo"

type command struct {
	Name string
	Help string

	Exec func(*discordgo.Session, *discordgo.MessageCreate, []string) error
}

type insult struct {
	Insult struct {
		ID            int         `json:"id"`
		Value         string      `json:"value"`
		TotalVoteUp   int         `json:"total_vote_up"`
		TotalVoteDown int         `json:"total_vote_down"`
		TotalVote     int         `json:"total_vote"`
		CurrentVote   interface{} `json:"current_vote"`
	} `json:"insult"`
}