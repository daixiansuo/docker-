/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rmiCmd represents the rmi command
var rmiCmd = &cobra.Command{
	Use:   "rmi",
	Short: "Remove one or more images",
	Long:  `Remove one or more images`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rmi called")
	},
}

func init() {
	rootCmd.AddCommand(rmiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rmiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rmiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
