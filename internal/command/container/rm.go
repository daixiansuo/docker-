/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package container

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rmCmd represents the rm command
var RmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove one or more containers",
	Long:  `Remove one or more containers`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rm called")
	},
}
