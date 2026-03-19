package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	InitDB()

	r := gin.Default()

	r.POST("/parse", ParseHandler)
	r.POST("/bill", BillHandler)

	r.Run(":8080")
}