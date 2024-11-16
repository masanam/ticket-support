package ticketmodel

import (
	"context"

	"go-ticket-support/config/app"
	"go-ticket-support/entity/ticketentity"

	"github.com/pkg/errors"
)

func (r *TicketRepository) UpdateTicketData(ctx context.Context, ticketID string, payload *ticketentity.Ticket) (*ticketentity.TicketData, error) {
	var tickets ticketentity.Ticket

	err := r.db.Debug().Model(&tickets).Where("id = ?", ticketID).Take(&tickets).UpdateColumns(payload).Error

	if err != nil {
		parsed := r.codeError.ParseSQLError(err)
		switch parsed {
		case app.ErrNoRowsFound:
			return nil, app.ErrNoRowsFound
		case app.ErrUniqueViolation:
			return nil, app.ErrUniqueViolation
		default:
			return nil, errors.Wrap(parsed, "build statement query to update ticket data from database")
		}
	}

	ticketData := &ticketentity.TicketData{
		ID:            ticketID,
		TicketTitle:   tickets.TicketTitle,
		TicketMessage: tickets.TicketMessage,
		StatusCode:    tickets.StatusCode,
		UserID:        tickets.UserID,
		CreatedAt:     tickets.CreatedAt.String(),
		UpdatedAt:     tickets.UpdatedAt.String(),
	}

	return ticketData, nil
}
