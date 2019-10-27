package tweet_test

import (
	"testing"

	"github.com/nickrisaro/workshop-go/tweet"
	"github.com/stretchr/testify/assert"
)

func TestCanSaveAndRetrieveATweet(t *testing.T) {

	newTweet := tweet.New("@nickrisaro", "Testeando mi twitter")
	repository := tweet.NewRepository()
	repository.Save(newTweet)

	tweets := repository.GetAllTweets()

	assert.Contains(t, tweets, newTweet, "No se encontró el tweet en la lista de tweets")

}
