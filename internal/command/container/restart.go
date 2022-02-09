/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package container

import (
	"fmt"

	"github.com/spf13/cobra"
)

// restartCmd represents the restart command
var RestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart one or more containers",
	Long:  `Restart one or more containers`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("restart called")
	},
}
