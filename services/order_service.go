package services

import (
	"Abishar-BPJS_Test-Joe_Allen_Butarbutar/config"
	"Abishar-BPJS_Test-Joe_Allen_Butarbutar/models"
	"fmt"
	"log"
	
)

const (
	CreateOrder  	 = `INSERT INTO requests (request_id, created_at) VALUES($1, $2) RETURNING request_id, created_at`
	CreateordData  = `INSERT INTO request_data (id, customer, quantity, price, timestamp, request_id) VALUES($1, $2, $3, $4, $5, $6) RETURNING id, customer, quantity, price, timestamp, request_id`
	GetAllOrder   	 = `SELECT * FROM requests`
	GetRequestDataByRequestId = `SELECT * FROM request_data WHERE request_id = $1`
)

var OrderService orderService = &orderRepo{}

type orderService interface {
	CreateOrder(*models.Order) *models.Order
	GetOrder() *[]models.Order
}

type orderRepo struct{}

func (m *orderRepo) CreateOrder(reqPayload *models.Order) *models.Order {
	db := db.GetDB()
	fmt.Println(reqPayload.RequestId)
	row := db.QueryRow(CreateOrder, reqPayload.RequestId, reqPayload.CreatedAt)
	err := row.Scan(&reqPayload.RequestId, &reqPayload.CreatedAt)
	if err != nil {
		log.Fatal(err)
	}

	request_data := []models.RequestData{}
	for _, requestData := range reqPayload.RequestData {
		row = db.QueryRow(CreateordData, requestData.Id, requestData.Customer, requestData.Quantity, requestData.Price, requestData.Timestamp, reqPayload.RequestId)
		var requestDataResult models.RequestData
		err = row.Scan(&requestDataResult.Id, &requestDataResult.Customer, &requestDataResult.Quantity, &requestDataResult.Price, &requestDataResult.Timestamp, &requestDataResult.RequestId)
		if err != nil {
			log.Fatal(err)
		}
		request_data = append(request_data, requestDataResult)
	}

	reqPayload.RequestData = request_data

	return reqPayload
}

func (m *orderRepo) GetOrder() *[]models.Order {
	db := db.GetDB()
	row, err := db.Query(GetAllOrder)
	if err != nil {
		log.Fatal("fail to get data")
	}

	var dataOrder []models.Order
	for row.Next() {
		var ordData models.Order
		if err := row.Scan(&ordData.RequestId, &ordData.CreatedAt); err != nil {
			log.Fatal("err")
		}

		row2, err := db.Query(GetRequestDataByRequestId, &ordData.RequestId)
		if err != nil {
			log.Fatal("failed to query request data")
		}

		var dataRequestData []models.RequestData
		for row2.Next() {
			var requestData models.RequestData
			if err := row2.Scan(&requestData.Id, &requestData.Customer, &requestData.Quantity, &requestData.Price, &requestData.Timestamp, &requestData.RequestId); err != nil {
				log.Fatal(err)
			}
			dataRequestData = append(dataRequestData, requestData)

		}
		ordData.RequestData = append(ordData.RequestData, dataRequestData...)
		dataOrder = append(dataOrder, ordData)
	}

	if err = row.Err(); err != nil {
		log.Fatal("Error ")
	}

	return &dataOrder
}