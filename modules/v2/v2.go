package v2

import "github.com/gin-gonic/gin"
import "net/http"

func GetTwo(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg": "OK",
		"data": "v2",
	})
}
