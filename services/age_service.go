package services

import (
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"github.com/mauriciomartinezc/real-estate-mc-property/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AgeService struct {
	Repo *repository.AgeRepository
}

func NewAgeService(repo *repository.AgeRepository) *AgeService {
	return &AgeService{Repo: repo}
}

func (s *AgeService) GetAll() (managementTypes domain.Ages, err error) {
	return s.Repo.GetAll()
}

func (s *AgeService) Create(managementType *domain.Age) error {
	return s.Repo.Create(managementType)
}

func (s *AgeService) GetByID(id primitive.ObjectID) (*domain.Age, error) {
	return s.Repo.GetByID(id)
}

func (s *AgeService) Update(managementType *domain.Age) error {
	return s.Repo.Update(managementType)
}

func (s *AgeService) Delete(id primitive.ObjectID) error {
	return s.Repo.Delete(id)
}
