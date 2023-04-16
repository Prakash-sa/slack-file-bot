package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", "xoxb-5110959150006-5103108723319-pcLu72B39GavvCxa5ojZQ31d")
	os.Setenv("CHANNEL_ID", "C053J15K8NQ")
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"eHills.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Println("%s\n", err)
			return
		}
		fmt.Println("Name %s, URL:%s\n", file.Name, file.URL)
	}
}
