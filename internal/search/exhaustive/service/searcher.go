package service

import (
	"context"
	"fmt"

	"github.com/sourcegraph/sourcegraph/internal/search"
	"github.com/sourcegraph/sourcegraph/internal/search/client"
	"github.com/sourcegraph/sourcegraph/internal/search/exhaustive/types"
	"github.com/sourcegraph/sourcegraph/internal/search/job"
	"github.com/sourcegraph/sourcegraph/internal/search/job/jobutil"
	"github.com/sourcegraph/sourcegraph/internal/search/job/printer"
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

		planJob, err := jobutil.NewPlanJob(inputs, inputs.Plan)
		if err != nil {
			return nil, err
		}

		fmt.Println(printer.SexpVerbose(planJob, job.VerbosityMax, true))

		return searchQuery{
			client:  client,
			planJob: planJob,
		}, nil
	})
}

// TODO maybe reuse for the fake
type newSearcherFunc func(context.Context, string) (SearchQuery, error)

func (f newSearcherFunc) NewSearch(ctx context.Context, q string) (SearchQuery, error) {
	return f(ctx, q)
}

type searchQuery struct {
	client  client.SearchClient
	planJob job.Job
}

func (s searchQuery) RepositoryRevSpecs(context.Context) ([]types.RepositoryRevSpec, error) {
	return nil, nil
}

func (s searchQuery) ResolveRepositoryRevSpec(context.Context, types.RepositoryRevSpec) ([]types.RepositoryRevision, error) {
	return nil, nil
}

func (s searchQuery) Search(ctx context.Context, reporev types.RepositoryRevision, w CSVWriter) error {
	//planJob.Run(ctx, s.JobClients(), stream)
	return nil
}
