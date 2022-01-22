package job

import (
	"context"
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
	"github.com/starryrbs/kfan/app/history/service/internal/biz"
)

type consumeHandler struct {
	ready chan bool
	hc    biz.HistoryJobRepo
}

func NewConsumeHandler(hc biz.HistoryJobRepo) sarama.ConsumerGroupHandler {
	return &consumeHandler{ready: make(chan bool), hc: hc}
}

func (c *consumeHandler) Setup(session sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

func (c *consumeHandler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

func (c *consumeHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
		ctx := context.Background()
		var history biz.History

		err := json.Unmarshal(message.Value, &history)
		if err != nil {
			return err
		}
		_, err = c.hc.PersistentSaveHistory(ctx, &history)
		if err != nil {
			return err
		}
		session.MarkMessage(message, "")
	}
	return nil
}
