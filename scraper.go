package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/karma-234/blog-post/internal/database"
)

func startScarping(db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration) {
	log.Printf("Fetching on %v goroutines every %s duration", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil {
			log.Printf("Error fecthing rss feed: %v", err)
			continue
		}
		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}

}
func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, dbFeed database.Feed) {
	defer wg.Done()
	_, err := db.MarkFeedAsfetched(context.Background(), dbFeed.ID)
	if err != nil {
		log.Printf("Error markling feed as fetched: %v", err)
		return
	}
	rssFeed, err := urlTofeed(dbFeed.Url)
	if err != nil {
		log.Printf("Error converting url to feed: %v", err)
		return
	}
	for _, feed := range rssFeed.Channel.Items {
		log.Printf("Found post in feed: %v", feed.Title)
	}
	log.Printf("Feed %v collected, %v posts found", dbFeed.Name, len(rssFeed.Channel.Items))
}
