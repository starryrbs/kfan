package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	consulAPI "github.com/hashicorp/consul/api"
	housev1 "github.com/starryrbs/kfan/api/house/service/v1"
	"github.com/starryrbs/kfan/app/history/service/internal/conf"
	ent "github.com/starryrbs/kfan/app/history/service/internal/data/ent"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

var JobProviderSet = wire.NewSet(NewJobData, NewKafkaConsumeGroup, NewDiscovery, NewHistoryServiceClient, NewHistoryJobRepo)

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

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Consul.Address
	c.Scheme = conf.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	return r
}

func NewHistoryServiceClient(r registry.Discovery, tp *tracesdk.TracerProvider) housev1.HouseClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///kfan.house.service"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return housev1.NewHouseClient(conn)
}

func NewKafkaConsumeGroup(c *conf.Data) (sarama.ConsumerGroup, error) {
	return sarama.NewConsumerGroup(c.Kafka.Addrs, c.Kafka.GroupId, sarama.NewConfig())
}
