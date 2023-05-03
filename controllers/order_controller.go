package controllers

import (
	"Abishar-BPJS_Test-Joe_Allen_Butarbutar/models"
	service "Abishar-BPJS_Test-Joe_Allen_Butarbutar/services"

	"fmt"
	"time"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	start := time.Now()

	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"errMsg": err.Error(),
		})
		return
	}

	order_data := []models.RequestData{}
	for _, v := range data["data"].([]interface{}) {
		orderData := v.(map[string]interface{})

		var TimestampToString string
		// _ = TimestampToString
		if val, ok := orderData["timestamp"].(string); ok {
			TimestampToString = val
		}
		timestamp, err := time.Parse("2006-01-02 15:04:05", TimestampToString)
		if err != nil {
			fmt.Println(err)
		}

		orderDataReq := []models.RequestData{
			{
				Id		  : int(orderData["id"].(float64)),
				Customer  : orderData["customer"].(string),
				Quantity  : int(orderData["quantity"].(float64)),
				Price	  : float64(orderData["price"].(float64)),
				Timestamp : timestamp,
			},
		}
		order_data = append(order_data, orderDataReq...)
	}
	req_payload := models.Order{
		RequestId   : int(data["request_id"].(float64)),
		CreatedAt   : time.Now(),
		RequestData : order_data,
	}
	// _ = req

	_ = service.OrderService.CreateOrder(&req_payload)
	// c.JSON(201, res)
	timeElapsed := time.Since(start).Seconds()
	c.JSON(http.StatusOK, gin.H{
		"message": "Data has been store",
		"response_time":   timeElapsed,
	})

}

func GetOrder(c *gin.Context) {
	res := service.OrderService.GetOrder()
	// c.JSON(201, res)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":   res,
	})
}