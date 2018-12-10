package daemon // import "github.com/sdslabs/docker/daemon"

import (
	"testing"

	"github.com/sdslabs/docker/api/types/network"
	"github.com/sdslabs/docker/errdefs"
	"gotest.tools/assert"
)

// Test case for 35752
func TestVerifyNetworkingConfig(t *testing.T) {
	name := "mynet"
	endpoints := make(map[string]*network.EndpointSettings, 1)
	endpoints[name] = nil
	nwConfig := &network.NetworkingConfig{
		EndpointsConfig: endpoints,
	}
	err := verifyNetworkingConfig(nwConfig)
	assert.Check(t, errdefs.IsInvalidParameter(err))
}
