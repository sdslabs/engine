package image // import "github.com/sdslabs/docker/integration/image"

import (
	"context"
	"testing"

	"github.com/sdslabs/docker/api/types"
	"github.com/sdslabs/docker/api/types/versions"
	"github.com/sdslabs/docker/integration/internal/container"
	"github.com/sdslabs/docker/internal/test/request"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
	"gotest.tools/skip"
)

func TestCommitInheritsEnv(t *testing.T) {
	skip.If(t, versions.LessThan(testEnv.DaemonAPIVersion(), "1.36"), "broken in earlier versions")
	skip.If(t, testEnv.DaemonInfo.OSType == "windows", "FIXME")
	defer setupTest(t)()
	client := request.NewAPIClient(t)
	ctx := context.Background()

	cID1 := container.Create(t, ctx, client)

	commitResp1, err := client.ContainerCommit(ctx, cID1, types.ContainerCommitOptions{
		Changes:   []string{"ENV PATH=/bin"},
		Reference: "test-commit-image",
	})
	assert.NilError(t, err)

	image1, _, err := client.ImageInspectWithRaw(ctx, commitResp1.ID)
	assert.NilError(t, err)

	expectedEnv1 := []string{"PATH=/bin"}
	assert.Check(t, is.DeepEqual(expectedEnv1, image1.Config.Env))

	cID2 := container.Create(t, ctx, client, container.WithImage(image1.ID))

	commitResp2, err := client.ContainerCommit(ctx, cID2, types.ContainerCommitOptions{
		Changes:   []string{"ENV PATH=/usr/bin:$PATH"},
		Reference: "test-commit-image",
	})
	assert.NilError(t, err)

	image2, _, err := client.ImageInspectWithRaw(ctx, commitResp2.ID)
	assert.NilError(t, err)
	expectedEnv2 := []string{"PATH=/usr/bin:/bin"}
	assert.Check(t, is.DeepEqual(expectedEnv2, image2.Config.Env))
}
