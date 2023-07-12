package ducounter

import "time"

type Info struct {
	Time  time.Time `json:"time" gorm:"column:time"`
	Name  string    `json:"name" gorm:"column:name"`
	Count int64     `json:"count" gorm:"column:count"`
}

type TableNames struct {
	DnaTable string `json:"dna_table"`
	DdiTable string `json:"ddi_table"`
}
