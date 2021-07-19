package mysql

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
)

type NLModel struct {
	NLStr string
}

func GetDate() (s string) {

	//endY := 2000
	////dateSlice := make([]string, 0, 38000)
	//for y := 1999; y <= endY; y++ {
	//
	//	for m := 1; m <= 12; m++ {
	//		days := pkg.GetDays(y, m)
	//		for d := 1; d <= days; d++ {
	//			ystr := strconv.Itoa(y)
	//			mstr := strconv.Itoa(m)
	//			dstr := strconv.Itoa(d)
	//
	//			nlStr := ""
	//			sqlStr := `SELECT nongli FROM hl WHERE myear = ? AND mmonth = ? AND mday = ?`
	//			err1 := DB.Get(&nlStr, sqlStr, ystr, mstr, dstr)
	//			if err1 != nil {
	//				//panic(err1)
	//				continue
	//			}
	//			bbbstr := []rune(nlStr)
	//			nnstr1 := string(bbbstr[2:6])
	//			nnstr2 := strings.Split(strings.Split(nlStr, `年`)[1], `月`)[0]
	//			nnstr3 := strings.Split(strings.Split(nlStr, `年`)[1], `月`)[1]
	//			fmt.Printf("%s-%s-%s\n", nnstr1, nnstr2, nnstr3)
	//			//UPDATE table_name
	//			//SET column1=value1,column2=value2,...
	//			//WHERE some_column=some_value;
	//			sqlStr2 := `update hl set nlyear = ? ,nlmonth = ?,nlday = ? where myear = ? AND mmonth = ? AND mday = ? `
	//			_, err2 := DB.Exec(sqlStr2, nnstr1, nnstr2, nnstr3, ystr, mstr, dstr)
	//			if err2 != nil {
	//				continue
	//			}
	//
	//		}
	//	}
	//}

	//

	//ss := make([]*model.NongM, 0, 1000)
	//for _, v := range dateS {
	//	n := new(model.NongM)
	//	strYi := []rune(v)
	//	n.YearStr = string(strYi[2:6])
	//	n.MonStr = strings.Split(strings.Split(v, `年`)[1], `月`)[0]
	//	n.DayStr = strings.Split(strings.Split(v, `年`)[1], `月`)[1]
	//	ss = append(ss, n)
	//}

	//endY := 2021
	//
	//map2 := make(map[string][]map[string][]string)
	////daySlice := make([]string, 0, 35)
	//map1 := make(map[string][]string)
	//
	//for y := 2020; y <= endY; y++ {
	//
	//	m1 := make([]map[string][]string, 0, 20)
	//	for m := 1; m <= 12; m++ {
	//
	//		days := pkg.GetDays(y, m)
	//		daySlice := make([]string, 0, 35)
	//		for d := 1; d <= days; d++ {
	//
	//			daySlice = append(daySlice, strconv.Itoa(d))
	//			fmt.Printf("天数--%d    ", d)
	//		}
	//		fmt.Println("")
	//		map1[strconv.Itoa(m+1000)] = daySlice
	//		fmt.Printf("月份--%d    \n", m)
	//	}
	//
	//	m1 = append(m1, map1)
	//	map2[strconv.Itoa(y)] = m1
	//
	//}
	//b, err := json.Marshal(map2)
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//bb := string(b)
	//fmt.Println(bb)
	//return bb

	//return mapJson()
	return get()
}

