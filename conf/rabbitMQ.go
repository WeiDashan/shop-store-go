package conf

import (
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"reflect"
)

var RabbitConn *amqp.Connection
var RabbitChannel *amqp.Channel

type RabbitMQClient struct {
}

func InitRabbitMQ() (*RabbitMQClient, error) {
	var err error
	// const MQURL = "amqp://kuteng:kuteng@127.0.0.1:5672/kuteng"
	var RabbitMQURL = "amqp://" + viper.GetString("rabbitmq.username") + ":" +
		viper.GetString("rabbitmq.password") + "@" +
		viper.GetString("rabbitmq.host") + ":" +
		viper.GetString("rabbitmq.port")
	fmt.Println(RabbitMQURL)
	RabbitConn, err = amqp.Dial(RabbitMQURL)
	if err != nil {
		err = errors.New(fmt.Sprintf("连接RabbitMQ失败,error: %s", err.Error()))
		panic(err)
		return nil, err
	}
	RabbitChannel, err = RabbitConn.Channel()
	if err != nil {
		err = errors.New(fmt.Sprintf("获取RabbitMQ Channel失败,error: %s", err.Error()))
		panic(err)
		return nil, err
	}
	return &RabbitMQClient{}, nil
}
func (rbMQ *RabbitMQClient) GetChannel() *amqp.Channel {
	return RabbitChannel
}
func (rbMQ *RabbitMQClient) CheckRabbitClosed() int64 {
	d := reflect.ValueOf(RabbitChannel)
	i := d.FieldByName("closed").Int()
	return i
}
