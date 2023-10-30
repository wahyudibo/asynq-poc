package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hibiken/asynq"
)

const (
	TypeTask = "task"
)

type Task struct {
	ID       int           `json:"id"`
	Duration time.Duration `json:"duration"`
}

func NewTask(id int, duration time.Duration) (*asynq.Task, error) {
	payload, err := json.Marshal(Task{ID: id, Duration: duration})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeTask, payload), nil
}

func (processor *Task) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p Task
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	log.Printf("working on a task ID: %d. Will take %ds to finish\n", p.ID, p.Duration)
	time.Sleep(p.Duration * time.Second)
	log.Printf("task ID: %d done\n", p.ID)

	return nil
}
