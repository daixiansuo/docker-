/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package command

import (
	"docker-/internal/image"
	. "docker-/internal/log"
	"github.com/spf13/cobra"
)

// imagesCmd represents the images command
var ImagesCmd = &cobra.Command{
	Use:   "images",
	Short: "List images",
	Long:  `List images`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := image.ListHandler(); err != nil {
			Log.Errorf("List image error %v", err)
		}
	},
}
