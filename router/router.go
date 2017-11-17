package router 

import "github.com/gin-gonic/gin"
import "../modules/v1"
import "../modules/v2"

func Router(router *gin.Engine) {
    router.GET("/v1/search", v1.GetOne)
    router.GET("/v2/hehes", v2.GetTwo)
}
