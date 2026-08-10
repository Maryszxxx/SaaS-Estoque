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

func UpdatePatch(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "PATCH"})
}

func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "DELETE"})
}

func Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "GET"})
}
func GetById(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "GET"})
}

// implementação de login

func PostUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"method": "POST"})
}
