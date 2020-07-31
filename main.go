package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"score_query_http_server/modules"
	"sync"
)

var lock sync.RWMutex

func init() {
	// 日志格式设置
	log.SetFlags(log.Llongfile | log.LstdFlags)
	// 初始化数据库
	modules.Init()
	// 设置GIN为发布模式
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	// 启动HTTP服务
	log.Println("HTTP服务启动")
	runServer()
}

func runServer() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/send-student-id", func(c *gin.Context) {
		studentIdStr := c.Query("id")
		if studentIdStr == "" {
			c.JSON(200, gin.H{
				"status": -1,
			})
			return
		}
		// 输出到终端
		log.Println("添加学号:",studentIdStr)
		// 纳入数据库
		studentId := new(modules.Id)
		if err := studentId.GetAutoCreate(map[string]interface{}{"group": studentIdStr}); err != nil {
			c.JSON(200, gin.H{
				"status": -1,
			})
			return
		}
		// 返回状态
		c.JSON(200, gin.H{
			"status": 1,
		})
	})
	if err := r.Run(":9090"); err != nil {
		log.Panic(err)
	}
}
