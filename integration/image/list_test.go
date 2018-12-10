package image // import "github.com/sdslabs/docker/integration/image"

import (
	"context"
	"testing"

	"github.com/sdslabs/docker/api/types"
	"github.com/sdslabs/docker/api/types/filters"
	"github.com/sdslabs/docker/internal/test/request"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

// Regression : #38171
func TestImagesFilterMultiReference(t *testing.T) {
	defer setupTest(t)()
	client := request.NewAPIClient(t)
	ctx := context.Background()

	name := "images_filter_multi_reference"
	repoTags := []string{
		name + ":v1",
		name + ":v2",
		name + ":v3",
		name + ":v4",
	}

	for _, repoTag := range repoTags {
		err := client.ImageTag(ctx, "busybox:latest", repoTag)
		assert.NilError(t, err)
	}

	filter := filters.NewArgs()
	filter.Add("reference", repoTags[0])
	filter.Add("reference", repoTags[1])
	filter.Add("reference", repoTags[2])
	options := types.ImageListOptions{
		All:     false,
		Filters: filter,
	}
	images, err := client.ImageList(ctx, options)
	assert.NilError(t, err)

	assert.Check(t, is.Equal(len(images[0].RepoTags), 3))
	for _, repoTag := range images[0].RepoTags {
		if repoTag != repoTags[0] && repoTag != repoTags[1] && repoTag != repoTags[2] {
			t.Errorf("list images doesn't match any repoTag we expected, repoTag: %s", repoTag)
		}
	}
}
