package asynq

import (
	"context"
	"time"

	"github.com/go-hare/haremicro/broker"
)

const (
	// ClusterType means redis cluster.
	ClusterType = "cluster"
	// NodeType means redis node.
	NodeType = "node"

	DefaultDB = 0
)

type optionsKeyType struct{}
type publishKeyType struct{}
type subscribeKeyType struct{}

var (
	optionsKey   = optionsKeyType{}
	publishKey   = publishKeyType{}
	subscribeKey = subscribeKeyType{}
)

type brokerOptions struct {
	service  string
	nodeType string
	pass     string
	db       int
	tls      bool
	queues   map[string]int
}

// Service customizes the given service name
func Service(service string) broker.Option {
	return func(o *broker.Options) {
		bo := o.Context.Value(optionsKey).(*brokerOptions)
		bo.service = service
	}
}

// Cluster customizes the given Redis as a cluster
func Cluster() broker.Option {
	return func(o *broker.Options) {
		bo := o.Context.Value(optionsKey).(*brokerOptions)
		bo.nodeType = ClusterType
	}
}

// Pass is redis password
func Pass(pass string) broker.Option {
	return func(o *broker.Options) {
		bo := o.Context.Value(optionsKey).(*brokerOptions)
		bo.pass = pass
	}
}

// DB is the redis db to use
func DB(db int) broker.Option {
	return func(o *broker.Options) {
		bo := o.Context.Value(optionsKey).(*brokerOptions)
		bo.db = db
	}
}

// Tls if open tls
func Tls(b bool) broker.Option {
	return func(o *broker.Options) {
		bo := o.Context.Value(optionsKey).(*brokerOptions)
		bo.tls = b
	}
}

// Queues multiple queues with different priority level
func Queues(queues map[string]int) broker.Option {
	return func(o *broker.Options) {
		bo := o.Context.Value(optionsKey).(*brokerOptions)
		bo.queues = queues
	}
}

type publishOptions struct {
	Service   string
	Queue     string
	Opr       string
	ProcessIn time.Duration
	Retention time.Duration
}

func getPublishOptions(o *broker.PublishOptions) *publishOptions {
	if o.Context == nil {
		o.Context = context.Background()
	}
	if v, ok := o.Context.Value(publishKey).(*publishOptions); ok {
		return v
	} else {
		po := &publishOptions{}
		o.Context = context.WithValue(o.Context, publishKey, po)
		return po
	}
}

// PubService publish with service name
func PubService(service string) broker.PublishOption {
	return func(o *broker.PublishOptions) {
		po := getPublishOptions(o)
		po.Service = service
	}
}

// Queue publish with priority queue
func Queue(queue string) broker.PublishOption {
	return func(o *broker.PublishOptions) {
		po := getPublishOptions(o)
		po.Queue = queue
	}
}

// PubOpr pub opr
func PubOpr(opr string) broker.PublishOption {
	return func(o *broker.PublishOptions) {
		po := getPublishOptions(o)
		po.Opr = opr
	}
}

// ProcessIn schedule tasks to be processed in the future
func ProcessIn(t time.Duration) broker.PublishOption {
	return func(o *broker.PublishOptions) {
		po := getPublishOptions(o)
		po.ProcessIn = t
	}
}

// Retention task to be kept in the queue
func Retention(t time.Duration) broker.PublishOption {
	return func(o *broker.PublishOptions) {
		po := getPublishOptions(o)
		po.Retention = t
	}
}

type subscribeOptions struct {
	Opr string
}

// SubOpr sub opr
func SubOpr(opr string) broker.SubscribeOption {
	return func(o *broker.SubscribeOptions) {
		po := o.Context.Value(subscribeKey).(*subscribeOptions)
		po.Opr = opr
	}
}
