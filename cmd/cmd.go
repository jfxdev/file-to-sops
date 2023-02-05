package cmd

import (
	"github.com/jfxdev/file-to-sops/internal/converter"

	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "fts",
	Short: "File to SOPs",
	RunE:  run,
}

var flags Flags

type Flags struct {
	FilePath string
}

func Run() {
	cmd.Flags().StringVarP(&flags.FilePath, "file", "f", "", "required file path")
	cmd.Execute()
}

func run(cmd *cobra.Command, args []string) (err error) {
	err = converter.Parse(flags.FilePath)
	return
}
