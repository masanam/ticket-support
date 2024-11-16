package ticketmodel

import (
	"context"

	"github.com/pkg/errors"

	"go-ticket-support/config/app"
	"go-ticket-support/entity"
	"go-ticket-support/entity/ticketentity"
)

// SaveNewTicket is used to run query insert
func (r *TicketRepository) SaveNewTicket(ctx context.Context, payload *ticketentity.Ticket) error {

	if payload.StatusCode == "" {
		payload.StatusCode = "opn"
	}

	err := r.db.Create(payload).Error
	if err != nil {
		parsed := r.codeError.ParseSQLError(err)
		switch parsed {
		case app.ErrNoRowsFound:
			return entity.ErrTicketNotExist
		case app.ErrUniqueViolation:
			return entity.ErrTicketAlreadyExist
		default:
			return errors.Wrap(parsed, "build statement query to insert ticket from database")
		}
	}
	return nil
}
