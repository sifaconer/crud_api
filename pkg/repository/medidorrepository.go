package repository

import (
	"github.com/google/uuid"
	"github.com/sifaconer/crud_api/pkg/domain/entity"
)

type MedidorRepository interface {
	Create(model entity.Medidor) (entity.Medidor, error)
	Update(model entity.Medidor) (entity.Medidor, error)
	Delete(id uuid.UUID) error
	All() ([]entity.Medidor, error)
	ByID(id uuid.UUID) (entity.Medidor, error)
	BySerial(serial string) ([]entity.Medidor, error)
	RecentInstallation(serial string) (entity.Medidor, error)
	Inactive() ([]entity.Medidor, error) // condicionar para energía true o false
}
