package lib

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
}

func ReturnToJson(w *gin.Context, status int, codeStatus string, desc string, data interface{}) {
	fmt.Println("status ", status)
	var res Response

	res.Status = codeStatus
	res.Message = desc
	res.Data = data

	w.JSON(status, res)
}

func ReturnToJsonWithPaginate(w *gin.Context, status int, codeStatus string, desc string, data interface{}, pagination interface{}) {
	fmt.Println("status ", status)
	var res Response

	res.Status = codeStatus
	res.Message = desc
	res.Data = data
	res.Pagination = pagination

	w.JSON(status, res)
}
