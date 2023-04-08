package global

import (
	"github.com/scaf-fold/tools/pkg/gormer"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	CommodityDb        *gorm.DB
	ShipDb             *gorm.DB
	SinotransbookingDb *gorm.DB
	ConfigDb           *gorm.DB
)

func Conf(conf func() *gormer.Model) {
	ms := conf()
	CommodityDb, _ = gorm.Open(postgres.New(postgres.Config{
		DSN: ms.Models[0].Conn,
	}))
	ShipDb, _ = gorm.Open(postgres.New(postgres.Config{
		DSN: ms.Models[1].Conn,
	}))
	SinotransbookingDb, _ = gorm.Open(postgres.New(postgres.Config{
		DSN: ms.Models[2].Conn,
	}))

	ConfigDb, _ = gorm.Open(postgres.New(postgres.Config{
		DSN: ms.Models[4].Conn,
	}))
}
