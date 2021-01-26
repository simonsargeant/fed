package load

import (
	"io"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

func Write(res interface{}, fs afero.Afero, outputPath string, defaultOut io.Writer) {
	data, err := yaml.Marshal(res)

	if err != nil {
		log.Fatalf("Error marshalling yaml: %s", err)
	}

	out := Output(fs, outputPath, defaultOut)

	if _, err = out.Write(data); err != nil {
		log.Fatalf("Error writing yaml: %s", err)
	}
}
