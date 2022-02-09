/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package container

import (
	"fmt"

	"github.com/spf13/cobra"
)

// logsCmd represents the logs command
var LogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Fetch the logs of a container",
	Long:  `Fetch the logs of a container`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("logs called")
	},
}
