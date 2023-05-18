package rabbitmq

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/internal/model"
)

// RabbitMQConn тип для соединения с RabbitMQ
type RabbitMQConn struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	queueName string
}

// NewRabbitMQConn создание соединения с RabbitMQ
func NewRabbitMQConn(user, password, host, port string, queueName string) (*RabbitMQConn, error) {
	connAddr := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		user,
		password,
		host,
		port,
	)

	conn, err := amqp.Dial(connAddr)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}
	return &RabbitMQConn{conn, channel, queueName}, nil
}

// Close Закрытие соединения с RabbitMQ
func (r *RabbitMQConn) Close() error {
	if err := r.channel.Close(); err != nil {
		return err
	}

	if err := r.conn.Close(); err != nil {
		return err
	}

	return nil
}

// SendOrder отправка заказа в RabbitMQ
func (r *RabbitMQConn) SendOrder(order *model.Order) error {
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}

	return r.channel.Publish("", r.queueName, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
}
