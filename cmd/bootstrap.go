package cmd

import (
	"github.com/fer2d2/sievert/logger"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(bootstrapCmd)
}

var bootstrapCmd = &cobra.Command{
	Use:   Bootstrap,
	Short: "Generate the docker project files from the sievert.yml configuration",
	Run: func(cmd *cobra.Command, args []string) {
		logger.CommandCall(cmd.Use, args)

		// chain := new(action.Chain)
		//
		// chain.
		// 	Add(action.GenerateNginxServers).
		// 	Execute()
	},
}
