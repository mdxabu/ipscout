package cmd

import (
	"fmt"
	"runtime"
	"github.com/spf13/cobra"
)

var (
	Version        = "1.0.0"
	GitCommit      = ""
	BuildTime      = ""
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Version of ipscout",
	Long:    `Get ipscout version, git commit, go version, build time, release channel, etc.`,
	Example: `  ipscout version`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version:        %s\n", Version)
		fmt.Printf("Go version:     %s\n", runtime.Version())
		fmt.Printf("Git commit:     %s\n", GitCommit)
		fmt.Printf("Built:          %s\n", BuildTime)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
