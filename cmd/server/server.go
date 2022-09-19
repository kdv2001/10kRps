package server

import (
	"10kRps/app/handlers"
	"10kRps/app/repositories/redisRepo"
	"10kRps/app/usecases/impl"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

type Server struct {
	hackerHandlers handlers.HackersHandler
}

func CreateServer() Server {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_CONTAINER"),
		Password: "",
		DB:       0,
	})
	hackerHandlers := handlers.CreateHackersHandler(impl.CreateNewHackersUseCases(redisRepo.CreateHackersRedis(redisClient)))
	return Server{hackerHandlers: hackerHandlers}
}

func (serv *Server) Start() {
	app := fiber.New()
	app.Get("/json/:group?", serv.hackerHandlers.Get)
	log.Fatal(app.Listen(os.Getenv("LISTEN_PORT")))
}
