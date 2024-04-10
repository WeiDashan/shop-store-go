package pojo

type TestUser struct {
	Id          int64  `json:"id"`
	NickyName   string `json:"nickyName"`
	Email       string `json:"email"`
	Password    string `json:"-"`
	RawPassword string `json:"-" gorm:"-"`
}

//
//func (TestUser) TableName() string {
//	return "test_user"
//}
