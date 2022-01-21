package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/starryrbs/kfan/app/history/service/internal/conf"
	ent "github.com/starryrbs/kfan/app/history/service/internal/data/ent"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewHistoryRepo, NewKafkaProducer)

// Data .
type Data struct {
	// 依赖kafka，redis，mysql
	db  *ent.Client
	rdb *redis.Client
	kp  sarama.AsyncProducer
}

// NewData .
func NewData(kp sarama.AsyncProducer, c *conf.Data, logger log.Logger) (*Data, func(), error) {
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

	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		DB:           int(c.Redis.Db),
		DialTimeout:  c.Redis.DialTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
	})

	d := &Data{
		db:  client,
		kp:  kp,
		rdb: rdb,
	}

	return d, func() {

		if err := d.rdb.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}

func NewKafkaProducer(conf *conf.Data) sarama.AsyncProducer {
	c := sarama.NewConfig()
	p, err := sarama.NewAsyncProducer(conf.Kafka.Addrs, c)
	if err != nil {
		panic(err)
	}
	return p
}
