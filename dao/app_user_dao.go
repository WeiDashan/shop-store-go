package dao

import "github.com/WeiDashan/shop-go/pojo"

var appUserDao *AppUserDao

type AppUserDao struct {
	BaseDao
}

func NewAppUserDao() *AppUserDao {
	if appUserDao == nil {
		appUserDao = &AppUserDao{
			NewBaseDao(),
		}
	}
	return appUserDao
}

func (m *AppUserDao) GetAppUserByEmail(email string) pojo.AppUser {
	var iAppUser pojo.AppUser
	m.Orm.Model(&iAppUser).Where("email=?", email).Find(&iAppUser)
	return iAppUser
}
func (m *AppUserDao) CheckEmailExist(email string) bool {
	//var nTotal int64
	//m.Orm.Model(&pojo.TestUser{}).Where("email=?", email).Count(&nTotal)
	//return nTotal > 0
	var appUser pojo.AppUser
	err := m.Orm.Model(&pojo.AppUser{}).Where("email=?", email).First(&appUser)
	if err.Error != nil {
		return false
	}
	return true
}
