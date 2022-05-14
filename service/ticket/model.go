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
	Title              string `json:"title" binding:"required" example:"Harry potter" `
	Description        string `json:"description" example:"story of school magic"`
	ContactInformation string `json:"contact_information" example:"address......"`
}

type inputTicketUpdate struct {
	Id                 int64        `json:"id" binding:"required" example:"1" `
	Title              string       `json:"title" example:"Harry potter2"`
	Description        string       `json:"description" example:"bra2 bra2 bra2"`
	ContactInformation string       `json:"contact_information" example:"address2......"`
	Status             StatusTicket `json:"status" example:"2"`
}

type inputTicketList struct {
	Status               StatusTicket `json:"status"`
	FilterTitle          string       `json:"filter_title"`
	FilterCreateDateFrom time.Time    `json:"filter_create_date_from"`
	FilterCreateDateTo   time.Time    `json:"filter_create_date_to"`
	FilterUpdateDateFrom time.Time    `json:"filter_update_date_from"`
	FilterUpdateDateTo   time.Time    `json:"filter_update_date_to"`
	SortBy               string       `json:"sort_by"`
	SortType             int64        `json:"sort_type"`
	PagingIndex          int64        `json:"paging_index"`
	PagingSize           int64        `json:"paging_size"`
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