package ticket

import (
	"errors"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"net/http"
	mssql "nipa-interview/database/mssql"
	"time"
)

func addTicket(request inputTicket) (msg messageResponse, err error) {
	tx := mssql.DB.Begin()

	// request title
	if request.Title == "" {
		err = errors.New("missing argument title")
		log.Errorf("addNewTicket failed : [%+v] ", err.Error())
		msg = messageResponse{
			Status:             http.StatusBadRequest,
			MessageDescription: "missing argument title"}
		return
	}

	timeStamp := time.Now()
	//model inset ticket
	var ticket = ticket{
		Title:              request.Title,
		Description:        request.Description,
		ContactInformation: request.ContactInformation,
		Status:             1,
		CreatedTimestamp:   timeStamp,
		UpdateTimestamp:    timeStamp,
	}

	log.Info("insert ticket...")
	// insert ticket
	err = addNewTicket(tx, ticket)
	if err != nil {
		log.Errorf("addNewTicket failed : [%+v] ", err.Error())
		err = errors.New("add new ticket failed")
		log.Errorf("addNewTicket failed : [%+v] ", err.Error())
		msg = messageResponse{
			Status:             http.StatusBadRequest,
			MessageDescription: "add new ticket failed"}
		return
	}
	log.Info("insert success")

	msg = messageResponse{
		Status:             http.StatusCreated,
		MessageDescription: "add ticket success ",
	}
	tx.Commit()
	return
}

func updateStatusTicket(request inputTicketUpdate) (msg messageResponse, err error) {

	// check id
	if request.Id == 0 {
		err = errors.New("missing id")
		log.Errorf("updateStatusTicket failed : [%+v] ", err.Error())
		msg = messageResponse{
			Status:             http.StatusBadRequest,
			MessageDescription: "missing id"}
		return
	}

	// check status
	if request.Status != 0 && request.Status != Pending && request.Status != Accepted && request.Status != Resolved && request.Status != Rejected {
		err = errors.New("status not correct")
		log.Errorf("addNewTicket failed : [%+v] ", err.Error())
		msg = messageResponse{
			Status:             http.StatusBadRequest,
			MessageDescription: "status not correct"}
		return
	}

	// select ticket by id
	var detail ticket
	detail, err = getTicketById(request.Id)
	if err != nil || err == gorm.ErrRecordNotFound {
		log.Errorf("getTicketById failed : [%+v] ", err.Error())
		err = errors.New("getTicketById failed")
		log.Errorf("getTicketById failed : [%+v] ", err.Error())
		msg = messageResponse{
			Status:             http.StatusBadRequest,
			MessageDescription: "ticket not found"}
		return
	}
	log.Debugf("Detail : [%+v]", detail)

	// title == ""
	if request.Title == "" {
		request.Title = detail.Title
	}

	// description == ""
	if request.Description == "" {
		request.Description = detail.Description
	}

	// contact == ""
	if request.ContactInformation == "" {
		request.ContactInformation = detail.ContactInformation
	}

	//if status == 0
	if request.Status == 0 {
		request.Status = detail.Status
	}

	timeStamp := time.Now()
	var ticketUpdate = ticket{
		Title:              request.Title,
		Description:        request.Description,
		ContactInformation: request.ContactInformation,
		Status:             request.Status,
		UpdateTimestamp:    timeStamp,
	}

	log.Info("update status ticket...")
	// update ticket
	tx := mssql.DB.Begin()
	err = updateStatus(tx, request.Id, ticketUpdate)
	if err != nil {
		log.Errorf("updateStatus failed : [%+v] ", err.Error())
		err = errors.New("add new ticket failed")
		log.Errorf("updateStatus failed : [%+v] ", err.Error())
		msg = messageResponse{
			Status:             http.StatusBadRequest,
			MessageDescription: "update status ticket failed"}
		return
	}
	log.Info("update status success")

	msg = messageResponse{
		Status:             http.StatusCreated,
		MessageDescription: "update status ticket success ",
	}
	tx.Commit()
	return
}

func getTicketList(request inputTicketList) (result ticketListRes, msg messageResponse, err error) {

	log.Info("ticketListGet...")
	result, err = ticketListGet(request)
	if err != nil {
		log.Errorf("getTicketList failed : [%+v] ", err.Error())
		err = errors.New("ticket not found")
		log.Errorf("getTicketList failed : [%+v] ", err.Error())
		msg = messageResponse{
			Status:             http.StatusNotFound,
			MessageDescription: "ticket not found"}
		return
	}
	log.Info("list ticket success")
	return result, msg, err
}