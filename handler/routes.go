package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nipa-interview/service/ping"
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

	ro := gin.New()

	store := ro.Group("/app")
	for _, e := range txPing {
		store.Handle(e.Method, e.Pattern, e.Endpoint)
	}

	return ro
}
