package google_chat

import "google.golang.org/api/chat/v1"

func PullRequestCard(title string, subTitle string, avatarUrl string, contentTitle string, contentDesc string, link string) *chat.Card {

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