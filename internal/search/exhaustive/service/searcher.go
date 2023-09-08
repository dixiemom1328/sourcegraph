package service

import (
	"context"

	"github.com/sourcegraph/sourcegraph/internal/search"
	"github.com/sourcegraph/sourcegraph/internal/search/client"
	"github.com/sourcegraph/sourcegraph/internal/search/exhaustive/types"
)

func FromSearchClient(client client.SearchClient) NewSearcher {
	return newSearcherFunc(func(ctx context.Context, q string) (SearchQuery, error) {
		// TODO adjust NewSearch API to enforce the user passing in a user id.
		// IE do not rely on ctx actor since that could easily lead to a bug.
		inputs, err := client.Plan(
			ctx,
			"V3",
			nil,
			q,
			search.Precise,
			search.Streaming,
		)
		if err != nil {
			return nil, err
		}

		// TODO ensure plan is runnable by exhaustive

		return searchQuery{
			client: client,
			inputs: inputs,
		}, nil
	})
}

// TODO maybe reuse for the fake
type newSearcherFunc func(context.Context, string) (SearchQuery, error)

func (f newSearcherFunc) NewSearch(ctx context.Context, q string) (SearchQuery, error) {
	return f(ctx, q)
}

type searchQuery struct {
	client client.SearchClient
	inputs *search.Inputs
}

func (s searchQuery) RepositoryRevSpecs(context.Context) ([]types.RepositoryRevSpec, error) {
	return nil, nil
}

func (s searchQuery) ResolveRepositoryRevSpec(context.Context, types.RepositoryRevSpec) ([]types.RepositoryRevision, error) {
	return nil, nil
}

func (s searchQuery) Search(context.Context, types.RepositoryRevision, CSVWriter) error {
	return nil
}
