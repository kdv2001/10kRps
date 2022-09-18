package impl

import (
	"10kRps/app/models"
	"10kRps/app/repositories"
)

type hackersUseCases struct {
	repo repositories.HackersRepository
}

func CreateNewHackersUseCases(repo repositories.HackersRepository) *hackersUseCases {
	return &hackersUseCases{repo: repo}
}

func (impl *hackersUseCases) GetAllHackers(groupName string) ([]models.Hacker, error) {
	result, err := impl.repo.GetAll(groupName)
	if err != nil {
		return []models.Hacker{}, err
	}
	return result, nil
}
