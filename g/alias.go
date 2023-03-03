package g

import "gitlab.myshuju.top/heshiying/gin-skeleton/g/log"

var (
	logger log.Log = log.NewLog()
	Debug          = logger.Debug
	Info           = logger.Info
	Warn           = logger.Warn
	Error          = logger.Error
	Panic          = logger.Panic
)

func SetLogger(l log.Log) {
	logger = l
}
