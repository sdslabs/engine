package client // import "github.com/sdslabs/docker/client"

import (
	"context"
	"encoding/json"

	"github.com/sdslabs/docker/api/types"
)

// ServerVersion returns information of the docker client and server host.
func (cli *Client) ServerVersion(ctx context.Context) (types.Version, error) {
	resp, err := cli.get(ctx, "/version", nil, nil)
	if err != nil {
		return types.Version{}, err
	}

	var server types.Version
	err = json.NewDecoder(resp.body).Decode(&server)
	ensureReaderClosed(resp)
	return server, err
}
