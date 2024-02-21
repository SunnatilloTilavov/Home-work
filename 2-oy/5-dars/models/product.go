package models
import "github.com/google/uuid"

type Product struct {
	Id        uuid.UUID
	Name      string
	Code      int
	Category_id uuid.UUID
	CreatedAt string
	UpdatedAt string
}