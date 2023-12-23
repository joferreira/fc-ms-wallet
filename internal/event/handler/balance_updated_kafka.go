package handler

import (
	"fmt"
	"sync"

	"github.com/joferreira/fc-ms-wallet/pkg/events"
	"github.com/joferreira/fc-ms-wallet/pkg/kafka"
)

type UpdateBalancedKafkaHandler struct {
	Kafka *kafka.Producer
}

func NewUpdateBalancedKafkaHandler(kafka *kafka.Producer) *UpdateBalancedKafkaHandler {
	return &UpdateBalancedKafkaHandler{
		Kafka: kafka,
	}
}

func (h *UpdateBalancedKafkaHandler) Handle(message events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	h.Kafka.Publish(message, nil, "balances")
	fmt.Println("UpdateBalanceKafkaHandler called")
}
