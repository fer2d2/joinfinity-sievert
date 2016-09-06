package cmd

import (
	"github.com/fer2d2/sievert/logger"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   Delete,
	Short: "@TODO",
	Run: func(cmd *cobra.Command, args []string) {
		logger.CommandCall(cmd.Use, args)
	},
}
