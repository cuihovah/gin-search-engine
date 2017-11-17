package middleware

import "github.com/gin-gonic/gin"
import "time"
import "fmt"

func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()
        // 设置example变量到Context的Key中,通过Get等函数可以取得
        c.Set("example", "12345")
        // 发送request之前
        c.Next()
        // 发送request之后
        latency := time.Since(t)
        fmt.Println(latency)

        // 这个c.Write是ResponseWriter,我们可以获得状态等信息
        status := c.Writer.Status()
	fmt.Println(c.Request.Header.Get("User-Agent"))
	
        fmt.Println(status)
    }
}
