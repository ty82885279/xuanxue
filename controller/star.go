package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"xuanxue/dao/mysql"
	"xuanxue/logic"
	"xuanxue/model"
)

func T(c *gin.Context) {

	param := c.Param("id")
	fmt.Println(param)

	c.JSON(200, gin.H{
		"laidi": "asd ",
	})
}

// FortuneHandler 星座运势查询
// @Summary 星座运势查询
// @Description 星座运势查询
// @Tags 星座相关
// @Accept application/json
// @Produce application/json
// @Param object query model.ParamStar true "查询参数"
// @Success 2001 {object} _ResponseOneFortune "获取星座某时间段运势"
// @Success 2002 {object} _ResponseAllFortune "获取星座所有时间段运势"
// @Router /fortune [get]
func FortuneHandler(c *gin.Context) {
	var p = new(model.ParamStar)

	err := c.ShouldBind(p)
	if err != nil {
		ResponseErr(c, CodeInvalidParam)
		zap.L().Error("param err", zap.Any("err", err))
		return
	}
	fortune, err := logic.GetStarFortune(p)
	if err != nil {
		if err == mysql.ErrNoRows {
			ResponseErr(c, CodeNoData)
			return
		}
		ResponseErr(c, CodeServiceBusy)
		zap.L().Error("logic.GetStarFortune(p) err", zap.Error(err))
		return
	}
	if _, ok := fortune.(*model.OneModel); ok {
		ResponseSuccessWithCode(c, CodeSuccessOne, fortune)
	} else {
		ResponseSuccessWithCode(c, CodeSuccessMuti, fortune)
	}
}
