package dualarm

import (
	"git.du.com/cloud/du_component/dugorm"
	"time"
)

const (
	TableName = "du_alarm_stat"
)

type Report struct {
	ReportTime time.Time `gorm:"column:report_time;type:datetime"`
	ReportDate string    `gorm:"column:report_date;type:date"`
	AppName    string    `gorm:"column:app_name"`
	Int1       int64     `gorm:"column:int_1"`
	Int2       int64     `gorm:"column:int_2"`
	Int3       int64     `gorm:"column:int_3"`
	Int4       int64     `gorm:"column:int_4"`
	Int5       int64     `gorm:"column:int_5"`
	Int6       int64     `gorm:"column:int_6"`
	Float1     float64   `gorm:"column:float_1;type:decimal(10,2)"`
	Float2     float64   `gorm:"column:float_2;type:decimal(10,2)"`
	Float3     float64   `gorm:"column:float_3;type:decimal(10,2)"`
	String1    string    `gorm:"column:string_1"`
	String2    string    `gorm:"column:string_2"`
	String3    string    `gorm:"column:string_3"`
	String4    string    `gorm:"column:string_4"`
	String5    string    `gorm:"column:string_5"`
	String6    string    `gorm:"column:string_6"`
}

func Init() {
	initBaseGormDb(dugorm.Config{
		Address:  "rm-2zewy8sohv5a02jqv.mysql.rds.aliyuncs.com:3306",
		Username: "du_alarm",
		Password: "8Lt4OuFDmC1gAP3H",
		Database: "du_alarm",
	})
}

func AddReport(report *Report) {
	var reportMysql ReportMysql
	reportMysql.Report = *report
	reportMysql.CreatedAt = time.Now()
	reportMysql.UpdatedAt = time.Now()
	BaseGormDb.Table(TableName).Create(&reportMysql)
}

type ReportMysql struct {
	Id        int64     `gorm:"column:id"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime"`
	Report
}
