package cmd

import (
	"github.com/fer2d2/sievert/util"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(provisionCmd)
}

var provisionCmd = &cobra.Command{
	Use:   Provision,
	Short: "Provision the sievert environment",
	Run: func(cmd *cobra.Command, args []string) {
		util.LogRunningCommand(cmd.Use, args)
	},
}
