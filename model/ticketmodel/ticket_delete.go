package ticketmodel

import (
	"context"

	"go-ticket-support/config/app"
	"go-ticket-support/entity/ticketentity"

	"github.com/pkg/errors"
)

func (r *TicketRepository) DeleteTicket(ctx context.Context, ticketID string) error {
	var ticket ticketentity.Ticket

	err := r.db.Model(&ticket).Where("id = ?", ticketID).Take(&ticket).Delete(ticket).Error
	if err != nil {
		parsed := r.codeError.ParseSQLError(err)
		switch parsed {
		case app.ErrNoRowsFound:
			return app.ErrNoRowsFound
		case app.ErrUniqueViolation:
			return app.ErrUniqueViolation
		default:
			return errors.Wrap(parsed, "build statement query to delete ticket data from database")
		}
	}

	return nil
}
