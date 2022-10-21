package main

import (
	"context"
	"fmt"

	"github.com/beeper/groupme-lib"
)

// This is not a real Bot ID. Please find yours by logging
// into the GroupMe development website: https://dev.groupme.com/bots
const botID = "0123456789ABCDEF"

// A short program that gets the gets the first 5 groups
// the user is part of, and then the first 10 messages of
// the first group in that list
func main() {
	// Create a new client with your auth token
	client := groupme.NewClient("")
	fmt.Println(client.PostBotMessage(context.Background(), botID, "Your message here!", nil))
}
