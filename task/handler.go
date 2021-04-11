package task

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/jahangiranwari/cache-service/cache"
	"github.com/jahangiranwari/cache-service/httputil"
)

func HandleUpdateCacheTask(ctx context.Context, t *asynq.Task) error {
	repo, err := t.Payload.GetString("repoName")
	if err != nil {
		return err
	}
	value := httputil.GetGitHubRepo(repo)
	cache.Save(repo, value)
	return nil
}
