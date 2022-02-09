/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package image

import (
	"github.com/spf13/cobra"
)

// imageCmd represents the image command
var ImageCmd = &cobra.Command{
	Use:   "image",
	Short: "Manage images",
	Long:  `Manage images`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {

	ImageCmd.AddCommand(lsCmd)
	ImageCmd.AddCommand(rmCmd)
	ImageCmd.AddCommand(PullCmd)
}
