/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package network

import (
	. "docker-/internal/log"
	"docker-/internal/network"
	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove one or more networks",
	Long:  `Remove one or more networks`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := network.Remove(args[0]); err != nil {
			Log.Errorf("Remove netowrk %s error %v", args[0], err)
		}
	},
}
