package ticket

import (
	"context"
	"time"

	"go-ticket-support/entity/ticketentity"
	repositories "go-ticket-support/model"

	"github.com/google/uuid"
)

type service struct {
	repo *repositories.Repository
}

func NewService(repo *repositories.Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) InsertNewTicket(ctx context.Context, payload *ticketentity.TicketRequest) (*ticketentity.Ticket, error) {

	err := ticketentity.TicketRequestValidate(payload)
	if err != nil {
		return nil, err
	}
	// hashPassword, _ := helpers.HashPassword(payload.Password)

	ticket := &ticketentity.Ticket{
		ID:            uuid.NewString(),
		TicketTitle:   payload.TicketTitle,
		TicketMessage: payload.TicketMessage,
		StatusCode:    payload.StatusCode,
		UserID:        payload.UserID,
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
	}

	err = s.repo.Ticket.SaveNewTicket(ctx, ticket)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (s *service) FindTicket(ctx context.Context, ticketID string) (*ticketentity.FindTicket, error) {

	ticket, err := s.repo.Ticket.FindTicketByID(ctx, ticketID)
	if err != nil {
		return nil, err
	}

	resultTicket := ticketentity.FindTicket{
		TicketTitle:   ticket.TicketTitle,
		TicketMessage: ticket.TicketMessage,
		StatusCode:    ticket.StatusCode,
		UserID:        ticket.UserID,
		CreatedAt:     ticket.CreatedAt.String(),
		UpdatedAt:     ticket.UpdatedAt.String(),
	}

	return &resultTicket, nil
}

func (s *service) GetListTickets(ctx context.Context, sortBy, orderBy string, perPage, page int, filterName, filterType, filterValue string) (*[]ticketentity.Tickets, int64, error) {
	tickets, total, err := s.repo.Ticket.ListTicket(sortBy, orderBy, perPage, page, filterName, filterType, filterValue)
	if err != nil {
		return nil, 0, err
	}

	return tickets, total, nil
}

func (s *service) UpdateTicket(ctx context.Context, payload *ticketentity.TicketData) (*ticketentity.TicketData, error) {
	ticket, err := s.repo.Ticket.FindTicketByID(ctx, payload.ID)
	if err != nil {
		return nil, err
	}

	ticketdata := ticketentity.SetParameterUpdateTicket(ticket, payload)

	updateTicket, err := s.repo.Ticket.UpdateTicketData(ctx, payload.ID, ticketdata)
	if err != nil {
		return nil, err
	}

	return updateTicket, nil
}

func (s *service) DeleteDataTicket(ctx context.Context, ticketID string) error {

	err := s.repo.Ticket.DeleteTicket(ctx, ticketID)
	if err != nil {
		return err
	}

	return nil
}
