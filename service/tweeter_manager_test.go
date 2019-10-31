package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPublishedTweetIsSaved(t *testing.T) {
	var tweet string = "This is my first tweet"

	PublishTweet(tweet)

//	if GetTweet() != tweet {
//		t.Error("Expected tweet is", tweet)
//	}
	assert.Equal(t,GetTweet(),tweet,"Should be the first Tweet")

}
