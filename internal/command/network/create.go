/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package network

import (
	. "docker-/internal/log"
	"docker-/internal/network"
	"fmt"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a network",
	Long:  `Create a network`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("docker- network create requires exactly 1 argument")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if model, err := parse(cmd, args[0]); err != nil {
			Log.Errorf("docker- create network parse flags %v error %v", cmd.Flags(), err)
		} else {
			if err := network.Create(model); err != nil {
				Log.Errorf("docker- create network %s error %v", args[0], err)
			}
		}
	},
}

func init() {
	createCmd.Flags().StringP("subnet", "s", "172.17.0.0/16", "Subnet in CIDR format that represents a network segment")
	createCmd.Flags().StringP("driver", "d", "bridge", "Driver to manage the Network")
}

func parse(cmd *cobra.Command, name string) (*network.CreateModel, error) {

	subnet := cmd.Flag("subnet").Value.String()
	driver := cmd.Flag("driver").Value.String()
	model := &network.CreateModel{
		Name:   name,
		Subnet: subnet,
		Driver: driver,
	}

	return model, nil
}
