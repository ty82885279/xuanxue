package controller

import "xuanxue/model"

type _ResponseOneFortune struct {
	Code    ResCode         `json:"code"` // 状态码
	Message string          `json:"msg"`  // 提示信息
	Data    *model.OneModel `json:"data"` // 数据
}
type _ResponseAllFortune struct {
	Code    ResCode         `json:"code"` // 状态码
	Message string          `json:"msg"`  // 提示信息
	Data    *model.AllModel `json:"data"` // 数据
}

type _ResponseCalendar struct {
	Code    ResCode              `json:"code"` // 状态码
	Message string               `json:"msg"`  // 提示信息
	Data    *model.CalendarModel `json:"data"` // 数据
}
type _ResponseAuspicious struct {
	Code    ResCode                `json:"code"` // 状态码
	Message string                 `json:"msg"`  // 提示信息
	Data    []*model.CalendarModel `json:"data"` // 数据
}
