package repositories

import "github.com/jinzhu/gorm"

type Repository struct {
	Ticket Ticket
}

// NewRepository to setting services repositories
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Ticket: NewTicket(db),
	}
}
