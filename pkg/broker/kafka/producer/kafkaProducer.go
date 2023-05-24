package producer

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/tumbleweedd/mediasoft-intership/customer-service/internal/model"
)

type Message struct {
	Order model.OrdersByOffice `json:"Order"`
}

type Producer struct {
	p sarama.AsyncProducer
}

func NewProducer(broker string) (*Producer, error) {
	producer, err := sarama.NewAsyncProducer([]string{broker}, sarama.NewConfig())
	if err != nil {
		return nil, err
	}
	return &Producer{
		p: producer,
	}, nil
}

func (p *Producer) StartProduce(done chan struct{}, topic string, order model.OrdersByOffice) {
	msg := Message{
		Order: order,
	}

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		fmt.Printf("Failed to marshal message, err: %s\n", err)
		return
	}

	select {
	case <-done:
		return
	case p.p.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(msgBytes),
	}:
		fmt.Println("Order sent to Kafka")
	case err := <-p.p.Errors():
		fmt.Printf("Failed to send message to Kafka, err: %s, msg: %s\n", err, msgBytes)
	}
}

func (p *Producer) Close() error {
	if p != nil {
		return p.p.Close()
	}
	return nil
}
