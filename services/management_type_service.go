package services

import (
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"github.com/mauriciomartinezc/real-estate-mc-property/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ManagementTypeService struct {
	Repo *repositories.ManagementTypeRepository
}

func NewManagementTypeService(repo *repositories.ManagementTypeRepository) *ManagementTypeService {
	return &ManagementTypeService{Repo: repo}
}

func (s *ManagementTypeService) GetAll() (managementTypes domain.ManagementTypes, err error) {
	return s.Repo.GetAll()
}

func (s *ManagementTypeService) Create(managementType *domain.ManagementType) error {
	return s.Repo.Create(managementType)
}

func (s *ManagementTypeService) GetByID(id primitive.ObjectID) (*domain.ManagementType, error) {
	return s.Repo.GetByID(id)
}

func (s *ManagementTypeService) Update(managementType *domain.ManagementType) error {
	return s.Repo.Update(managementType)
}

func (s *ManagementTypeService) Delete(id primitive.ObjectID) error {
	return s.Repo.Delete(id)
}
