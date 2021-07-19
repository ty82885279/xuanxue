package model

type AllModel struct {
	Today    *OneModel `json:"today" db:"type"`    // 今日运势
	Tomorrow *OneModel `json:"tomorrow" db:"type"` // 明日运势
	Week     *OneModel `json:"week" db:"type"`     // 本周运势
	Month    *OneModel `json:"month" db:"type"`    // 本年运势
	Year     *OneModel `json:"year" db:"type"`
}

type OneModel struct {
	Ntype      string `json:"-" db:"ntype"`
	Summary    string `json:"summary,omitempty" db:"summary"`       // 综合运势
	Love       string `json:"love" db:"love"`                       // 爱情运势
	Career     string `json:"career" db:"career"`                   // 工作运势
	Money      string `json:"money" db:"money"`                     // 财运运势
	Health     string `json:"health,omitempty" db:"health"`         // 健康运势
	Color      string `json:"color,omitempty" db:"color"`           // 幸运颜色
	Number     string `json:"number,omitempty" db:"number"`         // 幸运数字
	Star       string `json:"star,omitempty" db:"star"`             // 贵人星座
	Presummary string `json:"presummary,omitempty" db:"presummary"` // 概述
	Date       string `json:"-" db:"n_date"`
}
