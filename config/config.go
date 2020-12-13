package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

var (
	services ServiceCfg
)

type ServiceCfg struct {
	ProductSrvAddr string `envconfig:"PRODUCT_URL" default:"http://localhost:3000"`
}

func NewConfig() {
	configs := map[string]interface{}{
		"service": &services,
	}

	for prefix, instance := range configs {
		err := envconfig.Process(prefix, instance)
		if err != nil {
			log.Fatalf("unable to init %s config", prefix)
		}
	}
}

func ServiceConfig() ServiceCfg {
	return services
}
