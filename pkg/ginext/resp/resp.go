package extend

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Resp struct {
	Success bool        `json:"success" binding:"required" example:"true"` // 请求结果，失败:false，成功:true
	Msg     string      `json:"msg" binding:"required" example:"ok"`       // 请求结果的message
	Data    interface{} `json:"data,omitempty" binding:"required"`         // 返回的数据
	Code    int32       `json:"code"`                                      // code
}

func SendData(ctx *gin.Context, res Resp) {
	ctx.JSON(http.StatusOK, res)
}
func SendParamError(ctx *gin.Context, err error) {
	errInfo := err.Error()
	if errs, ok := err.(validator.ValidationErrors); ok {
		errInfo = getErrorInfo(errs)
	}
	ctx.JSON(http.StatusBadRequest, Resp{
		Success: false,
		Msg:     errInfo,
		Data:    nil,
		Code:    0,
	})
}
func SendError(ctx *gin.Context, err error) {
	errInfo := err.Error()
	ctx.JSON(http.StatusInternalServerError, Resp{
		Success: false,
		Msg:     errInfo,
		Data:    nil,
		Code:    0,
	})
}
