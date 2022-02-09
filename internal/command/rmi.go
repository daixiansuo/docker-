/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package command

import (
	"docker-/internal/image"
	. "docker-/internal/log"
	"github.com/spf13/cobra"
)

// rmiCmd represents the rmi command
var RmiCmd = &cobra.Command{
	Use:   "rmi",
	Short: "Remove one or more images",
	Long:  `Remove one or more images`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, imageName := range args {
			if err := image.RemoveHandler(imageName); err != nil {
				Log.Errorf("remove image %s error: %v", imageName, err)
			}
		}
	},
}
