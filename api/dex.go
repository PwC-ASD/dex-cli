package api

import (
	"context"
	"errors"
	"github.com/coreos/dex/api"
	"google.golang.org/grpc"
)

type DexClient struct {
	conn *grpc.ClientConn
}

func NewDexClient(host string) (*DexClient, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, errors.New("dial to gRPC endpoint failed")
	}
	c := &DexClient{
		conn: conn,
	}

	return c, err
}

func (c *DexClient) ClientCreate(request *api.CreateClientReq) (*api.CreateClientResp, error) {
	conn := c.conn
	defer conn.Close()

	dexClient := api.NewDexClient(conn)
	resp, err := dexClient.CreateClient(context.TODO(), request)
	if err != nil {
		return nil, err
	}
	if resp.AlreadyExists {
		return nil, errors.New("client already exists")
	}

	return resp, err
}

func (c *DexClient) ClientAddRedirectUri(request *api.AddClientRedirectUriReq) (*api.AddClientRedirectUriResp, error) {
	conn := c.conn
	defer conn.Close()

	dexClient := api.NewDexClient(conn)
	resp, err := dexClient.AddClientRedirectUri(context.TODO(), request)
	if err != nil {
		return nil, err
	}
	if resp.ClientNotFound {
		return nil, errors.New("client not found")
	}
	if resp.RedirectUriAlreadyExists {
		return nil, errors.New("redirect uri already exists")
	}

	return resp, err
}

func (c *DexClient) ClientRemoveRedirectUri(request *api.RemoveClientRedirectUriReq) (*api.RemoveClientRedirectUriResp, error) {
	conn := c.conn
	defer conn.Close()

	dexClient := api.NewDexClient(conn)
	resp, err := dexClient.RemoveClientRedirectUri(context.TODO(), request)
	if err != nil {
		return nil, err
	}
	if resp.ClientNotFound {
		return nil, errors.New("client not found")
	}
	if resp.RedirectUriNotFound {
		return nil, errors.New("redirect uri not found")
	}

	return resp, err
}

func (c *DexClient) ClientDelete(request *api.DeleteClientReq) (*api.DeleteClientResp, error) {
	conn := c.conn
	defer conn.Close()

	dexClient := api.NewDexClient(conn)
	resp, err := dexClient.DeleteClient(context.TODO(), request)
	if err != nil {
		return nil, err
	}
	if resp.NotFound {
		return nil, errors.New("client not found")
	}

	return resp, err
}
