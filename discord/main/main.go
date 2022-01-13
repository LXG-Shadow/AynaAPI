package main

import (
	"AynaAPI/discord"
	"fmt"
	"log"
	"os"
	"os/signal"
)

func main() {
	session, err := discord.CreateBot("")
	if err != nil {
		return
	}
	defer session.Close()
	err = session.Open()
	if err != nil {
		fmt.Println("cannot open the session")
		return
	}
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt) //nolint: staticcheck
	<-stop
	log.Println("Gracefully shutting down")
}
