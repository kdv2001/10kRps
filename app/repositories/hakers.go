package repositories

import "10kRps/app/models"

type HackersRepository interface {
	GetAll(string) ([]models.Hacker, error)
}
