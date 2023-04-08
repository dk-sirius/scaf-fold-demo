package business

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/dk-sirius/scaf-fold-demo/api"
)

func init() {
	impl = new(api.ConfigServiceClientImpl)
	config.SetConsumerService(impl)
	if err := config.Load(); err != nil {
		panic(err)
	}
}

var impl *api.ConfigServiceClientImpl

type ConfigClient struct {
	impl *api.ConfigServiceClientImpl
}

func NewConfigClient() *ConfigClient {
	return &ConfigClient{impl: impl}
}
func (receiver *ConfigClient) ListHarbors() (*api.HarborResponse, error) {
	req := &api.HarborRequest{}
	return receiver.impl.ListHarbors(context.Background(), req)
}
