package ticket

import (
	"context"

	"go-ticket-support/entity/ticketentity"
)

type Service interface {
	InsertNewTicket(ctx context.Context, payload *ticketentity.TicketRequest) (*ticketentity.Ticket, error)
	FindTicket(ctx context.Context, ticketID string) (*ticketentity.FindTicket, error)
	GetListTickets(ctx context.Context, sortBy, orderBy string, perPage, page int, filterName, filterType, filterValue string) (*[]ticketentity.Tickets, int64, error)
	UpdateTicket(ctx context.Context, payload *ticketentity.TicketData) (*ticketentity.TicketData, error)
	DeleteDataTicket(ctx context.Context, ticketID string) error
}
