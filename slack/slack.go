package slack

import (
	"fmt"
	"github.com/slack-go/slack"
)

const TOKEN = "xoxb-1984448532487-3392524779652-QcW8cVAbs1Qjc513cEh9QQmg"
const ChannelID = "github-notification"

func SendMessageToSlack() error {
	api := slack.New(TOKEN)

	attachment := slack.Attachment{
		Pretext: "테스트 입력입니다.",
		Text:    "테스트 두번째 입력입니다.",
	}

	channelID, timestamp, err := api.PostMessage(
		ChannelID,
		slack.MsgOptionText("", false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(false), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return nil
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
	return nil
}
