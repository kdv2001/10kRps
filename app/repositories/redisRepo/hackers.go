package redisRepo

import (
	"10kRps/app/models"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
)

type hackersRedis struct {
	client *redis.Client
}

func CreateHackersRedis(cl *redis.Client) *hackersRedis {
	return &hackersRedis{client: cl}
}

func (h *hackersRedis) GetAll(kitName string) ([]models.Hacker, error) {
	res := h.client.ZRangeWithScores(context.Background(), kitName, 0, -1)

	if res.Err() != nil {
		fmt.Println(res.Err())
		return nil, fiber.ErrInternalServerError
	}
	r, err := res.Result()
	if err != nil {
		fmt.Println(err)
		return nil, fiber.ErrInternalServerError
	}
	if len(r) == 0 {
		return nil, fiber.ErrNotFound
	}
	var result []models.Hacker
	for _, val := range r {
		result = append(result, models.Hacker{Name: fmt.Sprintf("%v", val.Member), Score: val.Score})
	}

	return result, nil
}
