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

func TestRegionServiceOp_List(t *testing.T) {
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

		if strings.Contains(request.Query, "regions") {
			fmt.Fprint(w, `{"data":{"response":[{"name":"us-east-1"},{"name":"us-east-2"},{"name":"us-west-1"},{"name":"us-west-2"},{"name":"ap-south-1"},{"name":"ap-northeast-2"},{"name":"ap-southeast-1"},{"name":"ap-southeast-2"},{"name":"ap-northeast-1"},{"name":"ca-central-1"},{"name":"eu-central-1"},{"name":"eu-west-1"},{"name":"eu-west-2"},{"name":"eu-west-3"},{"name":"sa-east-1"}]}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ5dW51c0BoYXplbGNhc3QuY29tIiwicm9sZXMiOlt7InRlYW1JZCI6IjMiLCJhdXRob3JpdHkiOiJURUFNX0FETUlOIn0seyJ0ZWFtSWQiOiIxIiwiYXV0aG9yaXR5IjoiVEVBTV9GSU5BTkNFIn0seyJ0ZWFtSWQiOm51bGwsImF1dGhvcml0eSI6IkFETUlOIn0seyJ0ZWFtSWQiOiIxIiwiYXV0aG9yaXR5IjoiVEVBTV9BRE1JTiJ9LHsidGVhbUlkIjpudWxsLCJhdXRob3JpdHkiOiJVU0VSIn0seyJ0ZWFtSWQiOm51bGwsImF1dGhvcml0eSI6IkRFRElDQVRFRF9VU0VSIn0seyJ0ZWFtSWQiOiIyIiwiYXV0aG9yaXR5IjoiVEVBTV9BRE1JTiJ9LHsidGVhbUlkIjoiMiIsImF1dGhvcml0eSI6IlRFQU1fRklOQU5DRSJ9LHsidGVhbUlkIjpudWxsLCJhdXRob3JpdHkiOiJBQ0NPVU5USU5HIn1dLCJ0b2tlbiI6IjE1YjY5MWQxLThmOWUtNGQ4Zi04NzNkLTk4ZWI0NGU0ODk5NSIsImV4cCI6MTc1NzQyODg3MH0.HM3vLZbR4H8LIu0Quqm3dqwCj6V_XAYtaUGg5ZQkeefgvMA1LIoxJRyPgZYhJgJJ_aHPnBZ08wJwCrFADGHitA"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	regions, _, _ := NewRegionService(client).List(context.TODO(), &models.RegionInput{})

	//then
	assert.Len(t, *regions, 15)
}

func ExampleRegionService_list() {
	client, _, _ := New()
	regions, _, _ := client.Region.List(context.Background(), &models.RegionInput{CloudProvider: "aws"})
	fmt.Printf("Result: %#v", regions)
	//Output:Result: &[]models.Region{models.Region{Name:"us-east-1", IsEnabledForStarter:false, IsEnabledForEnterprise:true}, models.Region{Name:"us-east-2", IsEnabledForStarter:false, IsEnabledForEnterprise:true}, models.Region{Name:"us-west-1", IsEnabledForStarter:false, IsEnabledForEnterprise:true}, models.Region{Name:"us-west-2", IsEnabledForStarter:true, IsEnabledForEnterprise:true}, models.Region{Name:"ap-south-1", IsEnabledForStarter:false, IsEnabledForEnterprise:true}, models.Region{Name:"ap-northeast-2", IsEnabledForStarter:false, IsEnabledForEnterprise:true}, models.Region{Name:"ap-southeast-1", IsEnabledForStarter:false, IsEnabledForEnterprise:true}, models.Region{Name:"ap-southeast-2", IsEnabledForStarter:false, IsEnabledForEnterprise:true}, models.Region{Name:"ap-northeast-1", IsEnabledForStarter:false, IsEnabledForEnterprise:true}, models.Region{Name:"ca-central-1", IsEnabledForStarter:false, IsEnabledForEnterprise:true}, models.Region{Name:"eu-central-1", IsEnabledForStarter:false, IsEnabledForEnterprise:true}, models.Region{Name:"eu-west-1", IsEnabledForStarter:false, IsEnabledForEnterprise:true}, models.Region{Name:"eu-west-2", IsEnabledForStarter:false, IsEnabledForEnterprise:true}, models.Region{Name:"eu-west-3", IsEnabledForStarter:false, IsEnabledForEnterprise:true}, models.Region{Name:"sa-east-1", IsEnabledForStarter:false, IsEnabledForEnterprise:true}}
}
