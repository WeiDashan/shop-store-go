package rabbitmq

import (
	"encoding/json"
	"errors"
	"github.com/WeiDashan/shop-go/global"
	"github.com/streadway/amqp"
)

func EmailProducer(to, code string) error {
	var queueName = "email"
	var err error
	//declare, err := global.RabbitMQClient.GetChannel().QueueDeclare(queueName,
	//	true, false, false, false, nil)
	email := EmailDTO{Subject: "获取验证码", To: to, Message: code}
	if err != nil {
		err = errors.New("声明email队列失败")
		panic(err)
		return err
	}
	marshal, _ := json.Marshal(email)
	err = global.RabbitMQClient.GetChannel().Publish("", queueName,
		false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(marshal),
		})
	if err != nil {
		err = errors.New("邮件生产者发送消息失败")
	}
	return err
}
