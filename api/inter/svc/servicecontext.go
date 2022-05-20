package svc

import (
	"github.com/jinxin-fu/hermes/api/inter/config"
)

type ServiceContext struct {
	Config config.Config
	//Transformer transformer.Transformer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		//Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transform)),
	}
}
