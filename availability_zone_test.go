package hazelcastcloud

import (
	"context"
	"fmt"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

func ExampleAvailabilityZoneService_list() {
	client, _, _ := New()
	availabilityZones, _, _ := client.AvailabilityZone.List(context.Background(), &models.AvailabilityZoneRequest{
		CloudProvider: "aws",
		Region:        "us-east-2",
		InstanceType:  "m5.large",
		InstanceCount: 1,
	})
	fmt.Printf("Results: %#v", availabilityZones)
	//Output:Result: &[]models.AvailabilityZone{models.AvailabilityZone{Name:"us-east-2a"}, models.AvailabilityZone{Name:"us-east-2b"}, models.AvailabilityZone{Name:"us-east-2c"}}
}
