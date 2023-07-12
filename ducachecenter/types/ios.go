package types

type CustomerIos struct {
	Id        int    `gorm:"primaryKey" json:"id"`    // id
	Name      string `gorm:"column:name" json:"name"` // pkgName
	TableName string `gorm:"column:tableName" json:"table_name"`
	IsActive  int    `gorm:"column:isActive" json:"is_active"`
}

type ProductIos struct {
	Id      int    `gorm:"primaryKey" json:"id"`                // id
	PkgName string `gorm:"column:package_name" json:"pkg_name"` // pkgName
	AppName string `gorm:"column:product_name" json:"app_name"` // AppName
	Cid     int    `gorm:"column:cp_customer_id" json:"cid"`    // 客户id
}

type IdMap struct {
	AndId int `gorm:"column:acid" json:"and_id"` //安卓客户id
	IosId int `gorm:"column:icid" json:"ios_id"` //Ios客户id
}

type CustomerInfoByPkg struct {
	PkgName   string `gorm:"column:pkg_name" json:"pkg_name"` // 包名
	IosCid    string `gorm:"column:ios_cid" json:"ios_cid"`   // ios客户id
	SecretKey string `gorm:"secret_key" json:"secret_key"`    // 由cdid获取did的密钥
}
