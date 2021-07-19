package model

type CalendarModel struct {
	Date       string `json:"date" db:"date"`             // 日期
	Year       string `json:"year" db:"year"`             // 日期
	Month      string `json:"month" db:"month"`           // 日期
	Day        string `json:"day" db:"day"`               // 日期
	Days       int    `json:"days"`                       // 还有几天
	Yangli     string `json:"yangli" db:"yangli"`         // 阳历
	Nongli     string `json:"nongli" db:"nongli"`         // 农历
	Star       string `json:"star" db:"star"`             // 星座
	Taishen    string `json:"taishen" db:"taishen"`       // 胎神
	Wuxing     string `json:"wuxing" db:"wuxing"`         // 五行
	Chong      string `json:"chong" db:"chong"`           // 冲
	Sha        string `json:"sha" db:"sha"`               // 煞
	Shengxiao  string `json:"shengxiao" db:"shengxiao"`   // 生肖
	Jiri       string `json:"jiri" db:"jiri"`             // 吉日
	Zhiri      string `json:"zhiri" db:"zhiri"`           // 值日
	Xiongshen  string `json:"xiongshen" db:"xiongshen"`   // 凶神
	Jishenyiqu string `json:"jishenyiqu" db:"jishenyiqu"` // 吉神宜趋
	Caishen    string `json:"caishen" db:"caishen"`       // 财神
	Xishen     string `json:"xishen" db:"xishen"`         // 喜神
	Fushen     string `json:"fushen" db:"fushen"`         // 福神
	Eweek      string `json:"eweek" db:"eweek"`           // 英文周
	Emonth     string `json:"emonth" db:"emonth"`         // 英文月
	Week       string `json:"week" db:"week"`             // 中文周
	Suici      string `json:"suici" db:"suici"`           // 岁次
	Yi         string `json:"yi" db:"yi"`                 // 宜
	Ji         string `json:"ji" db:"ji"`                 // 忌
}
