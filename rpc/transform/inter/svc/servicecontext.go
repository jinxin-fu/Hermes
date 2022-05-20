package svc

import (
	"github.com/Hermes/rpc/transform/inter/config"
	"github.com/Hermes/rpc/transform/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  model.HermesdModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewHermesdModel(sqlx.NewMysql(c.DataSource), c.Cache),
	}
}
