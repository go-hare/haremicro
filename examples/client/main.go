package main

import (
	"context"
	"fmt"

	"github.com/kong11213613/haremicro"
	"github.com/kong11213613/haremicro/client"
	greeter "github.com/kong11213613/haremicro/examples/helloworld/proto"
	"github.com/kong11213613/haremicro/metadata"
	"github.com/kong11213613/haremicro/registry"
	"github.com/kong11213613/haremicro/registry/etcd"
)

func call(i int, c client.Client) {
	// Create new request to service go.micro.srv.example, method Example.Call
	req := c.NewRequest("helloworld", "Greeter.Hello", &greeter.Request{})

	// create context with metadata
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "john",
		"X-From-Id": "script",
	})

	rsp := &greeter.Response{}

	// Call service
	if err := c.Call(ctx, req, rsp); err != nil {
		fmt.Println("call err: ", err, rsp)
		return
	}

	fmt.Println("Call:", i, "rsp:", rsp)
}

func main() {
	service := haremicro.NewService(haremicro.Registry(etcd.NewRegistry(registry.Addrs([]string{"127.0.0.1:2379"}...))))
	service.Init()

	fmt.Println("\n--- Call example ---")
	for i := 0; i < 10; i++ {
		call(i, service.Client())
	}
}
