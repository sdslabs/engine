package client // import "github.com/sdslabs/docker/client"

import (
	"context"

	"github.com/sdslabs/docker/api/types"
)

type apiClientExperimental interface {
	CheckpointAPIClient
}

// CheckpointAPIClient defines API client methods for the checkpoints
type CheckpointAPIClient interface {
	CheckpointCreate(ctx context.Context, container string, options types.CheckpointCreateOptions) error
	CheckpointDelete(ctx context.Context, container string, options types.CheckpointDeleteOptions) error
	CheckpointList(ctx context.Context, container string, options types.CheckpointListOptions) ([]types.Checkpoint, error)
}
