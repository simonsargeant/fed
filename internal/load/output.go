package load

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

func Output(fs afero.Afero, outputPath string, defaultOut io.Writer) io.Writer {
	if outputPath == "" {
		return defaultOut
	}

	_, err := fs.Stat(outputPath)
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error checking output file status: %s", err)
	}

	if os.IsNotExist(err) {
		out, err := fs.Create(outputPath)
		if err != nil {
			log.Fatalf("Error creating output file: %s", err)
		}

		return out
	}

	f, err := fs.OpenFile(outputPath, os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening output file: %s", err)
	}

	return f
}
