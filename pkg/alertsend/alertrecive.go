package alertsend

import (
	"github.com/gin-gonic/gin"
)

func AlertReceive(c *gin.Context) {
	//// webhook 来源包：github.com/prometheus/alertmanager/notify/webhhok
	//var msg = webhook.Message
	//err := c.BindJSON(&msg)
	//if err != nil {
	//	return
	//}
	//c.String(http.StatusOK, "ok")
}
