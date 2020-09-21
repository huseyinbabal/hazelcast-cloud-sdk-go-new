package hazelcastcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCloudProviderServiceOp_List(t *testing.T) {
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

		if strings.Contains(request.Query, "cloudProviders") {
			fmt.Fprint(w, `{"data":{"response":[{"name":"aws"},{"name":"azure"},{"name":"google"}]}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ5dW51c0BoYXplbGNhc3QuY29tIiwicm9sZXMiOlt7InRlYW1JZCI6IjMiLCJhdXRob3JpdHkiOiJURUFNX0FETUlOIn0seyJ0ZWFtSWQiOiIxIiwiYXV0aG9yaXR5IjoiVEVBTV9GSU5BTkNFIn0seyJ0ZWFtSWQiOm51bGwsImF1dGhvcml0eSI6IkFETUlOIn0seyJ0ZWFtSWQiOiIxIiwiYXV0aG9yaXR5IjoiVEVBTV9BRE1JTiJ9LHsidGVhbUlkIjpudWxsLCJhdXRob3JpdHkiOiJVU0VSIn0seyJ0ZWFtSWQiOm51bGwsImF1dGhvcml0eSI6IkRFRElDQVRFRF9VU0VSIn0seyJ0ZWFtSWQiOiIyIiwiYXV0aG9yaXR5IjoiVEVBTV9BRE1JTiJ9LHsidGVhbUlkIjoiMiIsImF1dGhvcml0eSI6IlRFQU1fRklOQU5DRSJ9LHsidGVhbUlkIjpudWxsLCJhdXRob3JpdHkiOiJBQ0NPVU5USU5HIn1dLCJ0b2tlbiI6IjE1YjY5MWQxLThmOWUtNGQ4Zi04NzNkLTk4ZWI0NGU0ODk5NSIsImV4cCI6MTc1NzQyODg3MH0.HM3vLZbR4H8LIu0Quqm3dqwCj6V_XAYtaUGg5ZQkeefgvMA1LIoxJRyPgZYhJgJJ_aHPnBZ08wJwCrFADGHitA"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	cloudProviders, _, _ := NewCloudProviderService(client).List(context.TODO())

	//then
	assert.Len(t, *cloudProviders, 3)
	assert.Equal(t, (*cloudProviders)[0].Name, "aws")
	assert.Equal(t, (*cloudProviders)[1].Name, "azure")
	assert.Equal(t, (*cloudProviders)[2].Name, "google")
}

func ExampleCloudProviderService_list() {
	client, _, _ := New()
	cloudProviders, _, _ := client.CloudProvider.List(context.Background())
	fmt.Printf("Result: %#v", cloudProviders)
	//Output:Result: &[]models.CloudProvider{models.CloudProvider{Name:"aws", IsEnabledForStarter:true, IsEnabledForEnterprise:true}, models.CloudProvider{Name:"azure", IsEnabledForStarter:false, IsEnabledForEnterprise:true}, models.CloudProvider{Name:"google", IsEnabledForStarter:false, IsEnabledForEnterprise:false}}
}