// 农历
func getNl() string {
	nlyear := make([]string, 0, 67) //得到所有农历年份
	sqlstr := `SELECT DISTINCT nlyear FROM hl limit 0,67`
	err := DB.Select(&nlyear, sqlstr)
	if err != nil {
		panic(err)
	}

	m1 := make(map[string]interface{})
	for i, v := range nlyear {

		nlmonth := make([]string, 0, 20) //得到所有农历月份
		sqlstr1 := `SELECT DISTINCT nlmonth FROM hl where nlyear = ? `
		err := DB.Select(&nlmonth, sqlstr1, v)
		if err != nil {
			panic(err)
		}
		c := i + 2000

		m1[strconv.Itoa(c)+v] = nlmonth
	}
	var testSlice []string
	for k, _ := range m1 {

		testSlice = append(testSlice, k)
	}

	sort.Strings(testSlice)
	fmt.Printf("%v\n", testSlice)
	/////

	for _, Key := range testSlice {
		/* 按顺序从MAP中取值输出 */
		if Value, ok := m1[Key]; ok {
			//fmt.Println(Key, ":", Value)
			//fmt.Printf("%T\n", Value)
			stingS, ok := Value.([]string)
			if !ok {
				panic("断言错误")
			}
			mm := make(map[string]interface{})
			for i, v := range stingS {
				//fmt.Printf("%d:%s\n", i, v)
				daySlice := make([]string, 0, 35)
				tk := []rune(Key)
				tkStr := string(tk[4:8])

				sqlstr := `select nlday from hl where nlyear = ? and nlmonth = ?`
				err := DB.Select(&daySlice, sqlstr, tkStr, v)
				if err != nil {
					panic(err)
				}
				//fmt.Printf("农历%s年%s月日子数组：%#v\n", tkStr, v, daySlice)

				tempS := make([]string, 0, 35)
				tempS = daySlice

				mm[strconv.Itoa(i+2000)+v] = tempS
			}
			for k, v := range mm {
				fmt.Printf("-===- %s:%v\n", k, v)
			}
			ss1 := make([]interface{}, 0, 1)
			ss1 = append(ss1, mm)
			m1[Key] = ss1
			fmt.Printf("=====%s===这个JB年结束了========\n", Key)
		}
	}

	b, err := json.Marshal(m1)
	if err != nil {
		panic(err)
		//return
	}
	bb := string(b)
	return bb
}

// 阳历
func get() string {
	nlyear := make([]string, 0, 67) //得到所有农历年份
	sqlstr := `SELECT DISTINCT myear FROM hl limit 0,67`
	err := DB.Select(&nlyear, sqlstr)
	if err != nil {
		panic(err)
	}

	m1 := make(map[string]interface{})
	for _, v := range nlyear {

		nlmonth := make([]string, 0, 20) //得到所有农历月份
		sqlstr1 := `SELECT DISTINCT mmonth FROM hl where myear = ? `
		err := DB.Select(&nlmonth, sqlstr1, v)
		if err != nil {
			panic(err)
		}
		//c := i + 2000

		m1[v] = nlmonth
	}
	var testSlice []string
	for k, _ := range m1 {

		testSlice = append(testSlice, k)
	}

	sort.Strings(testSlice)
	fmt.Printf("-- %v\n", testSlice)
	/////

	for _, Key := range testSlice {
		/* 按顺序从MAP中取值输出 */
		if Value, ok := m1[Key]; ok {
			//fmt.Println(Key, ":", Value)
			//fmt.Printf("%T\n", Value)
			stingS, ok := Value.([]string)
			if !ok {
				panic("断言错误")
			}
			mm := make(map[string]interface{})
			for _, v := range stingS {
				//fmt.Printf("%d:%s\n", i, v)
				daySlice := make([]string, 0, 35)
				//tk := []rune(Key)
				//tkStr := string(tk[4:8])

				sqlstr := `select mday from hl where myear = ? and mmonth = ?`
				err := DB.Select(&daySlice, sqlstr, Key, v)
				if err != nil {
					panic(err)
				}
				//fmt.Printf("农历%s年%s月日子数组：%#v\n", tkStr, v, daySlice)

				tempS := make([]string, 0, 35)
				tempS = daySlice
				atoi, _ := strconv.Atoi(v)
				mm[strconv.Itoa(1000+atoi)] = tempS
			}
			for k, v := range mm {
				fmt.Printf("-===- %s:%v\n", k, v)
			}
			ss1 := make([]interface{}, 0, 1)
			ss1 = append(ss1, mm)
			m1[Key] = ss1
			fmt.Printf("=====%s===这个JB年结束了========\n", Key)
		}
	}

	b, err := json.Marshal(m1)
	if err != nil {
		panic(err)
		//return
	}
	bb := string(b)
	return bb
}

type MapM struct {
	Date   string `db:"date"`
	Nyear  string `db:"nlyear"`
	Nmonth string `db:"nlmonth"`
	Nday   string `db:"nlday"`
}

func mapJson() string {

	s1 := make([]*MapM, 0, 38000)
	sqlStr := `select date,nlyear,nlmonth,nlday from hl WHere Date(date) between '1970-01-01' AND '2036-12-31'`
	err := DB.Select(&s1, sqlStr)

	if err != nil {
		panic(err)
	}
	fmt.Println(s1)

	mm := make(map[string]string)
	for _, v := range s1 {
		fmt.Printf("v:%s", v)
		mm[v.Nyear+v.Nmonth+v.Nday] = v.Date
	}
	b, err := json.Marshal(mm)
	if err != nil {
		panic(err)

	}

	return string(b)
}
