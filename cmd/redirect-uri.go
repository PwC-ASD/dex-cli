package cmd

import (
	"errors"
	"fmt"
	"github.com/coreos/dex/api"
	client "github.com/PwC-ASD/dex-cli/api"
	"github.com/spf13/cobra"
)

var redirectCmd = &cobra.Command{
	Use:     "redirect-uri",
	Short:   "Manage redirect URIs of a Dex client",
	Example: `  $ dex-cli redirect-uri add --client-id <a-client-id> --uri <http://host.name/a/callback/uri> `,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var redirectAddCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a redirect URI for an existing Dex client",
	Example: `  $ dex-cli redirect-uri add --client-id <a-client-id> --uri <http://host.name/a/callback/uri>`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if id == "" {
			return errors.New("client ID missing")
		}

		if redirectURI == "" {
			return errors.New("redirect URI missing")
		}

		addClientRedirectUriReq := &api.AddClientRedirectUriReq{
			Id: id,
			RedirectUri: redirectURI,
		}

		dexClient, err := client.NewDexClient(server)
		if err != nil {
			return err
		}
		_, err = dexClient.ClientAddRedirectUri(addClientRedirectUriReq)
		if err == nil {
			fmt.Printf(
				"redirect URI successfully added.\n")
		}

		return err
	},
}

var redirectRemoveCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove a redirect URI for an existing Dex client",
	Example: `  $ dex-cli redirect-uri remove --client-id <a-client-id> --uri <http://host.name/a/callback/uri>`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if id == "" {
			return errors.New("Client Id missing")
		}

		if redirectURI == "" {
			return errors.New("Redirect URI missing")
		}

		removeClientRedirectUriReq := &api.RemoveClientRedirectUriReq{
			Id: id,
			RedirectUri: redirectURI,
		}

		dexClient, err := client.NewDexClient(server)
		if err != nil {
			return err
		}
		_, err = dexClient.ClientRemoveRedirectUri(removeClientRedirectUriReq)
		if err == nil {
			fmt.Printf(
				"redirect URI successfully removed.\n")
		}

		return err
	},
}

func init() {
	rootCmd.AddCommand(redirectCmd)

	redirectCmd.AddCommand(redirectAddCmd)
	redirectAddCmd.Flags().StringVar(&id, "client-id", "",
		"\n\tId of the client for which the redirect URI will be added.")
	redirectAddCmd.Flags().StringVar(&redirectURI, "uri", "",
		"\n\tRedirect URI to add.")

	redirectCmd.AddCommand(redirectRemoveCmd)
	redirectRemoveCmd.Flags().StringVar(&id, "client-id", "",
		"\n\tId of the client for which the redirect URI will be removed.")
	redirectRemoveCmd.Flags().StringVar(&redirectURI, "uri", "",
		"\n\tRedirect URI to remove.")
}
