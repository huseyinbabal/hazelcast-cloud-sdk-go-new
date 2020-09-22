package main

import (
	"context"
	hazelcastcloud "github.com/huseyinbabal/hazelcast-cloud-sdk-go-new"
)

func main() {
	client, response, err := hazelcastcloud.NewFromCredentials("f3cbMEJIjvIc673XtlWgJPmjs", "eJumbXU55gWx4YM6RGfpDqTAcgwsawviiRTdU6dxocrsjvjGZaQcNU8ZOpHD")
	if err != nil {
		panic(err)
	}
	cp, response, err := client.StarterCluster.List(context.Background())
	if err != nil {
		panic(err)
	}
	print(response)
	print(cp)
}
