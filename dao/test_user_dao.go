package dao

import (
	"github.com/WeiDashan/shop-go/pojo"
	"github.com/WeiDashan/shop-go/service/dto"
)

var testUserDao *TestUserDao

type TestUserDao struct {
	BaseDao
}

func NewTestUserDao() *TestUserDao {
	if testUserDao == nil {
		testUserDao = &TestUserDao{
			NewBaseDao(),
		}
	}
	return testUserDao
}

//	func (m *AppUserDao) GetAppUserByEmail(email string) pojo.AppUser {
//		var iAppUser pojo.AppUser
//		m.Orm.Model(&iAppUser).Where("email=?", email).Find(&iAppUser)
//		return iAppUser
//	}

func (m *TestUserDao) CheckEmailExist(email string) bool {
	//var nTotal int64
	//m.Orm.Model(&pojo.TestUser{}).Where("email=?", email).Count(&nTotal)
	//return nTotal > 0
	var testUser pojo.TestUser
	err := m.Orm.Model(&pojo.TestUser{}).Where("email=?", email).First(&testUser)
	if err.Error != nil {
		return false
	}
	return true
}

func (m *TestUserDao) AddTestUser(iTestAddUserDTO *dto.TestUserAddDTO) error {
	var iTestUser pojo.TestUser
	iTestAddUserDTO.ConvertToPojo(&iTestUser)
	err := m.Orm.Save(&iTestUser).Error
	if err == nil {
		iTestAddUserDTO.Id = iTestUser.Id
	}
	return err
}
