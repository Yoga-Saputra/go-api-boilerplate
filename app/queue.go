package app

import (
	"context"
	"fmt"
	"time"

	"github.com/Yoga-Saputra/go-boilerplate/config"
	"github.com/Yoga-Saputra/go-boilerplate/pkg/kemu"
	"github.com/hibiken/asynq"

	taskaq "github.com/Yoga-Saputra/go-boilerplate/internal/job"
)

// Queue driver pointer value fo redis adapter
var Queue *asynq.Server

// Start queue server
func queueUp(args *AppArgs) {
	rdsAddr := fmt.Sprintf("%s:%v", config.Of.Queue.Redis.Host, config.Of.Queue.Redis.Port)

	// Create queue server
	createServer(rdsAddr)

	// Create queue client
	createClient(rdsAddr)
}

// Stop queue server
func queueDown() {
	printOutFinishTask("Wait until all queue tasks is finished...")
	printOutFinishTask("and closing current Queue Redis connection...")

	// Shutdown the server
	Queue.Shutdown()
}

// Create queue server
func createServer(addr string) {
	// Define asynq config
	aqcfg := asynq.Config{
		Concurrency:         config.Of.Queue.Option.Concurrency,
		ShutdownTimeout:     time.Duration(config.Of.Queue.Option.ShutdownTimeout) * time.Second,
		HealthCheckInterval: time.Duration(config.Of.Queue.Option.HealthCheckInterval) * time.Second,
		RetryDelayFunc: func(n int, e error, t *asynq.Task) time.Duration {
			return 2 * time.Second
		},
		StrictPriority: true,
	}

	// Set the queues
	queues := make(map[string]int, len(taskaq.RegiteredTask))
	for _, taq := range taskaq.RegiteredTask {
		queues[taq.QueueName] = taq.Priority
	}
	if len(queues) > 0 {
		aqcfg.Queues = queues
	}

	// Create new queue server
	qSrv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     addr,
			Password: config.Of.Queue.Redis.Password,
			DB:       config.Of.Queue.Redis.Database,
			PoolSize: config.Of.Queue.Redis.PoolSize,
		},
		aqcfg,
	)
	Queue = qSrv

	// Create new ServerMux
	mux := asynq.NewServeMux()

	// Assign registered task into mux
	for _, taq := range taskaq.RegiteredTask {
		for tn, hd := range taq.Tasks {
			handler, ok := hd.(func(context.Context, *asynq.Task) error)
			if !ok {
				panic(fmt.Sprintf("mismatch queue task handler: %v", hd))
			}

			mux.HandleFunc(tn, handler)
		}
	}

	// Start Queue Server using goroutine
	go func() {
		if err := qSrv.Start(mux); err != nil {
			printOutUp("ERROR QUEUE")
			panic(err)
		}
	}()

	printOutUp("New Queue Server successfully open")
}

// Create queue client
func createClient(addr string) {
	// Asynq client
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     addr,
		Password: config.Of.Queue.Redis.Password,
		DB:       config.Of.Queue.Redis.Database,
	})

	kemu := kemu.New()

	taskaq.CreateClient(KafADPT, client, Cache2.RCch, kemu)
	printOutUp("New Queue Client successfully open")
}
