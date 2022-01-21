package job

import (
	"github.com/Shopify/sarama"
	"github.com/starryrbs/kfan/app/history/service/internal/biz"
	"github.com/starryrbs/kfan/app/history/service/internal/conf"
	"github.com/starryrbs/kfan/pkg/job"
)

func NewJobServer(hc biz.HistoryJobRepo, c *conf.Data, kcg sarama.ConsumerGroup) *job.SaramaConsumerJobServer {
	kc := NewConsumeHandler(hc)
	return job.NewSaramaConsumerJobServer(
		[]string{c.Kafka.Topic},
		kc,
		kcg,
	)
}
