package services

import (
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"github.com/mauriciomartinezc/real-estate-mc-property/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeatureTypeService struct {
	Repo *repositories.FeatureTypeRepository
}

func NewFeatureTypeService(repo *repositories.FeatureTypeRepository) *FeatureTypeService {
	return &FeatureTypeService{Repo: repo}
}

func (s *FeatureTypeService) GetAll() (featureTypes domain.FeatureTypes, err error) {
	return s.Repo.GetAll()
}

func (s *FeatureTypeService) Create(featureType *domain.FeatureType) error {
	return s.Repo.Create(featureType)
}

func (s *FeatureTypeService) GetByID(id primitive.ObjectID) (*domain.FeatureType, error) {
	return s.Repo.GetByID(id)
}

func (s *FeatureTypeService) Update(featureType *domain.FeatureType) error {
	return s.Repo.Update(featureType)
}

func (s *FeatureTypeService) Delete(id primitive.ObjectID) error {
	return s.Repo.Delete(id)
}
