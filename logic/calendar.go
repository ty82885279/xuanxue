package logic

import (
	"xuanxue/dao/mysql"
	"xuanxue/model"
)

func CalendarDetail(date string) (*model.CalendarModel, error) {

	return mysql.CalendarDetail(date)

}
func GetAuspiciousList(p *model.ParamAuspiciousList) ([]*model.CalendarModel, int, error) {

	return mysql.GetAuspiciousList(p)

}
