package command

import (
	"time"

	"github.com/simonsargeant/fed/internal/load"
	"github.com/simonsargeant/fed/internal/template"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func NewYAML() *cobra.Command {
	var (
		date       string
		configPath string
		indexPath  string
		outputPath string
		number     int
		debug      bool
	)

	cmd := &cobra.Command{
		Use:   "yaml",
		Short: "Creates an invoice schema",
		Long: `Creates an invoice from the parameters and customer config to an invoice yml config

Provide items in a comma separated list of arguments`,
		Run: func(cmd *cobra.Command, args []string) {
			fs := afero.Afero{Fs: afero.NewOsFs()}

			var n int
			if number == 0 {
				n = load.Index(fs, indexPath)
			} else {
				n = number
			}

			var err error
			t := time.Now().Local()
			if date != "" {
				t, err = time.Parse("1/2/2006", date)
				if err != nil {
					log.Fatalf("Error parsing date %q: %s", date, err)
				}
			}

			_, invoice, err := load.
				Config(fs, configPath).
				Create(n, t, args[0], load.Items(args[1:]...))

			if err != nil {
				log.Fatalf("Error creating invoice config: %s", err)
			}

			res := &template.InvoiceContainer{
				Invoice: invoice,
			}

			load.Write(res, fs, outputPath, cmd.OutOrStdout())

			if number == 0 {
				load.UpdateIndex(fs, indexPath, n)
			}
		},
	}

	cmd.Flags().StringVarP(&date, "date", "t", "", "issued at date, default is today, format <day-int>/<month-int>/<year-int>")
	cmd.Flags().StringVarP(&configPath, "config", "c", "config.yml", "config file to use")
	cmd.Flags().StringVarP(&outputPath, "output-path", "o", "", "filename to write output to")
	cmd.Flags().StringVarP(&indexPath, "index-path", "i", "index.txt", "index file to")
	cmd.Flags().IntVarP(&number, "number", "n", 0, "invoice number override, will not be persisted")
	cmd.Flags().BoolVarP(&debug, "debug", "d", false, "enable debug mode, draws debug lines on PDFs")

	return cmd
}
