package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	log "github.com/sirupsen/logrus"
	"github.com/turnage/graw/reddit"
)

var (
	rb reddit.Bot
)

func main() {
	loadConfig()

	openStorage()
	defer closeStorage()

	dg, err := connectDiscord()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("error connecting to Discord")
	}

	/*rb, err = connectReddit()
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Fatal("error connecting to Reddit")
	}*/

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
