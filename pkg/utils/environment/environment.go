package environment

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type Environment struct {
	ConfigFilePath string `env:"CONFIG_CONNECTORCONFIG_FILE_PATH"`
}

var Settings Environment

func init() {
	if Settings == (Environment{}) {
		if err := env.Parse(&Settings); err != nil {
			log.Fatal(err)
		}
	}
}


