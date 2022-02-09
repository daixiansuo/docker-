/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package container

import (
	"fmt"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start one or more stopped containers",
	Long:  `Start one or more stopped containers`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
	},
}
