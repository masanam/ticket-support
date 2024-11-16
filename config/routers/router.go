package routers

import (
	repositories "go-ticket-support/model"
	"go-ticket-support/route"
	"go-ticket-support/src/ticket"
)

func (se *Serve) initializeRoutes() {
	//======================== REPOSITORIES ========================
	//initiate repository
	r := repositories.NewRepository(se.DB)

	//======================== ROUTER ========================
	//Setting Services
	//Setting Ticket Service

	s := ticket.NewService(r)
	h := route.NewTicketHandler(s)
	//=========================================================

	//======================== ENDPOINT ========================
	//Initialize endpoint route
	se.Router.HandleFunc("/ticket/create", h.RegisterNewTicket).Methods("POST")
	se.Router.HandleFunc("/ticket/{id}/find", h.FindTicketByTicketID).Methods("GET")
	se.Router.HandleFunc("/ticket", h.GetTickets).Methods("GET")
	se.Router.HandleFunc("/ticket", h.UpdateDataTickets).Methods("PUT")
	se.Router.HandleFunc("/ticket/{id}", h.TicketDelete).Methods("DELETE")
	//==========================================================

}
