package dto

import "github.com/WeiDashan/shop-go/pojo"

type TestUserAddDTO struct {
	Id          int64  `json:"id" form:"id"`
	NickyName   string `json:"nickyName" form:"nickyName"`
	Email       string `json:"email" form:"email"`
	Password    string `json:"password" form:"-"`
	RawPassword string `json:"-" form:"password"`
}

func (m *TestUserAddDTO) ConvertToPojo(iTestUser *pojo.TestUser) {
	iTestUser.NickyName = m.NickyName
	iTestUser.Email = m.Email
	iTestUser.Password = m.Password
}
