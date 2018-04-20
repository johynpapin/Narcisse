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
	"github.com/coreos/bbolt"
)

var (
	Token string
	Db    *bolt.DB
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	db, err := bolt.Open("narcisse_timezones.Db", 0600, nil)
	Db = db
	if err != nil {
		fmt.Println("error connecting to database,", err)
	}
	defer Db.Close()

	Db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("timezones"))
		if err != nil {
			return fmt.Errorf("error creating bucket: %s", err)
		}
		return nil
	})

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
			s.ChannelMessageSend(m.ChannelID, "Veuillez passer au moins un paramètre.")
			return
		}

		switch msg[1] {
		case "timezone":
			if len(msg) <= 2 {
				s.ChannelMessageSend(m.ChannelID, "Veuillez passer un fuseau horaire en paramètre.")
			}

			loc, err := time.LoadLocation(msg[1])
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Je ne connais pas ce fuseau horaire.")
				return
			}

			now := time.Now().In(loc)

			s.ChannelMessageSend(m.ChannelID, "Dans ce fuseau horaire, il est **"+now.String()+"**.")
		case "set":
			if len(msg) <= 2 {
				s.ChannelMessageSend(m.ChannelID, "Veuillez indiquer votre fuseau horaire en paramètre.")
			}

			loc, err := time.LoadLocation(msg[2])
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Je ne connais pas ce fuseau horaire.")
				return
			}

			err = Db.Update(func(tx *bolt.Tx) error {
				b := tx.Bucket([]byte("timezones"))
				err := b.Put([]byte(m.Author.ID), []byte(msg[2]))
				return err
			})
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Erreur lors de l’accès à la base de données.")
			}

			now := time.Now().In(loc)

			s.ChannelMessageSend(m.ChannelID, "Votre fuseau horaire est désormais **"+msg[2]+"**, dans ce fuseau horaire, il est **"+now.String()+"**.")
		default:
			if len(m.Mentions) == 0 {
				s.ChannelMessageSend(m.ChannelID, "Veuillez mentionner au moins un utilisateur.")
			} else {
				for _, user := range m.Mentions {
					Db.View(func(tx *bolt.Tx) error {
						b := tx.Bucket([]byte("timezones"))
						timezone := b.Get([]byte(user.ID))

						if timezone == nil {
							s.ChannelMessageSend(m.ChannelID, "Je connais pas le fuseau horaire de **"+user.Username+"**.")
						} else {
							loc, err := time.LoadLocation(string(timezone))
							if err != nil {
								s.ChannelMessageSend(m.ChannelID, "Le fuseau horaire de **"+user.Username+"** est invalide.")
							}

							now := time.Now().In(loc)

							s.ChannelMessageSend(m.ChannelID, "Dans le fuseau horaire de **"+user.Username+"**, il est **"+now.String()+"**.")
						}

						return nil
					})
				}
			}
		}
	}
}
