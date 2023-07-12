package types

// CustomerAnd represents a row from 'aicnt_core_db.product'.
type CustomerAnd struct {
	Id        int    `gorm:"primaryKey" json:"id"`    // id
	Name      string `gorm:"column:name" json:"name"` // pkgName
	TableName string `gorm:"column:tableName" json:"table_name"`
	IsActive  int    `gorm:"column:isActive" json:"is_active"`
}

type ProductAnd struct {
	Id      int    `gorm:"primaryKey" json:"id"`           // id
	PkgName string `gorm:"column:pkgName" json:"pkg_name"` // pkgName
	AppName string `gorm:"column:name" json:"app_name"`    // AppName
	Cid     int    `gorm:"column:cid" json:"cid"`          // 客户id
}
