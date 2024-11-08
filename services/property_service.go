package services

import (
	"github.com/mauriciomartinezc/real-estate-mc-property/domain"
	"github.com/mauriciomartinezc/real-estate-mc-property/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PropertyService struct {
	Repo *repository.PropertyRepository
}

func NewPropertyService(repo *repository.PropertyRepository) *PropertyService {
	return &PropertyService{Repo: repo}
}

func (s *PropertyService) GetAllPropertiesPaginated(page int, limit int) (domain.SimpleProperties, error) {
	return s.Repo.GetAllPropertiesPaginated(page, limit)
}

func (s *PropertyService) GetPropertiesByCompanyID(companyID string, page int, limit int) (domain.SimpleProperties, error) {
	return s.Repo.GetPropertiesByCompanyID(companyID, page, limit)
}

func (s *PropertyService) Create(property *domain.SimpleProperty) error {
	return s.Repo.Create(property)
}

func (s *PropertyService) Update(property *domain.SimpleProperty) error {
	return s.Repo.Update(property)
}

func (s *PropertyService) GetByID(id primitive.ObjectID) (*domain.SimpleProperty, error) {
	return s.Repo.GetByID(id)
}

func (s *PropertyService) GetDetailByID(id primitive.ObjectID) (*domain.DetailProperty, error) {
	return s.Repo.GetDetailByID(id)
}

func (s *PropertyService) ChangeStatus(property *domain.SimpleProperty) error {
	return s.Repo.ChangeStatus(property)
}
