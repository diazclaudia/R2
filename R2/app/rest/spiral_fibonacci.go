package rest

import (
	"net/http"
	"strconv"

	"R2/app/domain"
	"github.com/gin-gonic/gin"
)

func SpiralFibonacci(context *gin.Context) {
	rows, err := strconv.Atoi(context.Query("rows"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	cols, err := strconv.Atoi(context.Query("cols"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	response, err := domain.SpiralFibonacci(rows, cols)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"fibonacci_spiral": response,
	})
}
