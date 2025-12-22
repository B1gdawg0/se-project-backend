package rabbit

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/B1gdawg0/se-project-backend/internal/transaction/response"
	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	EXCHANGE_NAME = "logs_direct"
)

var rabbitConn *amqp.Connection
var rabbitCh *amqp.Channel
var rabbitEmailQueue amqp.Queue

func LogNormalFailed(err error, msg string){
	if err != nil {
	  log.Printf("%s: %s", msg, err)
	}
}

func LogSeriousFailed(err error, msg string){
	if err != nil{
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func InitRabbitMQConnection(){
	var err error
	rabbitConn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	LogSeriousFailed(err, "Can't connect to message queue")


	rabbitCh ,err = rabbitConn.Channel()
	LogSeriousFailed(err, "Can't open channel")

	err = rabbitCh.ExchangeDeclare(
		EXCHANGE_NAME,
		"direct",
		false,
		false,
		false,
		false,
		nil,
	)
	LogSeriousFailed(err, "Can't declare exchange queue")


	rabbitEmailQueue, err = rabbitCh.QueueDeclare(
		"email",
		false,
		false,
		false,
		false,
		nil,
	)
	LogSeriousFailed(err, "Can't declare queue")

	err = rabbitCh.QueueBind(
		rabbitEmailQueue.Name,
		"email",
		EXCHANGE_NAME,
		false,
		nil,
	)
	LogSeriousFailed(err, "Can't bing queue with exchange queue")


}

func SendMessage(msgType string, msgKey string, msgBody *response.GetOrderResponse){
	var err error
	body, err := json.Marshal(msgBody)
	if err != nil {
		log.Printf("Failed to serialize message body: %v", err)
		return
	}

	err = rabbitCh.Publish(
		EXCHANGE_NAME,
		msgKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil{
		LogNormalFailed(err, "Can't send message to queue")
	}else{
		log.Printf("Message sent: {\n Type: %s,\n Message: %s\n}", msgType, body)
	}
}

func CleanupRabbitMQ() {
	if rabbitCh != nil {
		rabbitCh.Close()
	}
	if rabbitConn != nil {
		rabbitConn.Close()
	}
	log.Println("RabbitMQ connection and channel closed")
}


func TestEmailQueueConnection(){
	msgKey := "email"
	msgType := "TEST"
	msgBody := "TEST"

	err := rabbitCh.Publish(
		EXCHANGE_NAME,
		msgKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("[%s] %s", msgType, msgBody)),
		},
	)
	if err != nil{
		LogSeriousFailed(err, "Can't send message to queue")
	}else{
		log.Printf("Succesful connect to email queue")
	}
}