package ticket

import (
	"time"
)

type StatusTicket int32

const (
	//status Ticket
	Pending  StatusTicket = 1
	Accepted StatusTicket = 2
	Resolved StatusTicket = 3
	Rejected StatusTicket = 4
)

type messageResponse struct {
	Status             int    `json:"status" example:"201"`
	MessageDescription string `json:"message_description"`
}

type ticket struct {
	Title              string       `json:"title"`
	Description        string       `json:"description"`
	ContactInformation string       `json:"contact_information"`
	Status             StatusTicket `json:"status"`
	CreatedTimestamp   time.Time    `json:"created_timestamp"`
	UpdateTimestamp    time.Time    `json:"update_timestamp"`
}

type inputTicket struct {
	Title              string `json:"title" example:"Harry potter" `
	Description        string `json:"description" example:"story of school magic"`
	ContactInformation string `json:"contact_information" example:"address......"`
}

type inputTicketUpdate struct {
	Id                 int64        `json:"id" example:"1"`
	Title              string       `json:"title" example:"Harry potter2"`
	Description        string       `json:"description" example:"bra2 bra2 bra2"`
	ContactInformation string       `json:"contact_information" example:"address2......"`
	Status             StatusTicket `json:"status" example:"2"`
}

type inputTicketList struct {
	Status               StatusTicket `json:"status" example:"1"`
	FilterTitle          string       `json:"filter_title" example:"Harryporter4"`
	FilterCreateDateFrom time.Time    `json:"filter_create_date_from" example:"2018-01-11T01:02:18.070Z" format:"date-time"`
	FilterCreateDateTo   time.Time    `json:"filter_create_date_to" example:"2023-01-11T01:02:18.070Z" format:"date-time"`
	FilterUpdateDateFrom time.Time    `json:"filter_update_date_from" example:"2018-01-11T01:02:18.070Z" format:"date-time"`
	FilterUpdateDateTo   time.Time    `json:"filter_update_date_to" example:"2023-01-11T01:02:18.070Z" format:"date-time"`
	SortBy               string       `json:"sort_by" example:"id"`
	SortType             int64        `json:"sort_type" example:"1"`
	PagingIndex          int64        `json:"paging_index" example:"1"`
	PagingSize           int64        `json:"paging_size" example:"10"`
}

type ticketCountRes struct {
	TicketCount int64 `json:"ticket_count"`
}

type ticketListRes struct {
	Header ticketCountRes `json:"header"`
	Detail []ticketRes    `json:"detail"`
}

type ticketRes struct {
	Sequence           int64     `json:"sequence"`
	Id                 int64     `json:"id"`
	Title              string    `json:"title"`
	Description        string    `json:"description"`
	ContactInformation string    `json:"contact_information"`
	StatusName         string    `json:"status_name"`
	CreatedTimestamp   time.Time `json:"created_timestamp"`
	UpdateTimestamp    time.Time `json:"update_timestamp"`
}

