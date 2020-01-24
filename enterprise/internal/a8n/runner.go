package a8n

import (
	"context"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/sourcegraph/sourcegraph/cmd/frontend/graphqlbackend"
	"github.com/sourcegraph/sourcegraph/internal/a8n"
	"github.com/sourcegraph/sourcegraph/internal/api"
	"github.com/sourcegraph/sourcegraph/internal/env"
	"gopkg.in/inconshreveable/log15.v2"
)

// maxWorkers defines the maximum number of repositories over which a Runner
// executes CampaignJobs in parallel.
var maxWorkers = env.Get("A8N_MAX_WORKERS", "8", "maximum number of repositories run campaigns over in parallel")

const defaultWorkerCount = 8

// A Runner executes a CampaignPlan.
type Runner struct {
	store         *Store
	defaultBranch repoDefaultBranch
	clock         func() time.Time

	started bool
}

// repoDefaultBranch takes in a RepositoryResolver and returns the ref name of
// the repositories default branch and its target commit ID.
type repoDefaultBranch func(ctx context.Context, repo *graphqlbackend.RepositoryResolver) (string, api.CommitID, error)

// ErrNoDefaultBranch is returned by a repoDefaultBranch when no default branch
// could be determined for a given repo.
var ErrNoDefaultBranch = errors.New("could not determine default branch")

// defaultRepoDefaultBranch is an implementation of repoDefaultBranch that uses
// methods defined on RepositoryResolver to talk to gitserver to determine a
// repository's default branch and its target commit ID.
var defaultRepoDefaultBranch = func(ctx context.Context, repo *graphqlbackend.RepositoryResolver) (string, api.CommitID, error) {
	var branch string
	var commitID api.CommitID

	defaultBranch, err := repo.DefaultBranch(ctx)
	if err != nil {
		return branch, commitID, err
	}
	if defaultBranch == nil {
		return branch, commitID, ErrNoDefaultBranch
	}
	branch = defaultBranch.Name()

	commit, err := defaultBranch.Target().Commit(ctx)
	if err != nil {
		return branch, commitID, err
	}

	commitID = api.CommitID(commit.OID())
	return branch, commitID, nil
}

// NewRunner returns a Runner for a given CampaignType.
func NewRunner(store *Store, defaultBranch repoDefaultBranch) *Runner {
	return NewRunnerWithClock(store, defaultBranch, func() time.Time {
		return time.Now().UTC().Truncate(time.Microsecond)
	})
}

// NewRunnerWithClock returns a Runner for a given CampaignType with the given clock used
// to generate timestamps
func NewRunnerWithClock(store *Store, defaultBranch repoDefaultBranch, clock func() time.Time) *Runner {
	runner := &Runner{
		store:         store,
		defaultBranch: defaultBranch,
		clock:         clock,
	}
	if runner.defaultBranch == nil {
		runner.defaultBranch = defaultRepoDefaultBranch
	}

	return runner
}

// RunChangesetJobs should run in a background goroutine and is responsible
// for finding pending jobs and running them.
// ctx should be canceled to terminate the function
func RunChangesetJobs(ctx context.Context, s *Store, clock func() time.Time, gitClient GitserverClient, backoffDuration time.Duration) {
	workerCount, err := strconv.Atoi(maxWorkers)
	if err != nil {
		log15.Error("Parsing max worker count, falling back to default of 8", "err", err)
		workerCount = defaultWorkerCount
	}
	process := func(ctx context.Context, s *Store, job a8n.ChangesetJob) error {
		c, err := s.GetCampaign(ctx, GetCampaignOpts{
			ID: job.CampaignID,
		})
		if err != nil {
			return errors.Wrap(err, "getting campaign")
		}
		_ = RunChangesetJob(ctx, clock, s, gitClient, nil, c, &job)
		// We ignore the error here so that we don't roll back the transaction
		// RunChangesetJob will save the error in the job row
		return nil
	}
	worker := func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				didRun, err := s.ProcessPendingChangesetJobs(context.Background(), process)
				if err != nil {
					log15.Error("Running changeset job", "err", err)
				}
				// Back off on error or when no jobs available
				if err != nil || !didRun {
					time.Sleep(backoffDuration)
				}
			}
		}
	}
	for i := 0; i < workerCount; i++ {
		go worker()
	}
}
