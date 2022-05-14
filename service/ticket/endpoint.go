package ticket

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type Endpoint struct {
}

func NewEndpoint() *Endpoint {
	return &Endpoint{}
}

func (ep *Endpoint) AddTicket(c *gin.Context) { //POST /app/addTicket
	defer c.Request.Body.Close()
	log.Info("Ticket : AddTicket")

	var request inputTicket //model รับ input จาก body
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}
	log.Debugf("Title : [%+v]", request.Title)
	log.Debugf("Description : [%+v]", request.Description)
	log.Debugf("Contact Information : [%+v]", request.ContactInformation)

	msg, err := addTicket(request)
	if err != nil {
		//return err
		c.JSON(msg.Status, msg)
		return
	}
	c.JSON(msg.Status, msg)
	return

}

func (ep *Endpoint) UpdateTicket(c *gin.Context) { //POST /app/addTicket
	defer c.Request.Body.Close()
	log.Info("Ticket : UpdateTicket")

	var request inputTicketUpdate //model รับ input จาก body
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}

	log.Debugf("Id : [%+v]", request.Id)
	log.Debugf("Title : [%+v]", request.Title)
	log.Debugf("Description : [%+v]", request.Description)
	log.Debugf("Contact Information : [%+v]", request.ContactInformation)
	log.Debugf("StatusTicket : [%+v]", request.Status)
	msg, err := updateStatusTicket(request)
	if err != nil {
		//return err
		c.JSON(msg.Status, msg)
		return
	}
	c.JSON(msg.Status, msg)
	return
}

func (ep *Endpoint) TicketList(c *gin.Context) { //POST /app/ticketList
	defer c.Request.Body.Close()
	log.Info("Ticket : TicketList")
	var request inputTicketList
	if err := c.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		//return err
		c.JSON(http.StatusBadRequest, err)
		return
	}

	log.Debugf("Status : [%+v]", request.Status)
	log.Debugf("FilterTitle : [%+v]", request.FilterTitle)
	log.Debugf("FilterCreateDateFrom : [%+v]", request.FilterCreateDateFrom)
	log.Debugf("FilterCreateDateTo : [%+v]", request.FilterCreateDateTo)
	log.Debugf("FilterUpdateDateFrom : [%+v]", request.FilterUpdateDateFrom)
	log.Debugf("FilterUpdateDateTo : [%+v]", request.FilterUpdateDateTo)
	log.Debugf("SortBy : [%+v]", request.SortBy)
	log.Debugf("SortType : [%+v]", request.SortType)
	log.Debugf("PagingIndex : [%+v]", request.PagingIndex)
	log.Debugf("PagingSize : [%+v]", request.PagingSize)

	result, msg, err := getTicketList(request)
	if err != nil {
		//return err
		c.JSON(msg.Status, msg)
		return
	}
	c.JSON(msg.Status, result)
	return

}