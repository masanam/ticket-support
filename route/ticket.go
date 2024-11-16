package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"go-ticket-support/entity"
	"go-ticket-support/entity/ticketentity"
	"go-ticket-support/helpers"
	"go-ticket-support/src/ticket"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

type TicketHandler struct {
	service ticket.Service
}

func NewTicketHandler(service ticket.Service) *TicketHandler {
	return &TicketHandler{
		service: service,
	}
}

// RegisterNewTicket is func to Handle ticket registration
func (uh *TicketHandler) RegisterNewTicket(w http.ResponseWriter, r *http.Request) {
	responder := helpers.NewHTTPResponse("registerNewTicket")
	ctx := r.Context()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
		return
	}

	var payload ticketentity.TicketRequest
	err = json.Unmarshal(body, &payload)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
		return
	}

	newTicket, err := uh.service.InsertNewTicket(ctx, &payload) //uh.service.InsertNewTicket(&payload)
	if err != nil {
		caticket := errors.Cause(err)
		switch caticket {
		case entity.ErrTicketAlreadyExist:
			responder.FieldErrors(w, err, http.StatusNotAcceptable, err.Error())
			return
		default:
			responder.FieldErrors(w, err, http.StatusInternalServerError, fmt.Sprint(err))
			return
		}
	}
	responder.SuccessJSON(w, newTicket, http.StatusCreated, "Succes to register new ticket")
	return
}

func (ud *TicketHandler) FindTicketByTicketID(w http.ResponseWriter, r *http.Request) {
	var (
		ticketID  = mux.Vars(r)["id"]
		responder = helpers.NewHTTPResponse("registerNewTicket")
		ctx       = r.Context()
	)

	findTicket, err := ud.service.FindTicket(ctx, ticketID)
	if err != nil {
		caticket := errors.Cause(err)
		switch caticket {
		case entity.ErrTicketNotExist:
			responder.ErrorJSON(w, http.StatusNotFound, "ticket not found")
			return
		default:
			responder.FailureJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	responder.SuccessJSON(w, findTicket, http.StatusOK, "Ticket found")
	return
}

func (uh *TicketHandler) GetTickets(w http.ResponseWriter, r *http.Request) {
	var (
		param            = r.URL.Query()
		paramPage        = param.Get("page")
		paramPerPage     = param.Get("per_page")
		paramOrderBy     = param.Get("order_by")
		paramSortBy      = param.Get("sort_by")
		paramFilterName  = param.Get("filter_name")
		paramFilterType  = param.Get("filter_type")
		paramFilterValue = param.Get("filter_value")

		responder = helpers.NewHTTPResponse("registerNewTicket")
		ctx       = r.Context()
	)

	filterParam := helpers.SetDefaultFilterParam(paramFilterName, paramFilterType, paramFilterValue)
	filterName := filterParam.FilterName
	filterType := filterParam.FilterType
	filterValue := filterParam.FilterValue

	filter, err := helpers.GetFilter(helpers.FilterParams{
		FilterName:  filterName,
		FilterType:  filterType,
		FilterValue: filterValue,
	})

	paginationParam, err := helpers.SetDefaultPginationParam(paramPage, paramPerPage, paramOrderBy, paramSortBy)
	if err != nil {
		responder.FieldErrors(w, err, http.StatusUnprocessableEntity, "value of query parameters has diferent type")
		return
	}
	sortBy := paginationParam.SortBy
	orderBy := paginationParam.OrderBy
	perPage := paginationParam.PerPage
	pagesize := int32(paginationParam.PerPage)
	page, _ := strconv.Atoi(paginationParam.Page)

	tickets, total, err := uh.service.GetListTickets(ctx, sortBy, orderBy, int(perPage), page, filterName, filterType, filterValue)
	perPage = perPage - (perPage % 10)
	log.Println(perPage)

	if err != nil {
		caticket := errors.Cause(err)
		switch caticket {
		case entity.ErrTicketNotExist:
			responder.ErrorJSON(w, http.StatusNotFound, "tickets list not found")
			return
		default:
			responder.FailureJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	pagination, err := helpers.GetPagination(helpers.PaginationParams{
		Path:        "list.tickets",
		Page:        strconv.Itoa(page),
		TotalRows:   int32(total),
		PerPage:     int32(perPage),
		OrderBy:     orderBy,
		SortBy:      sortBy,
		CurrentPage: int32(page),
	})

	if err != nil {
		responder.ErrorJSON(w, http.StatusConflict, "error pagination")
		return
	}
	pagesize = pagesize - (pagesize % 10)

	responder.SuccessWithMeta(w, tickets, filter, pagination, pagesize, http.StatusOK, "tickets list")
	return

}

func (uh *TicketHandler) UpdateDataTickets(w http.ResponseWriter, r *http.Request) {
	var (
		responder = helpers.NewHTTPResponse("updateDataTicket")
		ctx       = r.Context()
	)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
		return
	}

	var payload ticketentity.TicketData
	err = json.Unmarshal(body, &payload)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
		return
	}

	updatedTicket, err := uh.service.UpdateTicket(ctx, &payload)
	if err != nil {
		caticket := errors.Cause(err)
		switch caticket {
		case entity.ErrTicketNotExist:
			responder.FieldErrors(w, err, http.StatusNotExtended, err.Error())
			return
		default:
			responder.FieldErrors(w, err, http.StatusInternalServerError, fmt.Sprint(err))
			return
		}
	}
	responder.SuccessJSON(w, updatedTicket, http.StatusCreated, "Succes to update data ticket")
	return
}

func (uh *TicketHandler) TicketDelete(w http.ResponseWriter, r *http.Request) {
	var (
		responder = helpers.NewHTTPResponse("updateDataTicket")
		ctx       = r.Context()
		ticketID  = mux.Vars(r)["id"]
	)

	err := uh.service.DeleteDataTicket(ctx, ticketID)
	if err != nil {
		caticket := errors.Cause(err)
		switch caticket {
		case entity.ErrTicketNotExist:
			responder.ErrorJSON(w, http.StatusNotFound, "ticket not found")
			return
		default:
			responder.FailureJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	responder.SuccessWithoutData(w, http.StatusOK, "successfully to delete ticket")
	return
}
