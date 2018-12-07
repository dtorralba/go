package service

import (
	"fmt"

	"github.com/dtorralba/twitter/src/domain"
)

type TweetManager struct {
	SliceTweets []domain.Tweet
}

var emptyTweet domain.Tweet
var NewTweet domain.Tweet

var SliceTweetsByUser []domain.Tweet
var id int

func NewTweetManager() *TweetManager {
	var v TweetManager
	v.SliceTweets = make([]domain.Tweet, 0)
	return &v
}

func (tweetManager *TweetManager) PublishTweet(tweet domain.Tweet) (int, error) {

	if tweet.GetUser() == "" {
		return 0, fmt.Errorf("user is required")
	}

	if tweet.GetText() == "" {
		return 0, fmt.Errorf("text is required")
	}

	if len(tweet.GetText()) > 140 {
		return 0, fmt.Errorf("text exceed 140 characters")
	}

	NewTweet = tweet

	tweetManager.SliceTweets = append(tweetManager.SliceTweets, NewTweet)

	//Id por posicion en Slice
	var id = len(tweetManager.SliceTweets) - 1
	return id, nil
}

func (tweetManager *TweetManager) GetTweet() domain.Tweet {
	return NewTweet
}

func (tweetManager *TweetManager) InitializeService() {
	//limpiar tweets
	tweetManager.SliceTweets = tweetManager.SliceTweets[:0]
}

func (tweetManager *TweetManager) GetTweets() []domain.Tweet {
	return tweetManager.SliceTweets
}

func (tweetManager *TweetManager) GetTweetById(id int) (domain.Tweet, string) {
	length := len(tweetManager.SliceTweets)

	if id < length && id >= 0 {
		return tweetManager.SliceTweets[id], ""
	}
	return emptyTweet, "Id Invalido"
}

func (tweetManager *TweetManager) CountTweetsByUser(user string) int {
	count := 0
	for _, valor := range tweetManager.SliceTweets {
		if valor.GetUser() == user {
			count++
		}
	}

	return count
}

func (tweetManager *TweetManager) GetTweetsByUser(user string) []domain.Tweet {

	var tweetsByUser map[string][]domain.Tweet
	tweetsByUser = make(map[string][]domain.Tweet)

	for _, valor := range tweetManager.SliceTweets {
		if valor.GetUser() == user {
			SliceTweetsByUser = append(SliceTweetsByUser, valor)
		}
	}

	tweetsByUser[user] = SliceTweetsByUser

	println(tweetsByUser)
	return tweetsByUser[user]
}
