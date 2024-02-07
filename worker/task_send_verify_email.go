package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const TaskSendVerifyEmail = "task:send_verify_email"

type PayloadSendVerifyEmail struct {
	Username string `json:"username"`
}

func (distributor *RedisTaskDistributor) DistributeTaskSendVerifyEmail(
	ctx context.Context,
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option,
) error {
	// Since our payload is an object, we  will have to serialize it into JSON,
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to enqueue task: %w", err)
	}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")
	return nil
}

func (processor *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("failed to unmarshal payload: %w", asynq.SkipRetry)
	}

	// In some rare cases, anything could happen. So, one way to deal with that, is to always allow the task to retry.
	// And hopefully at that time, it will be successful, when the transaction has been committed.
	user, err := processor.store.GetUser(ctx, payload.Username)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	//verifyEmail, err := processor.store.CreateVerifyEmail(ctx, db.CreateVerifyEmailParams{
	//	Username:   user.Username,
	//	Email:      user.Email,
	//	SecretCode: util.RandomString(32),
	//})
	//if err != nil {
	//	return fmt.Errorf("failed to create verify email: %w", err)
	//}
	//
	//subject := "Welcome to Simple Bank"
	//// TODO: replace this URL with an environment variable that points to a front-end page
	//verifyUrl := fmt.Sprintf("http://localhost:8080/v1/verify_email?email_id=%d&secret_code=%s",
	//	verifyEmail.ID, verifyEmail.SecretCode)
	//content := fmt.Sprintf(`Hello %s,<br/>
	//Thank you for registering with us!<br/>
	//Please <a href="%s">click here</a> to verify your email address.<br/>
	//`, user.FullName, verifyUrl)
	//to := []string{user.Email}
	//
	//err = processor.mailer.SendEmail(subject, content, to, nil, nil, nil)
	//if err != nil {
	//	return fmt.Errorf("failed to send verify email: %w", err)
	//}

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("email", user.Email).Msg("processed task")
	return nil
}
