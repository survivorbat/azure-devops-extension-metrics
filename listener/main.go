package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	log.Fatal(r.Run())
}
