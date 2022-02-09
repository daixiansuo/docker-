/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package image

import (
	"docker-/internal/image"
	. "docker-/internal/log"
	"fmt"
	"github.com/spf13/cobra"
)

// pullCmd represents the pull command
var PullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull an image or a repository from a registry",
	Long:  `Pull an image or a repository from a registry`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("docker- pull requires exactly 1 argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, imageName := range args {
			Log.Debugf("download image: %s starting ...", imageName)
			if err := image.PullHandler(imageName); err != nil {
				Log.Errorf("download image: %s error: %v", imageName, err)
			}
		}
	},
}
