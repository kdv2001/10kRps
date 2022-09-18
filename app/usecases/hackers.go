package usecases

import "10kRps/app/models"

type HackersUseCases interface {
	GetAllHackers(string) ([]models.Hacker, error)
}
