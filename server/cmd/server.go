package main

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/dk-sirius/scaf-fold-demo/server/global"
	"github.com/dk-sirius/scaf-fold-demo/server/pkg/config_center"
	"github.com/scaf-fold/tools/pkg/gormer"
)

func init() {
	global.Conf(func() *gormer.Model {
		conf := "../conf/models.yaml"
		assemble := gormer.NewModelConverter(conf)
		m, err := assemble.Converter()
		if err != nil {
			panic(err)
		}
		return m
	})
}

func main() {
	config.SetProviderService(&config_center.ConfigCenter{})
	if err := config.Load(); err != nil {
		panic(err)
	}
	select {}
}
