package service_test

import (
	"testing"

	"github.com/dtorralba/twitter/src/service"

	"github.com/dtorralba/twitter/src/domain"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	//Initialization
	tweetManager := service.NewTweetManager()
	var tweet *domain.Tweet

	//var tweet string = "This is my first tweet"
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text) //dfevuelve un puntero a ese tw

	//Operation
	tweetManager.PublishTweet(tweet)

	//Validation
	publishedTweet := service.GetTweet()
	if publishedTweet.User != user &&
		publishedTweet.Text != text {
		t.Errorf("Expected tweet is %s: %s \nbut is %s: %s",
			user, text, publishedTweet.User, publishedTweet.Text)
	}
	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}
}
func TestTweetWithoutUserIsNotPublished(t *testing.T) {

	//Initialization
	tweetManager := service.NewTweetManager()
	var tweet *domain.Tweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	//Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWhitoutTextIsNotPublished(t *testing.T) {
	//Initialization
	tweetManager := service.NewTweetManager()
	var tweet *domain.Tweet

	var text string
	user := "Diego"

	tweet = domain.NewTweet(user, text)

	//Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	//Validation
	if err == nil {
		t.Error("error is expected")
		return
	}
	if err != nil && err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}

}
func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	//Initialization
	tweetManager := service.NewTweetManager()
	var tweet *domain.Tweet

	user := "Diego"
	text := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book."

	tweet = domain.NewTweet(user, text)

	//Operation
	var err error
	_, err = tweetManager.PublishTweet(tweet)

	//Validation
	if err != nil && err.Error() != "text exceed 140 characters" {
		t.Error("Expected error is text exceed 140 characters")
	}
}

func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	id := 0
	// Initialization
	tweetManager := service.NewTweetManager()
	var tweet, secondTweet *domain.Tweet // Fill the tweets with data

	user := "Diego"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, text)

	// Operation
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)

	// Validation
	publishedTweets := tweetManager.GetTweets()

	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}
	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]
	if !isValidTweet(t, firstPublishedTweet, id, user, text) {
		return
	}

	// Same for secondPublishedTweet
	if !isValidTweet(t, secondPublishedTweet, id, user, text) {
		return
	}
}

func TestCanRetrieveTweetById(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	//service.InitializeService()

	var tweet *domain.Tweet
	var id int

	user := "grupoesfera"
	text := "This is my first tweet"

	tweet = domain.NewTweet(user, text)

	// Operation
	id, _ = tweetManager.PublishTweet(tweet)

	// Validation
	publishedTweet, _ := tweetManager.GetTweetById(id)

	isValidTweet(t, publishedTweet, id, user, text)
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	// Initialization
	tweetManager := service.NewTweetManager()

	//service.InitializeService()

	var tweet, secondTweet, thirdTweet *domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)
	//RetrieveTheTweetsSentByAnUser
	// Operation
	count := tweetManager.CountTweetsByUser(user)

	// Validation
	if count != 2 {
		t.Errorf("Expected count is 2 but was %d", count)
	}
}

func TestCan(t *testing.T) {

	// Initialization
	tweetManager := service.NewTweetManager()

	//service.InitializeService()
	var id int

	var tweet, secondTweet, thirdTweet *domain.Tweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTweet(user, text)
	secondTweet = domain.NewTweet(user, secondText)
	thirdTweet = domain.NewTweet(anotherUser, text)

	// publish the 3 tweets
	tweetManager.PublishTweet(tweet)
	tweetManager.PublishTweet(secondTweet)
	tweetManager.PublishTweet(thirdTweet)

	// Operation
	tweets := tweetManager.GetTweetsByUser(user)

	// Validation
	if len(tweets) != 2 { /* handle error */
	}

	firstPublishedTweet := tweets[0]
	secondPublishedTweet := tweets[1]

	isValidTweet(t, firstPublishedTweet, id, user, text)
	isValidTweet(t, secondPublishedTweet, id, user, secondText)
}

func isValidTweet(t *testing.T, firstPublishedTweet domain.Tweet, id int, user string, text string) bool {
	if firstPublishedTweet.Text == text {
		if firstPublishedTweet.User == user {
			return true
		}
	}
	return false
}
