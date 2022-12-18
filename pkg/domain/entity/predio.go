package entity

import (
	"time"

	"github.com/google/uuid"
)

type Predio struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Address   string
	CreatedAt time.Time
}
