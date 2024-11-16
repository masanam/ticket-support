package ticketentity

import (
	"time"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Ticket struct {
	ID            string    `json:"id" gorm:"size:36;not null;unique index;primaryKey"`
	TicketTitle   string    `json:"ticket_title" gorm:"size:255;"`
	TicketMessage string    `json:"ticket_msg" gorm:"size:255;"`
	UserID        string    `json:"user_id" gorm:"size:50;"`
	StatusCode    string    `json:"status" gorm:"size:3;"`
	CreatedAt     time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

// Tickets represent body for list
type Tickets struct {
	ID            string       `json:"id"`
	TicketTitle   string       `json:"ticket_title"`
	TicketMessage string       `json:"ticket_msg"`
	UserID        string       `json:"user_id"`
	StatusCode    string       `json:"status_code"`
	CreatedAt     string       `json:"created_at"`
	UpdatedAt     string       `json:"updated_at"`
	Status        TicketStatus `gorm:"foreignKey:StatusCode;references:Code"`
}

type TicketStatus struct {
	ID        string    `json:"id" gorm:"size:36;not null;unique index;primaryKey"`
	Code      string    `json:"code" gorm:"size:3;"`
	Status    string    `json:"status" gorm:"size:255;"`
	CreatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
}

// Tickets represent body for get data from ticket
type TicketData struct {
	ID            string `json:"id"`
	TicketTitle   string `json:"ticket_title"`
	TicketMessage string `json:"ticket_msg"`
	UserID        string `json:"user_id"`
	StatusCode    string `json:"status"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

type TicketRequest struct {
	TicketTitle   string `json:"ticket_title"`
	TicketMessage string `json:"ticket_msg"`
	UserID        string `json:"user_id"`
	StatusCode    string `json:"status"`
}

// TicketRequestValidate is to validate input request
func TicketRequestValidate(ur *TicketRequest) error {
	err := validation.Errors{
		"ticket_title": validation.Validate(&ur.TicketTitle, validation.Required, validation.Length(10, 100).Error("ticket_title should be minimum 10 characters and maximum 100 characters")),
		"ticket_msg":   validation.Validate(&ur.TicketMessage, validation.Required, validation.Length(100, 0).Error("ticket_msg should be minimum 100 characters")),
		// "user_id":      validation.Validate(&ur.UserID, validation.Required, validation.Match(regexp.MustCompile(`^[0-9]+$`)).Error("user_id should be an integer / numeric")),
		"user_id": validation.Validate(&ur.UserID, validation.Required, is.Int.Error("user_id should be an integer / numeric")),
	}

	return err.Filter()
}

// FindTicket is struct to handle respone while find ticket by ID
type FindTicket struct {
	TicketTitle   string `json:"ticket_title"`
	TicketMessage string `json:"ticket_msg"`
	UserID        string `json:"user_id"`
	StatusCode    string `json:"status"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

func SetParameterUpdateTicket(ticket *Ticket, payload *TicketData) *Ticket {
	var ticketdata Ticket

	if payload.TicketTitle == "" {
		ticketdata.TicketTitle = ticket.TicketTitle
	} else {
		ticketdata.TicketTitle = payload.TicketTitle
	}

	if payload.TicketMessage == "" {
		ticketdata.TicketMessage = ticket.TicketMessage
	} else {
		ticketdata.TicketMessage = payload.TicketMessage
	}

	if payload.StatusCode == "" {
		ticketdata.StatusCode = ticket.StatusCode
	} else {
		ticketdata.StatusCode = payload.StatusCode
	}

	if payload.UserID == "" {
		ticketdata.UserID = ticket.UserID
	} else {
		ticketdata.UserID = payload.UserID
	}

	ticketdata.UpdatedAt = time.Now()

	return &ticketdata
}
