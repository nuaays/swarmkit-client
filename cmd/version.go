package cmd

import (
	"fmt"

	"github.com/shenshouer/swarmkit-client/version"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version ",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("%s [%s] (%s)", version.VERSION, version.BUILDTIME, version.GITCOMMIT))
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
