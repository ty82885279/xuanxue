package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"strconv"
	"xuanxue/model"
)

const (
	TODAY    = "today"
	TOMORROW = "tomorrow"
	WEEK     = "week"
	MONTH    = "month"
	YEAR     = "year"
)

var ErrNoRows = errors.New("没有数据")

func GetAllFortune(p *model.ParamStar) (m *model.AllModel, err error) {

	m = new(model.AllModel)
	var FortuneS = make([]*model.OneModel, 0, 5)
	sqlStr := fmt.Sprintf(`select ntype,summary,love,career,money,health,color,number,star,presummary from %s where date = ?`, p.Star)

	err = DB.Select(&FortuneS, sqlStr, p.Date)
	if err != nil {
		return nil, err
	}
	if len(FortuneS) == 0 {
		return nil, ErrNoRows
	}
	for _, v := range FortuneS {

		switch v.Ntype {
		case TODAY:
			m.Today = handleDecimal(v)
		case TOMORROW:
			m.Tomorrow = handleDecimal(v)
		case WEEK:
			m.Week = handleDecimal(v)
		case MONTH:
			m.Month = handleDecimal(v)
		case YEAR:
			m.Year = handleDecimal(v)
		}
	}
	return
}

func GetTimeFortune(p *model.ParamStar) (m *model.OneModel, err error) {
	m = new(model.OneModel)

	sqlStr := fmt.Sprintf(`select ntype,summary,love,career,money,health,color,number,star,presummary from %s where date = ? and ntype = ? `, p.Star)

	err = DB.Get(m, sqlStr, p.Date, p.Type)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoRows
		}
		return nil, err
	}

	//调用处理函数
	m = handleDecimal(m)

	return
}

func handleDecimal(m *model.OneModel) *model.OneModel {
	if m.Ntype == TODAY || m.Ntype == TOMORROW {
		sum, err := strconv.ParseFloat(m.Summary, 32)
		if err != nil {
			zap.L().Error("GetTimeFortune.strconv.ParseFloat", zap.Error(err))
		}
		if sum < 1 {
			sum = sum * 100
		}
		m.Summary = strconv.FormatFloat(sum, 'g', 3, 32)

		love, err := strconv.ParseFloat(m.Love, 32)
		if err != nil {
			zap.L().Error("GetTimeFortune.strconv.ParseFloat", zap.Error(err))
		}
		if love < 1 {
			love = love * 100
		}
		m.Love = strconv.FormatFloat(love, 'g', 3, 32)
		career, err := strconv.ParseFloat(m.Career, 32)
		if err != nil {
			zap.L().Error("GetTimeFortune.strconv.ParseFloat", zap.Error(err))
		}
		if career < 1 {
			career = career * 100
		}
		m.Career = strconv.FormatFloat(career, 'g', 3, 32)

		money, err := strconv.ParseFloat(m.Money, 32)
		if err != nil {
			zap.L().Error("GetTimeFortune.strconv.ParseFloat", zap.Error(err))
		}
		if money < 1 {
			money = money * 100
		}
		m.Money = strconv.FormatFloat(money, 'g', 3, 32)

		health, err := strconv.ParseFloat(m.Health, 32)
		if err != nil {
			zap.L().Error("GetTimeFortune.strconv.ParseFloat", zap.Error(err))
		}
		if health < 1 {
			health = health * 100
		}
		m.Health = strconv.FormatFloat(health, 'g', 3, 32)

	}
	return m
}
