package cmd

import (
	"github.com/fer2d2/sievert/logger"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:   New,
	Short: "Generate a new sievert project",
	Run: func(cmd *cobra.Command, args []string) {
		logger.CommandCall(cmd.Use, args)

		//
		// chain := new(action.Chain)
		//
		// chain.
		// 	Add(action.RequireConfigDirNotCreated).
		// 	Add(action.CreateConfigDir).
		// 	Add(action.CreateSievertYmlFile).
		// 	Execute()
	},
}
