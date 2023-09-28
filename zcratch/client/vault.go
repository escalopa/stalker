package client

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

func NewVault(url string) (*vault.Client, error) {
	client, err := vault.New(
		vault.WithAddress(url),
		vault.WithRequestTimeout(DefaultTimeout),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to vault: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	resp, err := client.Auth.AppRoleLogin(
		ctx,
		schema.AppRoleLoginRequest{
			RoleId:   os.Getenv("MY_APPROLE_ROLE_ID"),
			SecretId: os.Getenv("MY_APPROLE_SECRET_ID"),
		},
		vault.WithMountPath(os.Getenv("MY_APPROLE_PATH")),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to login: %v", err)
	}

	// authenticate with a root token (insecure)
	if err := client.SetToken(resp.Auth.ClientToken); err != nil {
		return nil, fmt.Errorf("failed to set token: %v", err)
	}

	return client, nil
}
