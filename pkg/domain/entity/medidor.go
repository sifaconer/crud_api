package entity

import (
	"time"

	"github.com/google/uuid"
)

type Medidor struct {
	ID               uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Brand            string    `gorm:"not null"`
	Address          string    `gorm:"not null"`
	InstallationDate time.Time `gorm:"not null"`
	RetirementDate   *time.Time
	Serial           string    `gorm:"not null"`
	Lines            uint32    `gorm:"not null;check:lines <= 10"`
	IsActive         bool      `gorm:"not null"`
	CreatedAt        time.Time `gorm:"not null;default:now()"`
}
