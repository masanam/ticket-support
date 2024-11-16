package ticketmodel

import (
	"go-ticket-support/entity/ticketentity"
	"log"
	"time"
)

func (r *TicketRepository) ListTicket(sortBy, orderBy string, perPage, page int, filterName string, filterType string, filterValue string) (*[]ticketentity.Tickets, int64, error) {
	var tickets []ticketentity.Ticket
	var ticketsGet []ticketentity.Tickets
	var count int64

	offset := (page - 1) * perPage
	order := orderBy + " " + sortBy

	// log.Println(&tickets)
	// log.Println(filterType)
	// log.Println(filterValue)

	if (filterName != "") && (filterName == `created_at`) {
		beforeDate, e := time.Parse("2006-01-02 15:04:05", filterValue)
		if e != nil {
			log.Println(beforeDate)
		}
		err := r.db.Model(&tickets).
			Preload("Status").
			Order(order).
			Limit(perPage).
			Offset(offset).
			Where(filterName+" "+filterType+" ?", beforeDate).
			Find(&ticketsGet).Error

		if err != nil {
			panic("failed to load data")
		}
	} else if (filterName != "") && (filterName != `created_at`) {
		err := r.db.Model(&tickets).
			Preload("Status").
			Order(order).
			Limit(perPage).
			Offset(offset).
			Where(filterName+" "+filterType+" ?", filterValue).
			Find(&ticketsGet).Error

		if err != nil {
			panic("failed to load data")
		}
	} else {
		// r.db.Debug().Preload("Status").Find(&ticketsGet)

		err := r.db.Model(&tickets).
			Preload("Status").
			Order(order).
			Limit(perPage).
			Offset(offset).
			Find(&ticketsGet).Error

		if err != nil {
			panic("failed to load data")
		}

	}

	count = int64(len(tickets))
	return &ticketsGet, count, nil

}
