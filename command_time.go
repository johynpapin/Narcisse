package main

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

func init() {
	newCommand("time", helpTime(), commandSet)
}

func helpTime() string {
	return "== Aide de la commande `time` ==\n" +
		"`time set <fuseau horaire>` définir votre fuseau horaire\n" +
		"`time get <fuseau horaire>` récupérer l’heure d’un fuseau horaire\n" +
		"`time @user` récupérer l’heure de l’utilisateur user"
}

func commandSet(s *discordgo.Session, m *discordgo.MessageCreate, args []string) error {
	if len(args) < 1 {
		s.ChannelMessageSend(m.ChannelID, "Il manque des paramètres. Utilisez la commande `help time` pour accéder à l’aide.")
		return nil
	}

	command := args[0]
	args = args[1:]

	switch command {
	case "set":
		if len(args) < 1 {
			s.ChannelMessageSend(m.ChannelID, "Il manque le fuseau horaire. Utilisez la commande `help time` pour accéder à l’aide.")
			return nil
		}

		loc, err := time.LoadLocation(args[0])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Je ne connais pas ce fuseau horaire, désolé.")
			return nil
		}

		err = setTimezoneForUser(m.Author, args[0])
		if err != nil {
			return err
		}

		now := time.Now().In(loc)

		s.ChannelMessageSend(m.ChannelID, "Votre fuseau horaire est désormais **"+args[0]+"**, dans ce fuseau horaire, il est **"+now.String()+"**.")
	case "get":
		if len(args) < 1 {
			s.ChannelMessageSend(m.ChannelID, "Il manque le fuseau horaire. Utilisez la commande `help time` pour accéder à l’aide.")
		}

		loc, err := time.LoadLocation(args[0])
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Je ne connais pas ce fuseau horaire, désolé.")
			return nil
		}

		now := time.Now().In(loc)

		s.ChannelMessageSend(m.ChannelID, "Dans ce fuseau horaire, il est **"+now.String()+"**.")
	default:
		if len(m.Mentions) == 0 {
			s.ChannelMessageSend(m.ChannelID, "Veuillez mentionner au moins un utilisateur.")
		} else {
			for _, user := range m.Mentions {
				timezone, err := getTimezoneByUser(user)
				if err != nil {
					return err
				}

				if timezone == "" {
					s.ChannelMessageSend(m.ChannelID, "Je connais pas le fuseau horaire de **"+user.Username+"**.")
				} else {
					loc, err := time.LoadLocation(timezone)
					if err != nil {
						s.ChannelMessageSend(m.ChannelID, "Le fuseau horaire de **"+user.Username+"** est invalide.")
					}

					now := time.Now().In(loc)

					s.ChannelMessageSend(m.ChannelID, "Dans le fuseau horaire de **"+user.Username+"**, il est **"+now.String()+"**.")
				}
			}
		}
	}

	return nil
}
