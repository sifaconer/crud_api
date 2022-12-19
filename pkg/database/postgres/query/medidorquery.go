package query

import (
	"github.com/google/uuid"
	"github.com/sifaconer/crud_api/pkg/domain/entity"
	"gorm.io/gorm"
)

// implements repository.MedidorRepository
type medidorQuery struct {
	DB *gorm.DB
}

func NewMedidorQuery(db *gorm.DB) *medidorQuery {
	return &medidorQuery{
		DB: db,
	}
}

func (s *medidorQuery) Create(model entity.Medidor) (entity.Medidor, error) {
	result := s.DB.Create(&model)
	if result.Error != nil {
		return entity.Medidor{}, result.Error
	}

	return model, nil
}

func (s *medidorQuery) Update(model entity.Medidor) (entity.Medidor, error) {
	var updated entity.Medidor

	result := s.DB.First(&updated, model.ID)
	if result.Error != nil {
		return entity.Medidor{}, result.Error
	}

	if updated.Address != model.Address {
		updated.Address = model.Address
	}
	if updated.RetirementDate != model.RetirementDate {
		updated.RetirementDate = model.RetirementDate
	}
	if updated.Lines != model.Lines {
		updated.Lines = model.Lines
	}
	if updated.IsActive != model.IsActive {
		updated.IsActive = model.IsActive
	}

	result = s.DB.Save(&updated)
	if result.Error != nil {
		return entity.Medidor{}, result.Error
	}

	return updated, nil
}

func (s *medidorQuery) Delete(id uuid.UUID) error {

	result := s.DB.Delete(&entity.Medidor{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *medidorQuery) All() ([]entity.Medidor, error) {
	var resp []entity.Medidor
	// Get all records
	result := s.DB.Find(&resp)
	if result.Error != nil {
		return resp, result.Error
	}

	return resp, nil
}

func (s *medidorQuery) RecentInstallation(serial string) (entity.Medidor, error) {
	var model entity.Medidor
	result := s.DB.Where("serial = ?", serial).Order("installation_date desc").Limit(1).Find(&model)
	if result.Error != nil {
		return model, result.Error
	}

	return model, nil
}

func (s *medidorQuery) ByID(id uuid.UUID) (entity.Medidor, error) {
	var resp entity.Medidor

	result := s.DB.First(&resp, id)
	if result.Error != nil {
		return resp, result.Error
	}

	return resp, nil
}

func (s *medidorQuery) Inactive() ([]entity.Medidor, error) {
	var model []entity.Medidor
	result := s.DB.Where("is_active = ?", false).Find(&model)
	if result.Error != nil {
		return model, result.Error
	}

	return model, nil
}

func (s *medidorQuery) BySerial(serial string) ([]entity.Medidor, error) {
	var model []entity.Medidor
	result := s.DB.Where("serial = ?", serial).Find(&model)
	if result.Error != nil {
		return model, result.Error
	}

	return model, nil
}
