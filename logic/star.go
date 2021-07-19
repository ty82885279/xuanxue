package logic

import (
	"xuanxue/dao/mysql"
	"xuanxue/model"
)

func GetStarFortune(p *model.ParamStar) (interface{}, error) {

	if p.Type == "" { //all
		return GetAllFortune(p)
	} else { //one
		return GetTimeFortune(p)
	}
}

func GetAllFortune(p *model.ParamStar) (*model.AllModel, error) {

	return mysql.GetAllFortune(p)

}
func GetTimeFortune(p *model.ParamStar) (*model.OneModel, error) {
	return mysql.GetTimeFortune(p)
}
