/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the Docker version information",
	Long:  `Show the Docker version information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version 1.0.0")
	},
}
