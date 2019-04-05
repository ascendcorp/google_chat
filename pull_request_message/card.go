package google_chat

import (
	"github.com/ascendcorp/google_chat/git"
	"google.golang.org/api/chat/v1"
)

func ApprovedPullRequestMessage(bitbucket git.BitbucketPullRequest) *chat.Message {
	var approvedPullRequestCard = approvedPullRequestCard(
		bitbucket.Approval.User.DisplayName,
		"Just approved this pull request",
		bitbucket.Approval.User.Links.Avatar.Href)
	var cards  = make([]*chat.Card, 1)
	cards = append(cards, approvedPullRequestCard)
	var (
		message  = &chat.Message {
			Text: "<users/all>",
			Cards: cards,
		}
	)
	return message
}

func PullRequestMessage(bitbucket git.BitbucketPullRequest, isUpdate bool) *chat.Message  {
	var subtitle string
	if isUpdate {
		subtitle = "Just update a pull request"
	} else {
		subtitle = "Just create a pull request"
		// card.Header = createPullRequestHeader(bitbucket.PullRequest.Author.DisplayName, bitbucket.PullRequest.Author.Links.Avatar.Href)
	}

	var pullRequestChatCard = pullRequestCard(
		bitbucket.PullRequest.Author.DisplayName,
		subtitle,
		bitbucket.PullRequest.Author.Links.Avatar.Href,
		bitbucket.PullRequest.Title,
		bitbucket.PullRequest.Summary.HTML,
		bitbucket.PullRequest.Links.HTML.Href)

	var cards  = make([]*chat.Card, 1)
	cards = append(cards, pullRequestChatCard)
	var (
		message  = &chat.Message {
			Text: "<users/all>",
			Cards: cards,
		}
	)

	return message
}

func approvedPullRequestCard(title string, subTitle string, avatarUrl string) *chat.Card  {

	var card = new(chat.Card)
	card.Header = &chat.CardHeader {
		Title: title,
		Subtitle: subTitle,
		ImageUrl: avatarUrl,
		ImageStyle: "AVATAR",
	}
	return card
}

func pullRequestCard(title string, subTitle string, avatarUrl string, contentTitle string, contentDesc string, link string) *chat.Card {

	var descSection = pullRequestDescription(contentTitle, contentDesc)
	var buttonSection = pullRequestReviewButton(link)

	var cardContent = make([]*chat.Section, 1)
	cardContent = append(cardContent, descSection)
	cardContent = append(cardContent, buttonSection)

	var card = new(chat.Card)
	card.Header = header(title, subTitle, avatarUrl)
	card.Sections = cardContent
	return card
}

func header(title string, subTitle string, avatarImgUrl string) *chat.CardHeader {
	var header = &chat.CardHeader {
		Title: title,
		Subtitle: subTitle,
		ImageUrl: avatarImgUrl,
		ImageStyle: "AVATAR",
	}
	return header
}

func pullRequestDescription(title string, desc string) *chat.Section  {

	var widgets = make([]* chat.WidgetMarkup, 2)
	var titleWidget = &chat.WidgetMarkup {
		KeyValue: &chat.KeyValue {
			TopLabel: "Title",
			Content: title,
		},
	}
	var descWidget  = &chat.WidgetMarkup {
		TextParagraph: &chat.TextParagraph {
			Text: desc,
		},
	}

	widgets = append(widgets, titleWidget)
	widgets = append(widgets, descWidget)

	var prDescSection = new(chat.Section)
	prDescSection.Widgets = widgets

	return prDescSection
}

func pullRequestReviewButton(link string) *chat.Section {
	var buttons = make([]* chat.Button, 1)
	var button = &chat.Button {
		TextButton: &chat.TextButton {
			Text: "REVIEW",
			OnClick: &chat.OnClick {
				OpenLink: &chat.OpenLink {
					Url: link,
				},
			},
		},
	}
	buttons = append(buttons, button)

	var widgets = make([]* chat.WidgetMarkup, 1)
	var widget = &chat.WidgetMarkup {
		Buttons: buttons,
	}
	widgets = append(widgets, widget)

	var buttonSection = new(chat.Section)
	buttonSection.Widgets = widgets

	return buttonSection
}