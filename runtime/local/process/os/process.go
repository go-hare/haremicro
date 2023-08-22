// Package os runs processes locally
package os

import (
	"github.com/kong11213613/haremicro/runtime/local/process"
)

type Process struct{}

func NewProcess(opts ...process.Option) process.Process {
	return &Process{}
}
