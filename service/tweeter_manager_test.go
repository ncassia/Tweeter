package service_test

import (
	"github.com/ncassia/Tweeter/domain"
	"github.com/ncassia/Tweeter/service"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {

	// Initialization
	var tweet *domain.Tweet
	user := "grupoesfera"
	text := "This is my first tweet"
	tweet = domain.NewTweet(user, text)

	// Operation
	service.PublishTweet(tweet)

	// Validation
	publishedTweet := service.GetTweet()
	assert.Equal(t, publishedTweet, *tweet, "Tweets must be equals")
	assert.Equal(t, publishedTweet.User, user, "The user must be the same")
	assert.Equal(t, publishedTweet.Text, text, "The text must be the same")
	assert.NotNil(t, publishedTweet, "The Date is nil")

}
