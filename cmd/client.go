package cmd

import (
	"errors"
	"fmt"
	client "github.com/PwC-ASD/dex-cli/api"
	"github.com/coreos/dex/api"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:     "client",
	Short:   "Manage a client on Dex",
	Example: `  $ dex-cli client create --name <a-new-app>`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var clientCreateCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a new Client on Dex",
	Example: `  $ dex-cli client create --name <a-new-app>`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		if name == "" {
			return errors.New("client name missing")
		}

		var kubeClient *client.KubeClient
		if k8sSecret != "" {
			kubeClient, err = client.NewKubeClient(kubeConfig)
			if err != nil {
				return err
			}
			exist, err := kubeClient.SecretExist(k8sSecret)
			if err != nil {
				return err
			}
			if exist {
				return errors.New("k8s Secret " + k8sSecret + " already exists.")
			}
		}

		var redirectUris []string
		if redirectURI != "" {
			redirectUris = []string{redirectURI}
		} else {
			redirectUris = []string{}
		}

		newClient := api.Client{
			Name:         name,
			Id:           id,
			Secret:       secret,
			RedirectUris: redirectUris,
		}

		newClientReq := api.CreateClientReq{
			Client: &newClient,
		}

		c, err := client.NewDexClient(server)
		if err != nil {
			return err
		}
		resp, err := c.ClientCreate(&newClientReq)
		if err != nil {
			return err
		}
		fmt.Printf(
			"client successfully created."+
				" \nname: %s\nid: %s\nsecret: %s\nredirect-uri: %s\n",
			resp.Client.Name, resp.Client.Id, resp.Client.Secret, resp.Client.RedirectUris)

		if k8sSecret != "" {
			_, err = kubeClient.SecretCreate(resp.Client.Id, resp.Client.Secret, k8sSecret)
			if err != nil {
				return err
			}
			fmt.Printf("client's ID and secret succesfully saved to %s k8s secret.\n", k8sSecret)
		}

		return nil
	},
}

var clientDeleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete an existing Dex client",
	Example: `  $ dex-cli client delete --id <an-app-id>`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if id == "" {
			return errors.New("client Id missing")
		}

		deleteClientReq := api.DeleteClientReq{
			Id: id,
		}

		dexClient, err := client.NewDexClient(server)
		if err != nil {
			return err
		}
		_, err = dexClient.ClientDelete(&deleteClientReq)
		if err == nil {
			fmt.Printf(
				"client successfully deleted.\n")
		}

		return err
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)

	clientCmd.AddCommand(clientCreateCmd)
	clientCreateCmd.Flags().StringVar(&name, "name", "",
		"\n\tName of the client to create.")
	clientCreateCmd.Flags().StringVar(&id, "id", "",
		"\n\tOptional: ID of the client to create.")
	clientCreateCmd.Flags().StringVar(&secret, "secret", "",
		"\n\tOptional: Secret of the client to create.")
	clientCreateCmd.Flags().StringVar(&redirectURI, "redirect-uri", "",
		"\n\tOptional: A redirect URI to register while creating the client.")
	clientCreateCmd.Flags().StringVar(&k8sSecret, "k8s-secret", "",
		"\n\tOptional: k8s secret name (NAMESPACE/SECRET_NAME) to store client's Id and secret.")

	clientCmd.AddCommand(clientDeleteCmd)
	clientDeleteCmd.Flags().StringVar(&id, "id", "",
		"\n\tID of the client to delete.")
}
