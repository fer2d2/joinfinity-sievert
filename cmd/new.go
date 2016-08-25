package cmd

import (
	"github.com/fer2d2/sievert/action"
	"github.com/fer2d2/sievert/util"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:   New,
	Short: "Install Sievert for first time",
	Run: func(cmd *cobra.Command, args []string) {
		util.LogRunningCommand(cmd.Use, args)

		chain := new(action.Chain)

		chain.
			Add(action.RequireConfigDirNotCreated).
			Add(action.CreateConfigDir).
			Add(action.CreateSievertYmlFile).
			Execute()
	},
}
