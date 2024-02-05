package worker

import (
	"context"
	"github.com/hibiken/asynq"
)

/*
This file will contain the codes to create tasks, and distribute them to the workers via Redis queue.
*/
type TaskDistributor interface {
	DistributeTaskSendVerifyEmail(
		ctx context.Context,
		payload *PayloadSendVerifyEmail,
		opts ...asynq.Option,
	) error
}

type RedisTaskDistributor struct {
	// We will use this client later to send the task to redis queue.
	client *asynq.Client
}

func NewRedisTaskDistributor(redisOpt asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOpt)
	return &RedisTaskDistributor{
		client: client,
	}
}
