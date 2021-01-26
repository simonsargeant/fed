package load

import (
	"fmt"
	"io"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

func Dir(fs afero.Fs, date time.Time, name string) (io.Writer, io.Writer) {
	y, m, _ := date.Date()
	dir := fmt.Sprintf("%d/Q%d", y, 1+(m-1)/3)

	if err := fs.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("Error creating directory %q: %s", dir, err)
	}

	ymlPath := dir + "/" + name + ".yml"
	pdfPath := dir + "/" + name + ".pdf"

	var ymlOut io.Writer
	_, err := fs.Stat(ymlPath)
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error checking yml file status: %s", err)
	}

	if os.IsNotExist(err) {
		ymlOut, err = fs.Create(ymlPath)
		if err != nil {
			log.Fatalf("Error creating yml file: %s", err)
		}

	} else {
		ymlOut, err = fs.OpenFile(ymlPath, os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Error opening yml file: %s", err)
		}
	}

	var pdfOut io.Writer
	_, err = fs.Stat(pdfPath)
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error checking yml file status: %s", err)
	}

	if os.IsNotExist(err) {
		pdfOut, err = fs.Create(pdfPath)
		if err != nil {
			log.Fatalf("Error creating yml file: %s", err)
		}

	} else {
		pdfOut, err = fs.OpenFile(pdfPath, os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Error opening yml file: %s", err)
		}
	}

	return ymlOut, pdfOut
}
