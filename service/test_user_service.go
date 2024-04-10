package service

import (
	"errors"
	"github.com/WeiDashan/shop-go/dao"
	"github.com/WeiDashan/shop-go/global"
	"github.com/WeiDashan/shop-go/global/constants"
	"github.com/WeiDashan/shop-go/middleware/rabbitmq"
	"github.com/WeiDashan/shop-go/service/dto"
	"github.com/WeiDashan/shop-go/utils"
	"github.com/spf13/viper"
	"time"
)

var testUserService *TestUserService

type TestUserService struct {
	BaseService
	Dao *dao.TestUserDao
}

func NewTestUserService() *TestUserService {
	if testUserService == nil {
		testUserService = &TestUserService{
			Dao: dao.NewTestUserDao(),
		}
	}
	return testUserService
}

//func (m *AppUserService) GetAppUserByEmail(iUserDTO dto.UserLoginDTO) (pojo.AppUser, error) {
//	var errResult error
//	iAppUser := m.Dao.GetAppUserByEmail(iUserDTO.Email)
//	if iAppUser.Id == 0 {
//		errResult = errors.New("invalid Email")
//	}
//	return iAppUser, errResult
//}

func (m *TestUserService) AddTestUser(iTestUserAddDTO *dto.TestUserAddDTO) error {
	if m.Dao.CheckEmailExist(iTestUserAddDTO.Email) {
		return errors.New("email exist")
	}
	psw, err := utils.Encrypt(iTestUserAddDTO.Password)
	if err != nil {
		return errors.New("password encrypt error")
	}
	iTestUserAddDTO.Password = psw
	return m.Dao.AddTestUser(iTestUserAddDTO)
}
func (m *TestUserService) GetCodeByEmailToLogin(email string) error {

	// 判断邮箱是否绑定过现有用户
	if !m.Dao.CheckEmailExist(email) {
		return errors.New("邮箱未绑定用户")
	}
	// 判断在1分钟内是否申请过验证码
	key := constants.LoginCodeGenerate(email)
	v, _ := global.RedisClient.Get(key)
	if v != "" {
		return errors.New("在" + viper.GetString("loginCode.generateCodeExpire") + "分钟内请勿重复申请")
	}
	// 在Redis中标记正在生成验证码
	err := global.RedisClient.Set(key, 1, viper.GetDuration("loginCode.generateCodeExpire")*time.Minute)
	if err != nil {
		return errors.New("generate redis LOGIN_CODE_GENERATE error")
	}
	// 生成验证码
	code := utils.GenerateCode()
	// redis中进行标记
	_ = global.RedisClient.Set(email, code,
		viper.GetDuration("loginCode.codeExpire")*time.Minute)
	// rabbitMQ发送
	err = rabbitmq.EmailProducer(email, code)
	return err
}
