package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(provisionCmd)
}

var provisionCmd = &cobra.Command{
	Use:   Provision,
	Short: "Provision the sievert environment",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd)
		fmt.Println("provision called")
	},
}
