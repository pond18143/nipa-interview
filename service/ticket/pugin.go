package ticket

import (
	"github.com/jinzhu/gorm"
	"nipa-interview/database/mssql"
	"time"
)

func addNewTicket(tx *gorm.DB, request ticket) (err error) {
	if tx == nil {
		tx = mssql.DB
	}
	if err = tx.Table("ticket").
		Create(&request).Error; err != nil {
		return
	}
	return
}

func updateStatus(tx *gorm.DB, id int64, detail ticket) (err error) {
	if tx == nil {
		tx = mssql.DB
	}

	if err = tx.Table("ticket").
		Where("id = ?", id).
		Update(map[string]interface{}{
			"title":               detail.Title,
			"description":         detail.Description,
			"contact_information": detail.ContactInformation,
			"status":              detail.Status,
			"update_timestamp":    detail.UpdateTimestamp,
		}).
		Error; err != nil {
		return
	}
	return
}

func getTicketById(id int64) (detail ticket, err error) {
	if err = mssql.DB.Select("*").
		Table("ticket").
		Where("id = ?", id).
		Find(&detail).Error; err != nil {
		return
	}
	return
}

func ticketListGet(request inputTicketList) (result ticketListRes, err error) {
	var zeroValueTime time.Time // 0001-01-01 00:00:00 +0000
	var zeroValueString string  // ""
	var zeroValueInt int64      // 0

	var SortingBy string
	switch request.SortBy {
	case "title":
		SortingBy = "t.title"
	case "id":
		SortingBy = "t.id"
	case "created_timestamp":
		SortingBy = "t.created_timestamp"
	case "update_timestamp":
		SortingBy = "t.update_timestamp"
	default:
		SortingBy = "t.status"
	}

	if request.SortType != 0 {
		SortingBy += " desc"
	}

	txSubQuery := mssql.DB
	txMainQuery := mssql.DB

	txSubQuery = txSubQuery.Select("ROW_NUMBER() OVER(ORDER BY (SELECT NULL) ) AS sequence, t.id, t.title, s.name AS status_name, t.description, t.contact_information, t.created_timestamp, t.update_timestamp").
		Table("ticket AS t").
		Joins("LEFT JOIN status AS s ON s.id = t.status")

	//!status
	if int64(request.Status) != zeroValueInt {
		txSubQuery = txSubQuery.Where("t.status = ?", request.Status)
	}
	//!FilterTitle
	if request.FilterTitle != zeroValueString {
		txSubQuery = txSubQuery.Where("t.title = ? ", request.FilterTitle)
	}
	//!FilterCreateDateFrom
	if request.FilterCreateDateFrom != zeroValueTime {
		txSubQuery = txSubQuery.Where("t.created_timestamp >= ? ", request.FilterCreateDateFrom)
	}
	//!FilterCreateDateTo
	if request.FilterCreateDateTo != zeroValueTime {
		txSubQuery = txSubQuery.Where("t.created_timestamp <= ? ", request.FilterCreateDateTo)
	}
	//!FilterUpdateDateFrom
	if request.FilterUpdateDateFrom != zeroValueTime {
		txSubQuery = txSubQuery.Where("t.update_timestamp >= ? ", request.FilterUpdateDateFrom)
	}
	//!FilterUpdateDateTo
	if request.FilterUpdateDateTo != zeroValueTime {
		txSubQuery = txSubQuery.Where("t.update_timestamp <= ? ", request.FilterUpdateDateTo)
	}

	if err = txMainQuery.Raw("SELECT COUNT (sub.title ) AS ticket_count FROM ? AS sub", txSubQuery.SubQuery()).Find(&result.Header).Error; err != nil {
		return
	}

	txSubQuery = txSubQuery.Order(SortingBy)
	//!zeroPagingIndex
	if request.PagingIndex != zeroValueInt {
		txSubQuery = txSubQuery.Offset((request.PagingIndex - 1) * request.PagingSize)
	}

	//!zeroPagingSize
	if request.PagingSize != zeroValueInt {
		txSubQuery = txSubQuery.Limit(request.PagingSize)
	}

	//defaultPagingIndex Page1
	if request.PagingIndex == 0 {
		txSubQuery = txSubQuery.Offset(0)
	}

	//defaultPagingSize Size10
	if request.PagingSize == 0 {
		txSubQuery = txSubQuery.Limit(100)
	}

	if err = txSubQuery.Find(&result.Detail).Error; err != nil {
		return
	}
	return

	return
}
