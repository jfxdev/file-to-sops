package cmd

import (
	"sops-for-files/app"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sopsff",
	Short: "SOPS for Files",
	RunE:  run,
}

var flags Flags

type Flags struct {
	FilePath string
}

func Run() {
	rootCmd.Flags().StringVarP(&flags.FilePath, "file", "f", "", "required file path")
	rootCmd.Execute()
}

func run(cmd *cobra.Command, args []string) (err error) {
	err = app.Single(flags.FilePath)
	return
}
