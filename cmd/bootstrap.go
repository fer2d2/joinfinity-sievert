package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(bootstrapCmd)
}

var bootstrapCmd = &cobra.Command{
	Use:   Bootstrap,
	Short: "Bootstrap the files in the ~/sievert/ directory",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bootstrap called")
	},
}
