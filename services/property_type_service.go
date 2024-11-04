package services

import (
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"github.com/mauriciomartinezc/real-estate-mc-property/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PropertyTypeService struct {
	Repo *repository.PropertyTypeRepository
}

func NewPropertyTypeService(repo *repository.PropertyTypeRepository) *PropertyTypeService {
	return &PropertyTypeService{Repo: repo}
}

func (s *PropertyTypeService) GetAll() (propertyTypes domain.PropertyTypes, err error) {
	return s.Repo.GetAll()
}

func (s *PropertyTypeService) Create(propertyType *domain.PropertyType) error {
	return s.Repo.Create(propertyType)
}

func (s *PropertyTypeService) GetByID(id primitive.ObjectID) (*domain.PropertyType, error) {
	return s.Repo.GetByID(id)
}

func (s *PropertyTypeService) Update(propertyType *domain.PropertyType) error {
	return s.Repo.Update(propertyType)
}

func (s *PropertyTypeService) Delete(id primitive.ObjectID) error {
	return s.Repo.Delete(id)
}
