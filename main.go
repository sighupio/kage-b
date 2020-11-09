package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/n", func(c *gin.Context) {
		n := 42
		c.String(http.StatusOK, fmt.Sprint(n))
	})

	router.Run(":8081")
}
