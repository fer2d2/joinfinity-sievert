package cmd

import (
	"github.com/fer2d2/sievert/logger"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use: Run,
	Short: "Run the sievert generated project. If Docker images are not " +
		"available, sievert will pull them before run (same as 'sievert build')",
	Run: func(cmd *cobra.Command, args []string) {
		logger.CommandCall(cmd.Use, args)
	},
}
