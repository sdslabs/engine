package volume // import "github.com/sdslabs/docker/api/server/router/volume"

import (
	"context"

	"github.com/sdslabs/docker/volume/service/opts"
	// TODO return types need to be refactored into pkg
	"github.com/sdslabs/docker/api/types"
	"github.com/sdslabs/docker/api/types/filters"
)

// Backend is the methods that need to be implemented to provide
// volume specific functionality
type Backend interface {
	List(ctx context.Context, filter filters.Args) ([]*types.Volume, []string, error)
	Get(ctx context.Context, name string, opts ...opts.GetOption) (*types.Volume, error)
	Create(ctx context.Context, name, driverName string, opts ...opts.CreateOption) (*types.Volume, error)
	Remove(ctx context.Context, name string, opts ...opts.RemoveOption) error
	Prune(ctx context.Context, pruneFilters filters.Args) (*types.VolumesPruneReport, error)
}
