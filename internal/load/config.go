package load

import (
	"github.com/simonsargeant/fed/internal/config"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

func Config(fs afero.Afero, configPath string) config.Data {
	data, err := fs.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error reading config file from %q: %s", configPath, err)
	}

	conf := config.Data{}
	if err := yaml.Unmarshal(data, &conf); err != nil {
		log.Fatalf("Error unmarshalling conf: %s", err)
	}

	return conf
}
