package republish

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Yoga-Saputra/go-boilerplate/pkg/kafadapter"
	"github.com/hibiken/asynq"
)

type (
	// QPayload define payload.
	QPayload struct {
		ProviderCode string `json:"provider_code"`

		MemberID uint64 `json:"member_id"`
		Username string `json:"username"`
		PID      string `json:"p_id"`

		BID          string `json:"b_id"`
		GameCategory string `json:"game_category"`
		WebID        uint32 `json:"web_id"`
		Currency     string `json:"currency"`
		GameCode     string `json:"game_code"`

		Result string `json:"result"`
		Event  string `json:"event"`
		Match  string `json:"match"`
		Market string `json:"market"`

		ResettlementAmount float64 `json:"resettlement_amount"`
		Status             string  `json:"status"`
		BTime              string  `json:"b_time"`
		Opt                string  `json:"opt"`
	}
)

// Package constant
const (
	QueueName = "REPUBLISH-DATALAKE"
	TaskName  = QueueName + ":Pub"

	maxRetry   = 3
	kafkaTopic = "Republish-Datalake"
)

// Local variable
var (
	RepublishTasks = map[string]interface{}{
		TaskName: Handler,
	}

	client        *asynq.Client
	kafkaProducer kafadapter.Builder
)

// CreateClient create new asynq client.
func CreateClient(c *asynq.Client, kp kafadapter.Builder) {
	client = c
	kafkaProducer = kp
}

// DispatchRepublishDatalake to dispatching data kafka to the queue.
func DispatchRepublishDatalake(p *QPayload) (*asynq.TaskInfo, error) {
	// Do aggregation process using queue
	task, err := enqueue(p)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func Handler(c context.Context, t *asynq.Task) error {
	// Prepare the task payload
	var p *QPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %s -> %s", err.Error(), asynq.SkipRetry)
	}

	err := kafkaProducer.Publish(
		p.ProviderCode,
		[]kafadapter.Messages{
			{Key: []byte(p.ProviderCode), Value: t.Payload()},
		},
	)
	if err != nil {
		p.logger(fmt.Sprintf("(KafkaPublish) Error: %s", err.Error()))
		return err
	}

	p.logger("(KafkaPublish) message successfully published")
	return nil
}

// Enqueue job to the queue system.
func enqueue(p *QPayload) (*asynq.TaskInfo, error) {
	// Create task
	b, err := json.Marshal(p)
	if err != nil {
		p.logger(fmt.Sprintf("(Init) Failed marshaling the payload: %s", err.Error()))
		return nil, err
	}
	task := asynq.NewTask(TaskName, b)

	// Enqueue job
	info, err := client.Enqueue(
		task,
		asynq.Queue(QueueName),
		asynq.MaxRetry(maxRetry),
	)

	if err != nil {
		p.logger(fmt.Sprintf("(Init) Failed to enqueue task: %s", err.Error()))
		return nil, err
	}

	p.logger(fmt.Sprintf("(Init) Successfully enqueue task: %s", info.ID))
	return info, nil
}

// Helper to print custom logging.
func (p *QPayload) logger(m string) {
	prefix := fmt.Sprintf("PubDatalakeQ-%v", p.ProviderCode)

	log.Printf("[%s] %s", prefix, m)
}
