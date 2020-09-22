package hazelcastcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAvailabilityZoneServiceOp_List(t *testing.T) {
	//given
	serveMux := http.NewServeMux()
	server := httptest.NewServer(serveMux)
	defer server.Close()

	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodPost; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		var request GraphQLQuery
		json.NewDecoder(r.Body).Decode(&request)

		if strings.Contains(request.Query, "availabilityZones") {
			fmt.Fprint(w, `{"data":{"response":[{"name":"us-west-2a"},{"name":"us-west-2b"},{"name":"us-west-2c"}]}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ5dW51c0BoYXplbGNhc3QuY29tIiwicm9sZXMiOlt7InRlYW1JZCI6IjMiLCJhdXRob3JpdHkiOiJURUFNX0FETUlOIn0seyJ0ZWFtSWQiOiIxIiwiYXV0aG9yaXR5IjoiVEVBTV9GSU5BTkNFIn0seyJ0ZWFtSWQiOm51bGwsImF1dGhvcml0eSI6IkFETUlOIn0seyJ0ZWFtSWQiOiIxIiwiYXV0aG9yaXR5IjoiVEVBTV9BRE1JTiJ9LHsidGVhbUlkIjpudWxsLCJhdXRob3JpdHkiOiJVU0VSIn0seyJ0ZWFtSWQiOm51bGwsImF1dGhvcml0eSI6IkRFRElDQVRFRF9VU0VSIn0seyJ0ZWFtSWQiOiIyIiwiYXV0aG9yaXR5IjoiVEVBTV9BRE1JTiJ9LHsidGVhbUlkIjoiMiIsImF1dGhvcml0eSI6IlRFQU1fRklOQU5DRSJ9LHsidGVhbUlkIjpudWxsLCJhdXRob3JpdHkiOiJBQ0NPVU5USU5HIn1dLCJ0b2tlbiI6IjE1YjY5MWQxLThmOWUtNGQ4Zi04NzNkLTk4ZWI0NGU0ODk5NSIsImV4cCI6MTc1NzQyODg3MH0.HM3vLZbR4H8LIu0Quqm3dqwCj6V_XAYtaUGg5ZQkeefgvMA1LIoxJRyPgZYhJgJJ_aHPnBZ08wJwCrFADGHitA"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	service := NewAvailabilityZoneService(client)
	availabilityZones, _, _ := service.List(context.TODO(), &models.AvailabilityZoneInput{})

	//then
	assert.Len(t, *availabilityZones, 3)
	assert.Equal(t, (*availabilityZones)[0].Name, "us-west-2a")
	assert.Equal(t, (*availabilityZones)[1].Name, "us-west-2b")
	assert.Equal(t, (*availabilityZones)[2].Name, "us-west-2c")

}

func ExampleAvailabilityZoneService_list() {
	client, _, _ := New()
	availabilityZones, _, _ := client.AvailabilityZone.List(context.Background(), &models.AvailabilityZoneInput{
		CloudProvider: "aws",
		Region:        "us-east-2",
		InstanceType:  "m5.large",
		InstanceCount: 1,
	})
	fmt.Printf("Results: %#v", availabilityZones)
	//Output:Result: &[]models.AvailabilityZone{models.AvailabilityZone{Name:"us-east-2a"}, models.AvailabilityZone{Name:"us-east-2b"}, models.AvailabilityZone{Name:"us-east-2c"}}
}
