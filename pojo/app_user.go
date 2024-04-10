package pojo

type AppUser struct {
	//gorm.Model
	Id          int64  `json:"id"`
	NickyName   string `json:"nickyName"`
	Email       string `json:"email"`
	Icon        string `json:"icon"`
	Phone       string `json:"phone"`
	Active      int8   `json:"-"`
	Password    string `json:"-"`
	RawPassword string `json:"-" gorm:"-"`
}
