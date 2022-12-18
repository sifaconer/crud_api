package entity

import (
	"time"

	"github.com/google/uuid"
)

type Medidor struct {
	ID               uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Brand            string
	PredioID         uuid.UUID
	Predio           Predio `gorm:"foreignKey:PredioID"`
	InstallationDate time.Time
	RetirementDate   *time.Time
	SerialID         uuid.UUID
	Serial           Serial `gorm:"foreignKey:SerialID"`
	Lines            uint8  `gorm:"check:lines <= 10"` //Int (valores permitidos: 0 a 10)	False	Número de líneas conectadas, puede ir de 1 a 10
	IsActive         bool
	CreatedAt        time.Time
}
