package redis

import (
	"github.com/kong11213613/haremicro/broker"
)

var (
	optionsKey = optionsKeyType{}
)

// options contain additional options for the broker.
type brokerOptions struct {
	password string
	db       int
}

type optionsKeyType struct{}

func Password(password string) broker.Option {
	return func(o *broker.Options) {
		bo := o.Context.Value(optionsKey).(*brokerOptions)
		bo.password = password
	}
}

func DB(db int) broker.Option {
	return func(o *broker.Options) {
		bo := o.Context.Value(optionsKey).(*brokerOptions)
		bo.db = db
	}
}
