package network

import (
	"context"
	"github.com/ascendcorp/google_chat/git"
	"github.com/ascendcorp/google_chat/pull_request_message"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/chat/v1"
	"log"
)

func AlertNewPullRequest(bitbucket git.BitbucketPullRequest, eventKey string, room string, googleChatAuth *jwt.Config) {
	conf := googleChatAuth //auth()
	client := conf.Client(context.Background())

	log.Print("callChatAPI 2")

	myChat, err := chat.New(client)
	if err != nil {
		log.Fatal(err)
	}

	var message *chat.Message
	if eventKey == "pullrequest:created" {
		message = google_chat.PullRequestMessage(bitbucket, false)
	} else if eventKey == "pullrequest:updated"   {
		message = google_chat.PullRequestMessage(bitbucket, true)
	} else if eventKey == "pullrequest:approved" {
		message = google_chat.ApprovedPullRequestMessage(bitbucket)
	}

	if message != nil {
		createCall := myChat.Spaces.Messages.Create("spaces/" + room, message).ThreadKey(bitbucket.PullRequest.Links.HTML.Href)
		resp, err := createCall.Do()
		if err != nil {
			log.Fatal(err)
		} else {
			log.Print(resp)
		}
	}

}