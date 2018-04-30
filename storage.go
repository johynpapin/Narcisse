package main

import (
	"github.com/coreos/bbolt"
	"github.com/bwmarrin/discordgo"
	"fmt"
	"strconv"
)

var db *bolt.DB

func openStorage() error {
	var err error
	db, err = bolt.Open("narcisse.db", 0600, nil)
	if err != nil {
		return err
	}

	return createBuckets()
}

func closeStorage() error {
	return db.Close()
}

func createBuckets() error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("timezones"))
		if err != nil {
			return fmt.Errorf("error creating timezones bucket: %s", err)
		}

		_, err = tx.CreateBucketIfNotExists([]byte("counter"))
		if err != nil {
			return fmt.Errorf("error creating counter bucket: %s", err)
		}

		return nil
	})
}

func setTimezoneForUser(user *discordgo.User, timezone string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("timezones"))

		return b.Put([]byte(user.ID), []byte(timezone))
	})
}

func getTimezoneByUser(user *discordgo.User) (string, error) {
	var timezone string

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("timezones"))
		timezonebytes := b.Get([]byte(user.ID))

		if timezonebytes != nil {
			timezone = string(timezonebytes)
		}

		return nil
	})

	return timezone, err
}

func incrementCounter() (int, error) {
	counter := 1

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("counter"))
		v := b.Get([]byte("counter"))

		if v != nil {
			vi, err := strconv.Atoi(string(v))
			if err != nil {
				return err
			}

			counter = vi + 1
		}

		b.Put([]byte("counter"), []byte(strconv.Itoa(counter)))

		return nil
	})

	return counter, err
}

func decrementCounter() (int, error) {
	counter := -1

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("counter"))
		v := b.Get([]byte("counter"))

		if v != nil {
			vi, err := strconv.Atoi(string(v))
			if err != nil {
				return err
			}

			counter = vi - 1
		}

		b.Put([]byte("counter"), []byte(strconv.Itoa(counter)))

		return nil
	})

	return counter, err
}
