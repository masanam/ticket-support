package app

import (
	"go-ticket-support/entity/ticketentity"
)

// SetMigrationTable is used to register entity model which want to be migrate
func SetMigrationTable() []interface{} {
	var migrationData = []interface{}{
		&ticketentity.Ticket{},
		&ticketentity.TicketStatus{},
	}

	return migrationData
}
