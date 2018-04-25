package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"strings"
	"github.com/spf13/viper"
)

func main() {
	loadConfig()

	openStorage()
	defer closeStorage()

	dg, err := discordgo.New("Bot " + viper.GetString("token"))
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	sayHello(dg)

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

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
