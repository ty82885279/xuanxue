package pkg

import (
	"encoding/json"
	"fmt"
)

func GetDays(year, month int) (day int) {

	switch month {

	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	default:
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			return 29
		} else {
			return 28
		}
	}

}

func JsonToMap(jsonStr string) map[string]interface{} {

	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
		panic(err)
	}

	return mapResult
}
