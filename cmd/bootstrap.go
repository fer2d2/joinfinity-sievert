package cmd

import (
	"github.com/fer2d2/sievert/action"
	"github.com/fer2d2/sievert/util"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(bootstrapCmd)
}

var bootstrapCmd = &cobra.Command{
	Use:   Bootstrap,
	Short: "Bootstrap the files in the ~/sievert/ directory",
	Run: func(cmd *cobra.Command, args []string) {
		util.LogRunningCommand(cmd.Use, args)
		action.BootstrapAction()
	},
}
