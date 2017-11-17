package main

import "github.com/gin-gonic/gin"
import "./router"
import "./modules/utils"
import "strconv"
import "./config"

func main() {
    config.ConfigInit()
    utils.SegmenterInit()
    utils.DBInit()
    r := gin.Default()
    r.Use(gin.Logger())
    router.Router(r) 
    r.Run(":"+strconv.Itoa(config.Port))
}
