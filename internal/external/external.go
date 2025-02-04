package external

import (
	"context"
	"net/http"
	"time"
)

type External struct {
	Kafka interface {
		ProduceKafkaMessage(context.Context, []byte) error
	}
	User interface {
		Profile(context.Context, string) (*Response, error)
	}
}

func NewExternal() External {
	return External{
		Kafka: &KafkaExternal{},
		User: &UserExternal{
			httpClient: &http.Client{
				Timeout: 5 * time.Second,
			},
		},
	}
}
