package dualarm

import (
	"git.du.com/cloud/du_component/dugorm"
	"gorm.io/gorm"
)

var (
	BaseGormDb *gorm.DB
)

func initBaseGormDb(cfg dugorm.Config) {
	BaseGormDb = dugorm.NewGormDb(cfg)
}
