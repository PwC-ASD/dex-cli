package cmd

import (
	"fmt"
	"runtime"
	"github.com/PwC-ASD/dex-cli/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Print the version",
	Example: `  $ dex-cli version `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(
			"dex-cli Version: %s\nGo Version: %s\nGo OS/ARCH: %s %s\n",
			version.Version, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
