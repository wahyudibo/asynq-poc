package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/hibiken/asynq"

	"github.com/wahyudibo/asynqmon-poc/pkg/tasks"
)

func main() {
	redisAddr := "redis:6379"

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		r := rand.Intn(60)

		task, err := tasks.NewTask(i, time.Duration(r))
		if err != nil {
			log.Fatalf("could not create task: %v", err)
		}
		info, err := client.Enqueue(task, asynq.MaxRetry(10), asynq.Timeout(time.Minute), asynq.Retention(time.Minute))
		if err != nil {
			log.Fatalf("could not enqueue task: %v", err)
		}

		log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
	}

	log.Printf("All tasks enqueued")
}
