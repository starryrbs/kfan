package data

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/starryrbs/kfan/app/house/internal/conf"
	"github.com/starryrbs/kfan/app/house/internal/data/ent"
)

// ProviderSet2 is data providers.
var ProviderSet2 = wire.NewSet(NewData, NewHouseRepo)

// Data .
type Data struct {
	db *ent.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
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
	d := &Data{
		db: client,
	}

	return d, func() {
		if err := d.db.Close(); err != nil {
			log.Error(err)
		}
	}, nil
}
