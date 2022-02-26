/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package container

import (
	"docker-/internal/log"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Call yourself, Do not call it outside.",
	Long:  `Call yourself, Do not call it outside.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Log.Debugf("")
	},
}

func init() {
	// hidden command
	InitCmd.Hidden = true
}
