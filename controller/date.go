package controller

import (
	"github.com/gin-gonic/gin"
	"xuanxue/logic"
)

func DataHandler(c *gin.Context) {

	//yM := make(map[string][]map[string][]string)
	//make(map[string]map[string][]string)
	//mM := make(map[string][]string)
	//
	//mS1 := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	//mS2 := []string{"11", "22", "33", "44", "55", "66", "77", "88", "99", "110"}
	//mmS := make([]map[string][]string, 0, 100)
	//
	//mM["1月"] = mS1
	//mM["2月"] = mS2
	//mmS = append(mmS, mM)
	//yM["2020"] = mmS
	//yM["2021"] = mmS
	//
	//b, err := json.Marshal(yM)
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//fmt.Printf("结果=%s\n", string(b))
	//c.String(200, string(b))
	date := logic.GetDate()
	c.String(200, date)
}
