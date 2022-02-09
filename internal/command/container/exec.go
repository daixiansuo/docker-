/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package container

import (
	"fmt"

	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var ExecCmd = &cobra.Command{
	Use:   "exec",
	Short: "Run a command in a running container",
	Long:  `Run a command in a running container`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("exec called")
	},
}
