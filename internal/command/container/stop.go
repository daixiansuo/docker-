/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package container

import (
	"fmt"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop one or more running containers",
	Long:  `Stop one or more running containers`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stop called")
	},
}
