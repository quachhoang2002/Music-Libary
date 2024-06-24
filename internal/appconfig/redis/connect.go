package redis

import (
	"fmt"
	"log"

	"github.com/quachhoang2002/Music-Library/config"
	"github.com/quachhoang2002/Music-Library/pkg/redis"
)

func Connect(redisConfig config.RedisConfig) (redis.Client, error) {
	redisOptions := redis.NewClientOptions().SetOptions(redisConfig)

	client, err := redis.Connect(redisOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	log.Println("Connected to Redis!")

	return client, nil
}

// Disconnect disconnects from the database.
func Disconnect(client redis.Client) {
	if client == nil {
		return
	}

	client.Disconnect()

	log.Println("Connection to MongoDB closed.")
}
