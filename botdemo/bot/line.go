package main

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
)

type Line struct {
	ChannelSecret string
	ChannelToken  string
	Client        *linebot.Client
}

func (c *Line) New(secret, token string) error {
	c.ChannelSecret = secret
	c.ChannelToken = token

	bot, err := linebot.New(
		c.ChannelSecret,
		c.ChannelToken,
	)
	if err != nil {
		return err
	}
	c.Client = bot
	return nil
}

func (r *Line) handleText(message *linebot.TextMessage, replyToken, userID string) {
	r.SendTextMessage(message.Text, replyToken)
}

func (r *Line) SendTextMessage(message string, replyToken string) error {
	return r.Reply(replyToken, linebot.NewTextMessage(message))
}

func (r *Line) EventRouter(eve []*linebot.Event) {
	for _, event := range eve {
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if message.Text == "今日の運勢" {
					message.Text = "最高の1日になるでしょう"
					r.handleText(message, event.ReplyToken, event.Source.UserID)
				}
			}
		}
	}
}

func (r *Line) SendTemplateMessage(replyToken, altText string, template linebot.Template) error {
	return r.Reply(replyToken, linebot.NewTemplateMessage(altText, template))
}

func (r *Line) Reply(replyToken string, message linebot.SendingMessage) error {
	if _, err := r.Client.ReplyMessage(replyToken, message).Do(); err != nil {
		log.Printf("Reply Error: %v", err)
		return err
	}
	return nil
}

func (r *Line) NewCarouselColumn(thumbnailImageURL, title, text string, actions ...linebot.TemplateAction) *linebot.CarouselColumn {
	return &linebot.CarouselColumn{
		ThumbnailImageURL: thumbnailImageURL,
		Title:             title,
		Text:              text,
		Actions:           actions,
	}
}

func (r *Line) NewCarouselTemplate(columns ...*linebot.CarouselColumn) *linebot.CarouselTemplate {
	return &linebot.CarouselTemplate{
		Columns: columns,
	}
}
