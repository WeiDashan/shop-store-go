package test

//import (
//	"github.com/streadway/amqp"
//	"log"
//)

//package main

//func main() {
// This example acts as a bridge, shoveling all messages sent from the source
// exchange "log" to destination exchange "log".

// Confirming publishes can help from overproduction and ensure every message
// is delivered.

// Setup the source of the store and forward
//source, _ := amqp.Dial("amqp://source/")
//defer source.Close()
//
//chs, err := source.Channel()
//
//shovel, _ := chs.Consume("remote-tee", "shovel", false, false, false, false, nil)
//
//destination, _ := amqp.Dial("amqp://destination/")
//defer destination.Close()
//chd, _ := destination.Channel()
//
//confirms := chd.NotifyPublish(make(chan amqp.Confirmation, 1))
//
//_= chd.Confirm(false)
//
//
//for {
//	msg, _ := <-shovel
//	err = chd.Publish("logs", msg.RoutingKey, false, false, amqp.Publishing{
//		// Copy all the properties
//		ContentType:     msg.ContentType,
//		ContentEncoding: msg.ContentEncoding,
//		DeliveryMode:    msg.DeliveryMode,
//		Priority:        msg.Priority,
//		CorrelationId:   msg.CorrelationId,
//		ReplyTo:         msg.ReplyTo,
//		Expiration:      msg.Expiration,
//		MessageId:       msg.MessageId,
//		Timestamp:       msg.Timestamp,
//		Type:            msg.Type,
//		UserId:          msg.UserId,
//		AppId:           msg.AppId,
//
//		// Custom headers
//		Headers: msg.Headers,
//
//		// And the body
//		Body: msg.Body,
//	})
//
//	if confirmed := <-confirms; confirmed.Ack {
//		msg.Ack(false)
//	} else {
//		msg.Nack(false, false)
//	}
//}
//}
