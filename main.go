package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
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
		if len(msg) == 1 {
			s.ChannelMessageSend(m.ChannelID, "Veuillez indiquer un fuseau horaire.")
			return
		}

		loc, err := time.LoadLocation(msg[1])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Je ne connais ce fuseau horaire.")
			return
		}

		//set timezone,
		now := time.Now().In(loc)

		s.ChannelMessageSend(m.ChannelID, "Dans ce fuseau horaire, il est **" + now.String() + "**.")
	}
}
