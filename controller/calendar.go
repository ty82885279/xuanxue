package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"xuanxue/dao/mysql"
	"xuanxue/logic"
	"xuanxue/model"
)

// CalendarHandler 黄历查询
// @Summary 黄历查询
// @Description 黄历查询
// @Tags 黄历相关
// @Accept application/json
// @Produce application/json
// @Param date query string true "日期：例如：2020-11-11"
// @Success 2000 {object} _ResponseCalendar "黄历详情"
// @Router /calendar [get]
func CalendarHandler(c *gin.Context) {

	date := c.Query("date")
	detail, err := logic.CalendarDetail(date)
	if err != nil {
		if err == mysql.ErrNoRows {
			zap.L().Error("CalendarDetail err", zap.Error(err))
			ResponseErr(c, CodeNoData)
			return
		}
		ResponseErr(c, CodeServiceBusy)
		return
	}
	ResponseSuccess(c, detail)
}

// AuspiciousHandler 择吉日
// @Summary 择吉日查询
// @Description 择吉日查询
// @Tags 黄历相关
// @Accept application/json
// @Produce application/json
// @Param object query model.ParamAuspiciousList true "查询参数"
// @Success 2000 {object} _ResponseAuspicious "吉日列表"
// @Router /calendar/auspicious [get]
func AuspiciousHandler(c *gin.Context) {
	var p = &model.ParamAuspiciousList{
		Weekend: false,
		Page:    1,
		Size:    20,
	}
	err := c.ShouldBind(p)
	if err != nil {
		ResponseErr(c, CodeInvalidParam)
		//zap.L().Error("AuspiciousHandler invalid param", zap.Error(err))
		return
	}
	//fmt.Printf("参数：%#v\n", p)
	//list, err := logic.GetAuspiciousYiList(p)
	list, total, err := logic.GetAuspiciousList(p)
	if err != nil {
		ResponseErr(c, CodeServiceBusy)
		return
	}
	type Temp struct {
		Total    int         `json:"total"`
		DataList interface{} `json:"list"`
	}
	ResponseSuccess(c, Temp{
		Total:    total,
		DataList: list,
	})

}
