package entity

import "github.com/google/uuid"

type Serial struct {
	ID           uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Code         string
	FabricanteID uuid.UUID
	Fabricante   Fabricante `gorm:"foreignKey:FabricanteID"`
}
