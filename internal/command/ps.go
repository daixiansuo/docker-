/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

// psCmd represents the ps command
var PsCmd = &cobra.Command{
	Use:   "ps",
	Short: "List containers",
	Long:  `List containers`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ps called")
	},
}
