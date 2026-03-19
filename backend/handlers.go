package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ParseHandler(c *gin.Context) {
	var req ParseRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ParseWithAI(req.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "AI failed"})
		return
	}

	c.JSON(http.StatusOK, result)
}

func BillHandler(c *gin.Context) {
	var data ParseResponse

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	total := len(data.Drugs)*500 + len(data.Tests)*1000

	c.JSON(http.StatusOK, BillResponse{
		Total: total,
	})
}