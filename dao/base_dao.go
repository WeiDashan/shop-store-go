package dao

import (
	"github.com/WeiDashan/shop-go/global"
	"gorm.io/gorm"
)

type BaseDao struct {
	Orm *gorm.DB
}

func NewBaseDao() BaseDao {
	return BaseDao{
		Orm: global.DB,
	}
}
