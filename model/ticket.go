package repositories

import (
	"context"

	"go-ticket-support/entity/ticketentity"
	"go-ticket-support/model/ticketmodel"

	"github.com/jinzhu/gorm"
)

type Ticket interface {
	SaveNewTicket(ctx context.Context, payload *ticketentity.Ticket) error
	FindTicketByID(ctx context.Context, ticketID string) (*ticketentity.Ticket, error)
	ListTicket(sortBy, orderBy string, perPage, page int, filterName, filterType, filterValue string) (*[]ticketentity.Tickets, int64, error)
	UpdateTicketData(ctx context.Context, ticketID string, payload *ticketentity.Ticket) (*ticketentity.TicketData, error)
	DeleteTicket(ctx context.Context, ticketID string) error
}

func NewTicket(db *gorm.DB) Ticket {
	return ticketmodel.NewTicketRepository(db)
}
