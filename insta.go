package main

import (
	"errors"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	ins "github.com/ahmdrz/goinsta"
	"github.com/spf13/viper"
	"github.com/tducasse/goinsta"
	"github.com/tducasse/goinsta/store"
)

// Insta is a goinsta.Instagram instance
var insta *goinsta.Instagram

// login will try to reload a previous session, and will create a new one if it can't
func login() {
	err := reloadSession()
	if err != nil {
		createAndSaveSession()
	}
}

func instagramPost() {

	now := time.Now()
	after := now.AddDate(0, -2, 0)

	User, err := ins.Profiles.ByName("icata")
	check(err)

	media := User.Feed(after.Unix())

	media.Sync()
	media.Next()
}

// Logins and saves the session
func createAndSaveSession() {
	insta = goinsta.New(viper.GetString("user.instagram.username"), viper.GetString("user.instagram.password"))
	err := insta.Login()
	check(err)

	key := createKey()
	bytes, err := store.Export(insta, key)

	check(err)
	err = ioutil.WriteFile("session", bytes, 0644)
	check(err)
	log.Println("Created and saved the session")
}

// reloadSession will attempt to recover a previous session
func reloadSession() error {
	if _, err := os.Stat("session"); os.IsNotExist(err) {
		return errors.New("No session found")
	}

	session, err := ioutil.ReadFile("session")
	check(err)
	log.Println("A session file exists")

	key, err := ioutil.ReadFile("key")
	check(err)

	insta, err = store.Import(session, key)
	if err != nil {
		return errors.New("Couldn't recover the session")
	}

	log.Println("Successfully logged in")
	return nil

}

// createKey creates a key and saves it to file
func createKey() []byte {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	check(err)
	err = ioutil.WriteFile("key", key, 0644)
	check(err)
	log.Println("Created and saved the key")
	return key
}
