package entity

type Error string

// Declare error messege
const (
	ErrPermissionNotAllowed = Error("permission.not_allowed")

	//Ticket Error
	ErrTicketNotExist            = Error("domain.ticket.error.not_exist")
	ErrTicketAlreadyExist        = Error("domain.ticket.error.ticket_already_exist")
	ErrTicketsCredentialNotExist = Error("domain.ticket.error.credential_not_exist")
)

func (e Error) Error() string {
	return string(e)
}
