package affected_apps

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/zapier/kubechecks/pkg/config"
	repo2 "github.com/zapier/kubechecks/pkg/repo"
)

func TestCreateNewMatcherWithNilVcsMap(t *testing.T) {
	// setup
	var (
		vcsMap config.VcsToArgoMap
		repo   repo2.Repo
		path   string
	)

	// run test
	matcher, err := NewArgocdMatcher(vcsMap, &repo, path)
	require.NoError(t, err)

	// verify results
	require.Nil(t, matcher.appsDirectory)
}

func TestFindAffectedAppsWithNilAppsDirectory(t *testing.T) {
	// setup
	var (
		ctx        = context.TODO()
		changeList = []string{"/go.mod"}
	)

	matcher := ArgocdMatcher{}
	items, err := matcher.AffectedApps(ctx, changeList)

	// verify results
	require.NoError(t, err)
	require.Len(t, items.Applications, 0)
	require.Len(t, items.ApplicationSets, 0)
}
