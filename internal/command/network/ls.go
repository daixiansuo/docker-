/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package network

import (
	. "docker-/internal/log"
	"docker-/internal/network"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List networks",
	Long:  `List networks`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := network.List(); err != nil {
			Log.Errorf("List networks error %v", err)
		}
	},
}
