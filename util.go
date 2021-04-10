package lpc

import (
	"context"
	"os"
	"path/filepath"

	"github.com/ansd/lastpass-go"
)

func GetSecrets(ctx context.Context, username, password string, trusted bool) ([]*lastpass.Account, error) {
	var client *lastpass.Client
	var err error

	if !trusted {
		client, err = lastpass.NewClient(ctx, username, password)
	} else {
		homeDir, homeDirErr := os.UserHomeDir()
		if homeDirErr != nil {
			return nil, homeDirErr
		}
		configDir := filepath.Join(homeDir, ".lpc")
		client, err = lastpass.NewClient(
			ctx,
			username,
			password,
			lastpass.WithConfigDir(configDir),
			lastpass.WithTrust())
	}

	if err != nil {
		return nil, err
	}
	return client.Accounts(ctx)
}
