/*
Copyright © 2022 daixiansuo

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package command

import (
	"docker-/internal/command/container"
	"docker-/internal/command/image"
	"docker-/internal/command/network"
	. "docker-/internal/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "docker-",
	Short: "Docker- is a simple version implementation that imitates docker。",
	Long:  `Docker- is a simple version implementation that imitates docker。`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// print help info
		_ = cmd.Help()
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// global setting logs level
		debug, _ := strconv.ParseBool(cmd.Flag("debug").Value.String())
		if debug {
			Log.SetLevel(logrus.DebugLevel)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	// container commands
	rootCmd.AddCommand(container.ContainerCmd)
	rootCmd.AddCommand(container.RunCmd)
	rootCmd.AddCommand(container.InitCmd)
	rootCmd.AddCommand(container.StartCmd)
	rootCmd.AddCommand(container.StopCmd)
	rootCmd.AddCommand(container.RestartCmd)
	rootCmd.AddCommand(container.ExecCmd)
	rootCmd.AddCommand(container.LogsCmd)
	rootCmd.AddCommand(container.RmCmd)

	// image commands
	rootCmd.AddCommand(image.ImageCmd)
	rootCmd.AddCommand(image.PullCmd)

	// network commands
	rootCmd.AddCommand(network.NetworkCmd)

	// common commands
	rootCmd.AddCommand(ImagesCmd)
	rootCmd.AddCommand(RmiCmd)
	rootCmd.AddCommand(PsCmd)
	rootCmd.AddCommand(inspectCmd)
	rootCmd.AddCommand(versionCmd)

	// global flag setting
	rootCmd.PersistentFlags().BoolP("debug", "", false, "debug logs")
}
