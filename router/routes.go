package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Post(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "POST"})
}

func Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "PUT"})
}

func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "DELETE"})
}

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "GET"})
}
