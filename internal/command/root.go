package command

import "github.com/spf13/cobra"

func NewRoot() *cobra.Command {
	root := &cobra.Command{
		Use:   "fed",
		Short: "Fed prints money",
		Long:  "Fed generates invoices so you can make that boom dollar",
	}

	root.AddCommand(
		NewYAML(),
		NewPDF(),
		NewNew(),
		NewVersion(),
	)

	return root
}
