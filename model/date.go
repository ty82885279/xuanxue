package model

// make(map[string][]map[string][]string)
type YearM struct {
	Year []map[string][]MonthM
}

type MonthM struct {
	Month map[string]DayS
}

type DayS struct {
	Day []string
}

type NongM struct {
	YearStr string
	MonStr  string
	DayStr  string
}
