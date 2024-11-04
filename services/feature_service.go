package services

import (
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"github.com/mauriciomartinezc/real-estate-mc-property/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FeatureService struct {
	Repo *repository.FeatureRepository
}

func NewFeatureService(repo *repository.FeatureRepository) *FeatureService {
	return &FeatureService{Repo: repo}
}

func (s *FeatureService) GetAll() (features domain.Features, err error) {
	return s.Repo.GetAll()
}

func (s *FeatureService) GetFeaturesGroupedByType() (map[string][]domain.Feature, error) {
	return s.Repo.GetFeaturesGroupedByType()
}

func (s *FeatureService) Create(feature *domain.Feature) error {
	return s.Repo.Create(feature)
}

func (s *FeatureService) GetByID(id primitive.ObjectID) (*domain.Feature, error) {
	return s.Repo.GetByID(id)
}

func (s *FeatureService) Update(feature *domain.Feature) error {
	return s.Repo.Update(feature)
}

func (s *FeatureService) Delete(id primitive.ObjectID) error {
	return s.Repo.Delete(id)
}
