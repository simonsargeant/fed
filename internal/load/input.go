package load

import (
	"io"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

func Input(res interface{}, fs afero.Afero, inputPath string, defaultIn io.Reader) {
	var data []byte
	var err error
	if inputPath == "" {
		data, err = ioutil.ReadAll(defaultIn)
		if err != nil {
			log.Fatalf("Error reading input: %s", err)
		}
	} else if data, err = fs.ReadFile(inputPath); err != nil {
		log.Fatalf("Error reading input file: %s", err)
	}

	err = yaml.Unmarshal(data, res)

	if err != nil {
		log.Fatalf("Error unmarshalling yaml: %s", err)
	}
}
