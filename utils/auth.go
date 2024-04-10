package utils

import (
	"fmt"
	"github.com/WeiDashan/shop-go/global"
	"github.com/WeiDashan/shop-go/global/constants"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strconv"
	"time"
)

const (
	TOKEN_NAME = "token"
)

func tokenErr(c *gin.Context) {
	RespError(c, "Invalid Token")
}
func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		if constants.NeedFilter(c.FullPath()) {
			token := c.GetHeader(TOKEN_NAME)
			if token == "" {
				tokenErr(c)
				return
			}
			iJwtCustClaims, err := ParseToken(token)
			Id := iJwtCustClaims.Id
			if err != nil || Id == 0 {
				tokenErr(c)
				return
			}

			// token与访问者登录对应的token不一致，直接返回
			stringId := strconv.Itoa(int(Id))
			redisId, err := global.RedisClient.Get(token)
			if err != nil || stringId != redisId {
				tokenErr(c)
				return
			}
			// token续期
			err = global.RedisClient.Set(token, redisId, viper.GetDuration("jwt.tokenExpire")*time.Minute)
			if err != nil {
				RespError(c, fmt.Sprintf("token续期失败：err: %s", err.Error()))
				return
			}
		}

	}
}
