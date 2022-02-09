/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package container

import (
	"github.com/spf13/cobra"
)

// containerCmd represents the container command
var ContainerCmd = &cobra.Command{
	Use:   "container",
	Short: "Manage containers",
	Long:  `Manage containers`,
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {

	ContainerCmd.AddCommand(ExecCmd)
	ContainerCmd.AddCommand(LogsCmd)
	ContainerCmd.AddCommand(RestartCmd)
	ContainerCmd.AddCommand(RmCmd)
	ContainerCmd.AddCommand(RunCmd)
	ContainerCmd.AddCommand(StartCmd)
	ContainerCmd.AddCommand(StopCmd)
}
