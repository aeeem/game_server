package redis_repository

import (
	"game_server/domain"

	"github.com/go-redis/redis"
)

type stateRedisRepository struct {
	Conn *redis.Client
}

func NewStateRedisRepository(Conn *redis.Client) (SR domain.StateRepository) {

	return
}
