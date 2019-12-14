package main

import (
	"fmt"

	"github.com/yoshimitsuEgashira/97programmerbot/api"
)

func main() {
	apiAuth := api.Auth()

	t, err := api.MakeText()
	if err != nil {
		fmt.Printf("Failed to make tweet text %s", err.Error())
	}
	if t == "" {
		fmt.Println("Tweet text is empty")
		return
	}

	tw, err := api.PostTweet(*apiAuth, "", nil)
	if err != nil {
		fmt.Printf("Tweet failed : %s", err.Error())
		return
	}
	fmt.Printf("tw %#v", tw)
}
