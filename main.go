package main

import (
	"fmt"

	"github.com/dmitryk-dk/slackbot/api"
	"github.com/dmitryk-dk/slackbot/client"
	"log"
)

func main() {
	slackResponse, err := api.SlackConnect()
	if err != nil {
		log.Fatal(err)
	}

	clientNew := client.NewClient(slackResponse.Token)
	if err := clientNew.Connect(slackResponse.URL); err != nil {
		log.Fatal(err)
	}

	clientNew.Loop()

	for {
		select {
		case err := <-clientNew.Errors:
			log.Fatal(err)
		case msg := <-clientNew.Incoming:
			parse(clientNew, msg)
		}
	}
}

func parse(client *client.Client, msg interface{}) {
	switch msg := msg.(type) {
	case api.Hello:
		fmt.Println("Slack says hello!")

	case api.Message:
		channel := msg.Channel
		text := "Simple text message here"
		if err := client.SendMessage(channel, text); err != nil {
			fmt.Println("An error occured while responding", err)
		}

	default:
		fmt.Println("event received", msg)
	}

}
