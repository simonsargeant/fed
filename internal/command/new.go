package command

import (
	"time"

	"github.com/simonsargeant/fed/internal/load"
	"github.com/simonsargeant/fed/internal/template"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

func NewNew() *cobra.Command {
	var (
		date       string
		fontPath   string
		configPath string
		indexPath  string
		debug      bool
	)

	cmd := &cobra.Command{
		Use:   "new",
		Short: "Creates a new invoice and PDF, and adds them to a folder",
		Long:  "Creates a new invoice and PDF, creates directories to store them and moves them there",
		Run: func(cmd *cobra.Command, args []string) {
			fs := afero.Afero{Fs: afero.NewOsFs()}

			n := load.Index(fs, indexPath)

			var err error
			t := time.Now().Local()
			if date != "" {
				t, err = time.Parse("2/1/2006", date)
				if err != nil {
					log.Fatalf("Error parsing date %q: %s", date, err)
				}
			}

			filename, invoice, err := load.
				Config(fs, configPath).
				Create(n, t, args[0], load.Items(args[1:]...))

			if err != nil {
				log.Fatalf("Error creating invoice config: %s", err)
			}

			ymlOut, pdfOut := load.Dir(fs, t, filename)

			data, err := yaml.Marshal(&template.InvoiceContainer{Invoice: invoice})

			if err != nil {
				log.Fatalf("Error marshalling yaml: %s", err)
			}

			if _, err := ymlOut.Write(data); err != nil {
				log.Fatalf("Error writing yaml: %s", err)
			}

			d := template.NewDrawer(debug, log.StandardLogger())
			if err := d.Draw(invoice, fontPath, pdfOut); err != nil {
				log.Fatalf("Error drawing PDF: %s", err)
			}

			load.UpdateIndex(fs, indexPath, n)
		},
	}

	cmd.Flags().StringVarP(&date, "date", "t", "", "issued at date, default is today, format <day-int>/<month-int>/<year-int>")
	cmd.Flags().StringVarP(&fontPath, "font", "f", "/usr/share/fonts/noto/NotoSans-Regular.ttf", "font .ttf file to use")
	cmd.Flags().StringVarP(&configPath, "config", "c", "config.yml", "config file to use")
	cmd.Flags().StringVarP(&indexPath, "index", "n", "index.txt", "index file to use")
	cmd.Flags().BoolVarP(&debug, "debug", "d", false, "enable debug mode")

	return cmd
}
