package command

import (
	"github.com/simonsargeant/fed/internal/load"
	"github.com/simonsargeant/fed/internal/template"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func NewPDF() *cobra.Command {
	var (
		fontPath   string
		inputPath  string
		outputPath string
		debug      bool
	)

	cmd := &cobra.Command{
		Use:   "pdf",
		Short: "Generates an invoice PDF",
		Long:  "Generates an invoice PDF from the invoice config passed to stdin",
		Run: func(cmd *cobra.Command, args []string) {
			d := template.NewDrawer(debug, log.StandardLogger())
			fs := afero.Afero{Fs: afero.NewOsFs()}
			invoice := &template.InvoiceContainer{}

			load.Input(invoice, fs, inputPath, cmd.InOrStdin())
			out := load.Output(fs, outputPath, cmd.OutOrStdout())

			if err := d.Draw(invoice.Invoice, fontPath, out); err != nil {
				log.Fatalf("Error drawing PDF: %s", err)
			}
		},
	}

	cmd.Flags().StringVarP(&fontPath, "font", "f", "/usr/share/fonts/noto/NotoSans-Regular.ttf", "font .ttf file to use")
	cmd.Flags().StringVarP(&inputPath, "input-path", "i", "", "filename to read input from")
	cmd.Flags().StringVarP(&outputPath, "output-path", "o", "", "filename to write output to")
	cmd.Flags().BoolVarP(&debug, "debug", "d", false, "enable debug mode")

	return cmd
}
