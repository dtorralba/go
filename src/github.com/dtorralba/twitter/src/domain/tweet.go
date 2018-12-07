package domain

import (
	"time"
)

type Stringer interface {
	String() string
}

type Tweet interface {
	GetUser() string
	GetText() string
	GetDate() *time.Time
	PrintableTweet() string
}

type TextTweet struct {
	User string
	Text string
	Date *time.Time
}

type ImageTweet struct {
	TextTweet
	URL string
}

type QuoteTweet struct {
	TextTweet
	QuoteText TextTweet
}

func NewImageTweet(user string, text string, URL string) *ImageTweet {
	textTweet := NewTextTweet(user, text)
	imageTweet := ImageTweet{textTweet, URL}
	return &imageTweet
}

func NewQuoteTweet(user string, text string, quote TextTweet) *QuoteTweet {
	textTweet := NewTextTweet(user, text)
	quoteTweet := QuoteTweet{textTweet, quote}
	return &quoteTweet
}

func NewTextTweet(user string, text string) TextTweet {
	var time = time.Now()
	textTweet := TextTweet{user, text, &time}
	return textTweet
}

func (textTweet *TextTweet) PrintableTweet() string {
	return "@" + textTweet.GetUser() + ": " + textTweet.GetText()
}

func (imageTweet *ImageTweet) PrintableTweet() string {
	return "@" + imageTweet.GetUser() + ": " + imageTweet.GetText() + " " + imageTweet.GetURL()
}

func (quoteTweet *QuoteTweet) PrintableTweet() string {
	return "@" + quoteTweet.GetQuoteName() + ": " + quoteTweet.GetText() + " \"" + quoteTweet.GetQuoteText() + "\""
}

func (textTweet *TextTweet) String() string {
	return "@" + textTweet.GetUser() + ": " + textTweet.GetText()
}

func (tweet TextTweet) GetUser() string {
	return tweet.User
}

func (tweet TextTweet) GetText() string {
	return tweet.Text
}

func (tweet TextTweet) GetDate() *time.Time {
	return tweet.Date
}

func (imageTweet ImageTweet) GetURL() string {
	return imageTweet.URL
}

func (quoteTweet QuoteTweet) GetQuoteName() string {
	return quoteTweet.User
}

func (quoteTweet QuoteTweet) GetQuoteText() string {

	return quoteTweet.QuoteText.PrintableTweet()
}
