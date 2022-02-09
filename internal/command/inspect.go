/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

// inspectCmd represents the inspect command
var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Return low-level information on Docker- objects",
	Long:  `Return low-level information on Docker- objects`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("inspect called")
	},
}
