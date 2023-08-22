// Package local provides a local runtime
package local

import (
	"github.com/kong11213613/haremicro/runtime"
)

// NewRuntime returns a new local runtime
func NewRuntime(opts ...runtime.Option) runtime.Runtime {
	return runtime.NewRuntime(opts...)
}
