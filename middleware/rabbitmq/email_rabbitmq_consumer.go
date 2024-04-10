package rabbitmq

import (
	"encoding/json"
	"github.com/WeiDashan/shop-go/global"
)

func EmailConsumer() {
	var queueName = "email"
	_ = global.RabbitMQClient.GetChannel().Qos(
		1,
		0,
		false)

	deliveries, _ := global.RabbitMQClient.GetChannel().Consume(queueName,
		"", false, false,
		false, false, nil)

	for d := range deliveries {
		msgData := string(d.Body)
		emailDto := EmailDTO{}
		_ = json.Unmarshal([]byte(msgData), &emailDto)
		// 发送验证码邮件
		_ = SendEmail(emailDto)
		_ = d.Ack(false)
	}

	//forever := make(chan bool)
	//go func() {
	//
	//}()
	//<-forever
}
