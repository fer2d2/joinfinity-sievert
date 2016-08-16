package cmd

import "github.com/spf13/cobra"

// SievertCmd is the root command, $ sievert <subcommand> <flags>
var RootCmd = &cobra.Command{
	Use:   "sievert",
	Short: "Sievert is a tool for easily generate docker-compose environments for PHP developers",
}
