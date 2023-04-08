package config_center

import (
	"context"
	"github.com/dk-sirius/scaf-fold-demo/api"
	"github.com/dk-sirius/scaf-fold-demo/server/db/configcenter/gen/query"
	"github.com/dk-sirius/scaf-fold-demo/server/global"
)

type ConfigCenter struct {
	api.UnimplementedConfigServiceServer
}

func (receiver *ConfigCenter) ListHarbors(ctx context.Context, in *api.HarborRequest) (*api.HarborResponse, error) {
	harbor := query.Use(global.ConfigDb).Harbour
	data, err := harbor.WithContext(context.Background()).Find()
	if err != nil {
		return nil, err
	}
	hs := make([]*api.Harbor, 0)
	for _, v := range data {
		sub := &api.Harbor{}
		sub.Id = v.FID
		sub.Name = v.FNameZh
		sub.EnName = v.FNameEn
		hs = append(hs, sub)
	}
	resp := &api.HarborResponse{}
	resp.Data = hs
	return resp, nil
}
