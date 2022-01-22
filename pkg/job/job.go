package job

import (
	"context"

	"github.com/Shopify/sarama"
)

type SaramaConsumerJobServer struct {
	topics []string
	kc     sarama.ConsumerGroupHandler
	kcg    sarama.ConsumerGroup
	ready  chan bool
}

func NewSaramaConsumerJobServer(topics []string,
	kc sarama.ConsumerGroupHandler, kcg sarama.ConsumerGroup) *SaramaConsumerJobServer {
	return &SaramaConsumerJobServer{topics: topics, kc: kc, kcg: kcg}
}

func (s *SaramaConsumerJobServer) Start(ctx context.Context) error {
	for {
		err := s.kcg.Consume(ctx, s.topics, s.kc)
		if err != nil {
			return err
		}
		if err := ctx.Err(); err != nil {
			return err
		}
		s.ready = make(chan bool)
	}
}

func (s *SaramaConsumerJobServer) Stop(ctx context.Context) error {
	ctx.Done()
	return nil
}
