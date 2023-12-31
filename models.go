package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/jegorie/rss-server/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"apiKey"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	OwnerId   uuid.UUID `json:"ownerId"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		OwnerId:   dbFeed.OwnerID,
	}
}

func databaseAllFeedToAllFeed(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}

	for _, dbFeedsItem := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeedsItem))
	}

	return feeds
}

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	OwnerId   uuid.UUID `json:"ownerId"`
	FeedId    uuid.UUID `json:"feedId"`
}

func databaseFeedFoolowToFeedFollow(dbFeedFollow database.FeedFollow) FeedFollow {
	return FeedFollow{
		ID:        dbFeedFollow.ID,
		CreatedAt: dbFeedFollow.CreatedAt,
		UpdatedAt: dbFeedFollow.UpdatedAt,
		FeedId:    dbFeedFollow.FeedID,
		OwnerId:   dbFeedFollow.OwnerID,
	}
}

func databaseFeedFollowsToFeedFollows(dbFeeds []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}

	for _, dbFeedFollowItem := range dbFeeds {
		feedFollows = append(feedFollows, databaseFeedFoolowToFeedFollow(dbFeedFollowItem))
	}

	return feedFollows
}
