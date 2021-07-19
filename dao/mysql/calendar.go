package mysql

import (
	"database/sql"
	"go.uber.org/zap"
	"time"
	"xuanxue/model"
)

type DBModel struct {
	ID    int64  `db:"id"`
	Year  string `db:"year"`
	Month string `db:"month"`
	Day   string `db:"day"`
}

func CalendarDetail(date string) (calendar *model.CalendarModel, err error) {
	calendar = new(model.CalendarModel)

	sqtStr := `SELECT date,yangli,nongli,star,taishen,wuxing,chong,sha,shengxiao,jiri,zhiri,xiongshen,jishenyiqu,
caishen,xishen,fushen,suici,yi,ji,eweek,emonth,week FROM huangli WHERE date = ?`
	err = DB.Get(calendar, sqtStr, date)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoRows
		}
		zap.L().Error("CalendarDetail err", zap.Error(err))
		return nil, err
	}
	//去掉多余字串
	strYi := []rune(calendar.Yi)
	if string(strYi[0]) == `【` {
		calendar.Yi = string(strYi[12:])
	}
	strJi := []rune(calendar.Ji)
	if string(strJi[0]) == `【` {
		calendar.Ji = string(strJi[12:])
	}
	//计算多少天
	t := time.Now()
	PT, err := time.Parse("2006-01-02", date)
	if err != nil {
		return
	}
	//nowT := t.Format("2006-01-02")
	duration := PT.Sub(t)
	//fmt.Println(duration)
	hours := duration.Hours()

	if (int(hours))%24 == 0 {
		calendar.Days = int(hours) / 24
	} else {
		calendar.Days = int(hours)/24 + 1
	}

	return
}
func GetAuspiciousList(p *model.ParamAuspiciousList) (calendarList []*model.CalendarModel, count int, err error) {
	var days int64
	switch p.Dur {
	case "week":
		days = 7
	case "month":
		days = 30
	case "month3":
		days = 90
	default:
		days = 365
	}
	t := time.Now()
	weekTime := t.Add((time.Hour * 24 * time.Duration(days)))
	nowT := t.Format("2006-01-02")
	endT := weekTime.Format("2006-01-02")
	//fmt.Println(endT)
	sqlStr1 := `SELECT date,yangli,nongli,star,taishen,wuxing,chong,sha,shengxiao,jiri,zhiri,xiongshen,jishenyiqu,
caishen,xishen,fushen,suici,yi,ji,eweek,emonth,week FROM huangli WHERE ` + p.Type + ` LIKE ` + `'%` + p.Name + `%'` + `And Date(date) between '` + nowT + `'` + `and '` + endT + `'`
	sqlWeek := `AND week IN ('周六','周日')`
	sqlStr := `order by id limit ?,? `

	//查询总数的sql
	sqlCount := `SELECT count(date) FROM huangli WHERE ` + p.Type + ` LIKE ` + `'%` + p.Name + `%'` + `And Date(date) between '` + nowT + `'` + `and '` + endT + `'`

	//拼接是否只看周末的sql
	if p.Weekend {
		sqlStr = sqlStr1 + sqlWeek + sqlStr
		sqlCount = sqlCount + sqlWeek
	} else {
		sqlStr = sqlStr1 + sqlStr
	}
	//只在第1页返回总天数
	if p.Page == 1 {
		err = DB.Get(&count, sqlCount)
		if err != nil && err != sql.ErrNoRows {
			panic(err)
			return
		}
	}
	//fmt.Printf("天数---%d\n", count)
	calendarList = make([]*model.CalendarModel, 0, 20)
	err = DB.Select(&calendarList, sqlStr, (p.Page-1)*p.Size, p.Size)
	if err != nil {
		panic(err)
	}
	//计算还差多少天
	for _, calendar := range calendarList {

		t1, err := time.Parse("2006-01-02", calendar.Date)
		if err != nil {
			zap.L().Error("t1 time.Parse err", zap.Error(err))
			continue
		}
		t2, err := time.Parse("2006-01-02", nowT)
		if err != nil {
			zap.L().Error("t2 time.Parse err", zap.Error(err))
			continue
		}
		durtion := t1.Sub(t2)
		hours := durtion.Hours()

		if (int(hours))%24 == 0 {
			calendar.Days = int(hours) / 24
		} else {
			calendar.Days = int(hours)/24 + 1
		}
		//fmt.Println(calendar.Days)

		//去掉多余字串
		strYi := []rune(calendar.Yi)
		if string(strYi[0]) == `【` {
			calendar.Yi = string(strYi[12:])
		}
		strJi := []rune(calendar.Ji)
		if string(strJi[0]) == `【` {
			calendar.Ji = string(strJi[12:])
		}

	}
	//fmt.Println(calendarList)
	return
}
