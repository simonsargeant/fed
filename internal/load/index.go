package load

import (
	"io"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
)

func Index(fs afero.Afero, indexPath string) int {
	_, err := fs.Stat(indexPath)

	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error checking index file status: %s", err)
	}

	if os.IsNotExist(err) {
		log.Warnf("No file found at %q, intialising with invoice number 1", indexPath)
		return 1
	}

	data, err := fs.ReadFile(indexPath)
	if err != nil {
		log.Fatalf("Error reading index file from %q: %s", indexPath, err)
	}

	n, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		log.Fatalf("Error converting index string %q to integer: %s", data, err)
	}

	return n + 1
}

func UpdateIndex(fs afero.Afero, indexPath string, n int) {
	_, err := fs.Stat(indexPath)
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error checking index file status: %s", err)
	}

	var out io.Writer
	if os.IsNotExist(err) {
		out, err = fs.Create(indexPath)
		if err != nil {
			log.Fatalf("Error creating index file: %s", err)
		}

	} else {
		out, err = fs.OpenFile(indexPath, os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Error opening index file: %s", err)
		}
	}

	if _, err = out.Write([]byte(strconv.Itoa(n) + "\n")); err != nil {
		log.Fatalf("Error writing index file: %s", err)
	}
}
