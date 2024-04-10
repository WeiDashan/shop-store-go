package global

import (
	"github.com/WeiDashan/shop-go/conf"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger         *zap.SugaredLogger
	DB             *gorm.DB
	RedisClient    *conf.RedisClient
	RabbitMQClient *conf.RabbitMQClient
)
