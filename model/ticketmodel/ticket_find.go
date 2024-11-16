package ticketmodel

import (
	"context"

	"go-ticket-support/config/app"
	"go-ticket-support/entity/ticketentity"

	"github.com/pkg/errors"
)

func (r *TicketRepository) FindTicketByID(ctx context.Context, ticketID string) (*ticketentity.Ticket, error) {
	var ticket ticketentity.Ticket
	err := r.db.First(&ticket, "id = ?", ticketID).Error
	if err != nil {
		parsed := r.codeError.ParseSQLError(err)
		switch parsed {
		case app.ErrNoRowsFound:
			return nil, app.ErrNoRowsFound
		case app.ErrUniqueViolation:
			return nil, app.ErrUniqueViolation
		default:
			return nil, errors.Wrap(parsed, "build statement query to find ticket from database")
		}
	}
	return &ticket, nil
}
