package task

import (
	"fmt"

	"github.com/hibiken/asynq"
)

func NewUpdateCacheTask(repo string) *asynq.Task {
	fmt.Printf("Adding repo '%s' to queue\n", repo)
	payload := map[string]interface{}{"repoName": repo}
	return asynq.NewTask(UpdateCache, payload)
}
