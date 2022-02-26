/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package container

import (
	"docker-/internal/container"
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var RunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a command in a new container",
	Long:  `Run a command in a new container`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := container.Run(); err != nil {
			return
		}
	},
}

func init() {

	RunCmd.Flags().BoolP("detach", "d", false, "Run container in background and print container ID")
	RunCmd.Flags().StringP("name", "", "", "Assign a name to the container")
	RunCmd.Flags().StringP("hostname", "", "", "Container host name")
	RunCmd.Flags().StringP("dns", "", "8.8.8.8,4.4.4.4", "Set custom DNS servers")
	RunCmd.Flags().StringP("network", "", "", "Connect a container to a network")

	RunCmd.Flags().StringP("env", "e", "", "Set environment variables")
	RunCmd.Flags().StringP("volume", "v", "", "Bind mount a volume")
	RunCmd.Flags().StringP("publish", "p", "", "Publish a container's port(s) to the host")
	RunCmd.Flags().StringP("storage-driver", "", "", "Storage driver options for the container")

	RunCmd.Flags().StringP("memory-limit", "", "", "memory limit")
	RunCmd.Flags().StringP("cpu-share", "", "", "cpu shares")
	RunCmd.Flags().StringP("cpu-set", "", "", "cpu set")
}
