package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nipa-interview/service/ping"
	"nipa-interview/service/ticket"
)

type route struct {
	Name        string
	Description string
	Method      string
	Pattern     string
	Endpoint    gin.HandlerFunc
	//Validation  gin.HandlerFunc
}

type Routes struct {
	transaction []route
}

func (r Routes) InitTransactionRoute() http.Handler {


	pingSrv := ping.NewEndpoint()
	ticketSrv := ticket.NewEndpoint()


	txPing := []route{
		{
			Name:        "pingGet",
			Description: "pingGet",
			Method:      http.MethodGet,
			Pattern:     "/ping",
			Endpoint:    pingSrv.PingGetEndpoint,
		},
		{
			Name:        "PingPost",
			Description: "PingPost",
			Method:      http.MethodPost,
			Pattern:     "/pong",
			Endpoint:    pingSrv.PingPostEndpoint,
		},
	}
	txTicket := []route{
		{
			Name:        "addTicket",
			Description: "addTicket",
			Method:      http.MethodPost,
			Pattern:     "/addTicket",
			Endpoint:    ticketSrv.AddTicket,
		},
		{
			Name:        "updateTicket",
			Description: "updateTicket",
			Method:      http.MethodPost,
			Pattern:     "/updateTicket",
			Endpoint:    ticketSrv.UpdateTicket,
		},
		{
			Name:        "listTicket",
			Description: "listTicket",
			Method:      http.MethodPost,
			Pattern:     "/listTicket",
			Endpoint:    ticketSrv.TicketList,
		},
	}

	ro := gin.New()

	store := ro.Group("/test")
	for _, e := range txPing {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	store = ro.Group("/app")
	for _, e := range txTicket {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	return ro
}
