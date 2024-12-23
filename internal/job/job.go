package taskaq

import (
	"github.com/Yoga-Saputra/go-boilerplate/internal/job/republish"
	"github.com/Yoga-Saputra/go-boilerplate/pkg/kafadapter"
	"github.com/Yoga-Saputra/go-boilerplate/pkg/kemu"
	"github.com/go-redis/cache/v8"
	"github.com/hibiken/asynq"
)

type taskAsynQ struct {
	QueueName string
	Priority  int

	Tasks map[string]interface{}
}

// Map of function that will be called on Up() method based on their order.
// If have new services, just create new file and their method and register here
var RegiteredTask = []taskAsynQ{

	{
		QueueName: republish.QueueName,
		Priority:  1,
		Tasks:     republish.RepublishTasks,
	},
}

// CreateClient create new asynq client
// and another needed value
func CreateClient(KafkaProducer kafadapter.Builder, cln *asynq.Client, cch *cache.Cache, k *kemu.Mutex) {

	republish.CreateClient(cln, KafkaProducer)
}
