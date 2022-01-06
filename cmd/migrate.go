package cmd

import (
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate command",
	Long:  "migrate command to handel sql migration ",
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
