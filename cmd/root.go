package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var (
	name			string
	id				string
	secret			string
	redirectURI		string
	k8sSecret 		string
	kubeConfig		string
	server			string
)

const (
	defaultCliServer = "localhost:5557"
)

var rootCmd = &cobra.Command{
	Use:     "dex-cli",
	Short:   "Command line program for Dex",
	Long:    "A command line program to interact with Dex API.",
	Example: `  $ dex-cli create client --name <a-new-app>`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// Execute launch dex-cli command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		//fmt.Println(err) // Should be use for logging
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&server, "server", "s", "",
		"\n\tDex server. Can also be set using the"+
			"\n\tenvironment variable DEX_CLI_SERVER (default value is "+defaultCliServer+")")
	rootCmd.PersistentFlags().StringVarP(&kubeConfig, "kube-config", "k", "",
		"\n\tPath to a kube config file, otherwise use an in-cluster configuration. Can also be"+
			"\n\tset using the environment variable DEX_CLI_KUBECONFIG")
	rootCmd.SilenceUsage = true
}

func initConfig() {
	if server == "" {
		server = os.Getenv("DEX_CLI_SERVER")
		if server == "" {
			server = defaultCliServer
		}
	}
	if kubeConfig == "" {
		kubeConfig = os.Getenv("DEX_CLI_KUBECONFIG")
	}
}
