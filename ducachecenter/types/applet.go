package types

type ProductApplet struct {
	AppId     string `gorm:"column:app_id" json:"app_id"`
	Name      string `gorm:"column:name" json:"name"`
	Cid       int    `gorm:"column:cid" json:"cid"`
	SecretKey string `gorm:"secret_key" json:"secret_key"`
	UserName  string `gorm:"user_name" json:"user_name"`
	Cheating  int    `gorm:"cheating" json:"cheating"`
}
