package registry

import (
	"go.uber.org/zap"
	"gopkg.in/resty.v1"
)

// Registry needs a comment
type Registry struct {
	client *resty.Client
	log    *zap.Logger
}

// Dial creates a Registry client
func Dial(url, username, password string) (*Registry, error) {
	logger, _ := zap.NewProduction()
	return &Registry{
		client: resty.New().SetHostURL(url).SetBasicAuth(username, password),
		log:    logger,
	}, nil
}
