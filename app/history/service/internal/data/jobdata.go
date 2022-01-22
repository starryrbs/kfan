package data

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/starryrbs/kfan/app/history/service/internal/conf"
	ent "github.com/starryrbs/kfan/app/history/service/internal/data/ent"
)

var JobProviderSet = wire.NewSet(NewJobData, NewKafkaConsumeGroup, NewHistoryJobRepo)

type JobData struct {
	db *ent.Client
}

func NewJobData(c *conf.Data, logger log.Logger) (*JobData, func(), error) {
	log := log.NewHelper(logger)
	drv, err := sql.Open(
		c.Database.Driver,
		c.Database.Source,
	)

	client := ent.NewClient(ent.Driver(drv))
	if err != nil {
		return nil, nil, err
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Errorf("failed creating schema resources: %v", err)
		return nil, nil, err
	}

	return &JobData{
			db: client,
		}, func() {
			if err := client.Close(); err != nil {
				log.Error(err)
			}
		}, nil
}

func NewKafkaConsumeGroup(c *conf.Data) (sarama.ConsumerGroup, error) {
	return sarama.NewConsumerGroup(c.Kafka.Addrs, c.Kafka.GroupId, sarama.NewConfig())
}
