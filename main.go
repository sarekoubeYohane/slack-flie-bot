package main

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	os.Setenv("SLACK_BOT_TOKEN", "xoxb-4392660304743-4431384157168-5PV9C5pKwZhSa0Zp8Sm25pbd")
	os.Setenv("CHANNEL_ID", "C04CNSMA0DN")

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"text.txt"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}
