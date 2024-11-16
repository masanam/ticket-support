package ticketmodel

import (
	"go-ticket-support/helpers/errorcodehandling"

	"github.com/jinzhu/gorm"
)

type TicketRepository struct {
	db        *gorm.DB
	codeError *errorcodehandling.CodeError
}

func NewTicketRepository(db *gorm.DB) *TicketRepository {
	return &TicketRepository{
		db: db,
	}
}
