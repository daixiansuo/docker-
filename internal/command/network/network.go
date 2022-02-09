/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package network

import (
	"github.com/spf13/cobra"
)

// networkCmd represents the network command
var NetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "Manage networks",
	Long:  `Manage networks`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {

	NetworkCmd.AddCommand(createCmd)
	NetworkCmd.AddCommand(lsCmd)
	NetworkCmd.AddCommand(rmCmd)
}
