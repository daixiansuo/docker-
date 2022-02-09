/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package image

import (
	"docker-/internal/image"
	. "docker-/internal/log"
	"github.com/spf13/cobra"
)

// lsCmd represents the ls command
var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List images",
	Long:  "List images",
	Run: func(cmd *cobra.Command, args []string) {
		if err := image.ListHandler(); err != nil {
			Log.Errorf("List image error %v", err)
		}
	},
}
