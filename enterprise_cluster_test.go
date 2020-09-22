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

func TestEnterpriseClusterServiceOp_Create(t *testing.T) {
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

		if strings.Contains(request.Query, "createEnterpriseCluster") {
			fmt.Fprint(w, `{"data":{"response":{"id":"123456","customerId":17000,"teamId":null,"name":"test-cluster","password":"e6838c596a0342d4918cf89a8d071023","port":34005,"hazelcastVersion":"3.12.2-4","isAutoScalingEnabled":false,"isHotBackupEnabled":false}}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ5dW51c0BoYXplbGNhc3QuY29tIiwicm9sZXMiOlt7InRlYW1JZCI6IjMiLCJhdXRob3JpdHkiOiJURUFNX0FETUlOIn0seyJ0ZWFtSWQiOiIxIiwiYXV0aG9yaXR5IjoiVEVBTV9GSU5BTkNFIn0seyJ0ZWFtSWQiOm51bGwsImF1dGhvcml0eSI6IkFETUlOIn0seyJ0ZWFtSWQiOiIxIiwiYXV0aG9yaXR5IjoiVEVBTV9BRE1JTiJ9LHsidGVhbUlkIjpudWxsLCJhdXRob3JpdHkiOiJVU0VSIn0seyJ0ZWFtSWQiOm51bGwsImF1dGhvcml0eSI6IkRFRElDQVRFRF9VU0VSIn0seyJ0ZWFtSWQiOiIyIiwiYXV0aG9yaXR5IjoiVEVBTV9BRE1JTiJ9LHsidGVhbUlkIjoiMiIsImF1dGhvcml0eSI6IlRFQU1fRklOQU5DRSJ9LHsidGVhbUlkIjpudWxsLCJhdXRob3JpdHkiOiJBQ0NPVU5USU5HIn1dLCJ0b2tlbiI6IjE1YjY5MWQxLThmOWUtNGQ4Zi04NzNkLTk4ZWI0NGU0ODk5NSIsImV4cCI6MTc1NzQyODg3MH0.HM3vLZbR4H8LIu0Quqm3dqwCj6V_XAYtaUGg5ZQkeefgvMA1LIoxJRyPgZYhJgJJ_aHPnBZ08wJwCrFADGHitA"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.CreateEnterpriseClusterInput{}

	//when
	clusterResponse, _, _ := NewEnterpriseClusterService(client).Create(context.TODO(), request)

	//then
	assert.Equal(t, (*clusterResponse).Id, "123456")
	assert.Equal(t, (*clusterResponse).CustomerId, 17000)
	assert.Equal(t, (*clusterResponse).Name, "test-cluster")
	assert.Equal(t, (*clusterResponse).Password, "e6838c596a0342d4918cf89a8d071023")
	assert.Equal(t, (*clusterResponse).Port, 34005)
	assert.Equal(t, (*clusterResponse).HazelcastVersion, "3.12.2-4")
	assert.False(t, (*clusterResponse).IsAutoScalingEnabled)
	assert.False(t, (*clusterResponse).IsHotBackupEnabled)
}

func TestEnterpriseClusterServiceOp_List(t *testing.T) {
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

		if strings.Contains(request.Query, "cluster") {
			fmt.Fprint(w, `{"data":{"response":[{"id":"427","name":"demo"},{"id":"429","name":"demo-play2"},{"id":"437","name":"demo-play3"},{"id":"438","name":"demo-sdk"},{"id":"439","name":"mycluster"},{"id":"445","name":"test-cluster"}]}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ5dW51c0BoYXplbGNhc3QuY29tIiwicm9sZXMiOlt7InRlYW1JZCI6IjMiLCJhdXRob3JpdHkiOiJURUFNX0FETUlOIn0seyJ0ZWFtSWQiOiIxIiwiYXV0aG9yaXR5IjoiVEVBTV9GSU5BTkNFIn0seyJ0ZWFtSWQiOm51bGwsImF1dGhvcml0eSI6IkFETUlOIn0seyJ0ZWFtSWQiOiIxIiwiYXV0aG9yaXR5IjoiVEVBTV9BRE1JTiJ9LHsidGVhbUlkIjpudWxsLCJhdXRob3JpdHkiOiJVU0VSIn0seyJ0ZWFtSWQiOm51bGwsImF1dGhvcml0eSI6IkRFRElDQVRFRF9VU0VSIn0seyJ0ZWFtSWQiOiIyIiwiYXV0aG9yaXR5IjoiVEVBTV9BRE1JTiJ9LHsidGVhbUlkIjoiMiIsImF1dGhvcml0eSI6IlRFQU1fRklOQU5DRSJ9LHsidGVhbUlkIjpudWxsLCJhdXRob3JpdHkiOiJBQ0NPVU5USU5HIn1dLCJ0b2tlbiI6IjE1YjY5MWQxLThmOWUtNGQ4Zi04NzNkLTk4ZWI0NGU0ODk5NSIsImV4cCI6MTc1NzQyODg3MH0.HM3vLZbR4H8LIu0Quqm3dqwCj6V_XAYtaUGg5ZQkeefgvMA1LIoxJRyPgZYhJgJJ_aHPnBZ08wJwCrFADGHitA"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	clusterResponses, _, _ := NewEnterpriseClusterService(client).List(context.TODO())

	//then
	assert.Len(t, *(clusterResponses), 6)
}

func TestEnterpriseClusterServiceOp_Get(t *testing.T) {
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

		if strings.Contains(request.Query, "cluster") {
			fmt.Fprint(w, `{"data":{"response":{"id":"123456","customerId":17000,"teamId":null,"name":"test-cluster","password":"e6838c596a0342d4918cf89a8d071023","port":34005,"hazelcastVersion":"3.12.2-4","isAutoScalingEnabled":false,"isHotBackupEnabled":false}}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ5dW51c0BoYXplbGNhc3QuY29tIiwicm9sZXMiOlt7InRlYW1JZCI6IjMiLCJhdXRob3JpdHkiOiJURUFNX0FETUlOIn0seyJ0ZWFtSWQiOiIxIiwiYXV0aG9yaXR5IjoiVEVBTV9GSU5BTkNFIn0seyJ0ZWFtSWQiOm51bGwsImF1dGhvcml0eSI6IkFETUlOIn0seyJ0ZWFtSWQiOiIxIiwiYXV0aG9yaXR5IjoiVEVBTV9BRE1JTiJ9LHsidGVhbUlkIjpudWxsLCJhdXRob3JpdHkiOiJVU0VSIn0seyJ0ZWFtSWQiOm51bGwsImF1dGhvcml0eSI6IkRFRElDQVRFRF9VU0VSIn0seyJ0ZWFtSWQiOiIyIiwiYXV0aG9yaXR5IjoiVEVBTV9BRE1JTiJ9LHsidGVhbUlkIjoiMiIsImF1dGhvcml0eSI6IlRFQU1fRklOQU5DRSJ9LHsidGVhbUlkIjpudWxsLCJhdXRob3JpdHkiOiJBQ0NPVU5USU5HIn1dLCJ0b2tlbiI6IjE1YjY5MWQxLThmOWUtNGQ4Zi04NzNkLTk4ZWI0NGU0ODk5NSIsImV4cCI6MTc1NzQyODg3MH0.HM3vLZbR4H8LIu0Quqm3dqwCj6V_XAYtaUGg5ZQkeefgvMA1LIoxJRyPgZYhJgJJ_aHPnBZ08wJwCrFADGHitA"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.GetEnterpriseClusterInput{}

	//when
	clusterResponse, _, _ := NewEnterpriseClusterService(client).Get(context.TODO(), request)

	//then
	assert.Equal(t, (*clusterResponse).Id, "123456")
	assert.Equal(t, (*clusterResponse).CustomerId, 17000)
	assert.Equal(t, (*clusterResponse).Name, "test-cluster")
	assert.Equal(t, (*clusterResponse).Password, "e6838c596a0342d4918cf89a8d071023")
	assert.Equal(t, (*clusterResponse).Port, 34005)
	assert.Equal(t, (*clusterResponse).HazelcastVersion, "3.12.2-4")
	assert.False(t, (*clusterResponse).IsAutoScalingEnabled)
	assert.False(t, (*clusterResponse).IsHotBackupEnabled)
}

func TestEnterpriseClusterServiceOp_Delete(t *testing.T) {
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

		if strings.Contains(request.Query, "deleteCluster") {
			fmt.Fprint(w, `{"data":{"response":{"clusterId":123456}}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"eyJhbGciOiJIUzUxMiJ9.eyJzdWIiOiJ5dW51c0BoYXplbGNhc3QuY29tIiwicm9sZXMiOlt7InRlYW1JZCI6IjMiLCJhdXRob3JpdHkiOiJURUFNX0FETUlOIn0seyJ0ZWFtSWQiOiIxIiwiYXV0aG9yaXR5IjoiVEVBTV9GSU5BTkNFIn0seyJ0ZWFtSWQiOm51bGwsImF1dGhvcml0eSI6IkFETUlOIn0seyJ0ZWFtSWQiOiIxIiwiYXV0aG9yaXR5IjoiVEVBTV9BRE1JTiJ9LHsidGVhbUlkIjpudWxsLCJhdXRob3JpdHkiOiJVU0VSIn0seyJ0ZWFtSWQiOm51bGwsImF1dGhvcml0eSI6IkRFRElDQVRFRF9VU0VSIn0seyJ0ZWFtSWQiOiIyIiwiYXV0aG9yaXR5IjoiVEVBTV9BRE1JTiJ9LHsidGVhbUlkIjoiMiIsImF1dGhvcml0eSI6IlRFQU1fRklOQU5DRSJ9LHsidGVhbUlkIjpudWxsLCJhdXRob3JpdHkiOiJBQ0NPVU5USU5HIn1dLCJ0b2tlbiI6IjE1YjY5MWQxLThmOWUtNGQ4Zi04NzNkLTk4ZWI0NGU0ODk5NSIsImV4cCI6MTc1NzQyODg3MH0.HM3vLZbR4H8LIu0Quqm3dqwCj6V_XAYtaUGg5ZQkeefgvMA1LIoxJRyPgZYhJgJJ_aHPnBZ08wJwCrFADGHitA"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))
	request := &models.ClusterDeleteInput{}

	//when
	clusterResponse, _, _ := NewEnterpriseClusterService(client).Delete(context.TODO(), request)

	//then
	assert.Equal(t, (*clusterResponse).ClusterId, 123456)
}

func ExampleEnterpriseClusterService_create() {
	client, _, _ := New()
	cluster, _, _ := client.EnterpriseCluster.Create(context.Background(), &models.CreateEnterpriseClusterInput{
		Name:                  "enterprise-cluster",
		CloudProvider:         "aws",
		Region:                "us-east-2",
		Zones:                 []string{"us-east-2a", "us-east-2b", "us-east-2c"},
		InstanceType:          "m5.large",
		InstancePerZone:       1,
		HazelcastVersion:      "3.12.6",
		IsPublicAccessEnabled: true,
		CidrBlock:             "10.18.0.0/16",
		NativeMemory:          3,
		IsAutoScalingEnabled:  false,
		IsHotRestartEnabled:   false,
		IsHotBackupEnabled:    false,
		IsTLSEnabled:          false,
		DataStructure:         models.DataStructureInput{},
	})
	fmt.Printf("Result: #%v ", cluster)
	//Output:Result: &models.Cluster{Id:"53807", CustomerId:10090, Name:"enterprise-cluster", Password:"e75140b884db4e488ecfb77fca92988c", Port:34023, HazelcastVersion:"3.12.6", IsAutoScalingEnabled:false, IsHotBackupEnabled:false, IsHotRestartEnabled:false, IsIpWhitelistEnabled:false, IsTlsEnabled:false, ProductType:struct { Name models.ProductTypeName "json:\"name\""; IsFree bool "json:\"isFree\"" }{Name:"Enterprise", IsFree:false}, State:"PENDING", CreatedAt:"2020-09-08T07:39:38.000Z", StartedAt:"2020-09-08T07:39:38.000Z", StoppedAt:"", Progress:struct { Status string "json:\"status\""; TotalItemCount int "json:\"totalItemCount\""; CompletedItemCount int "json:\"completedItemCount\"" }{Status:"Preparing", TotalItemCount:4, CompletedItemCount:0}, CloudProvider:struct { Name string "json:\"name\""; Region string "json:\"region\""; AvailabilityZones []string "json:\"availabilityZones\"" }{Name:"aws", Region:"us-east-2", AvailabilityZones:[]string{"us-east-2c", "us-east-2a", "us-east-2b"}}, DiscoveryTokens:[]models.DiscoveryToken{models.DiscoveryToken{Source:"public", Token:"RA6aaSPU9B0PouzK9VRYAISRYBFPCmqVl9E5lU737pyCat0NWQ"}}, Specs:struct { TotalMemory float64 "json:\"totalMemory\""; HeapMemory int "json:\"heapMemory\""; NativeMemory int "json:\"nativeMemory\""; Cpu int "json:\"cpu\""; InstanceType string "json:\"instanceType\""; InstancePerZone int "json:\"instancePerZone\"" }{TotalMemory:21, HeapMemory:4, NativeMemory:3, Cpu:1500, InstanceType:"m5.large", InstancePerZone:1}, Networking:struct { Type string "json:\"type\""; CidrBlock string "json:\"cidrBlock\""; Peering struct { IsEnabled bool "json:\"is_enabled\"" } "json:\"peering\""; PrivateLink struct { Url string "json:\"url\""; State string "json:\"state\"" } "json:\"privateLink\"" }{Type:"PUBLIC", CidrBlock:"10.18.0.0/16", Peering:struct { IsEnabled bool "json:\"is_enabled\"" }{IsEnabled:false}, PrivateLink:struct { Url string "json:\"url\""; State string "json:\"state\"" }{Url:"", State:"NOT_ACTIVE"}}, DataStructures:models.DataStructureResponse{MapConfigs:[]models.MapConfigResponse{}, JCacheConfigs:[]models.JCacheConfigResponse{}, ReplicatedMapConfigs:[]models.ReplicatedMapConfigResponse{}, QueueConfigs:[]models.QueueConfigResponse{}, SetConfigs:[]models.SetConfigResponse{}, ListConfigs:[]models.ListConfigResponse{}, TopicConfigs:[]models.TopicConfigResponse{}, MultiMapConfigs:[]models.MultiMapConfigResponse{}, RingBufferConfigs:[]models.RingBufferConfigResponse{}, ReliableTopicConfigs:[]models.ReliableTopicConfigResponse{}}}
}

func ExampleEnterpriseClusterService_list() {
	client, _, _ := New()
	clusters, _, _ := client.EnterpriseCluster.List(context.Background())
	fmt.Printf("Result: #%v ", clusters)
	//Output:Result: &[]models.Cluster{models.Cluster{Id:"53807", CustomerId:10090, Name:"enterprise-cluster", Password:"", Port:34023, HazelcastVersion:"3.12.6", IsAutoScalingEnabled:false, IsHotBackupEnabled:false, IsHotRestartEnabled:false, IsIpWhitelistEnabled:false, IsTlsEnabled:false, ProductType:struct { Name models.ProductTypeName "json:\"name\""; IsFree bool "json:\"isFree\"" }{Name:"Enterprise", IsFree:false}, State:"PENDING", CreatedAt:"2020-09-08T07:39:38.000Z", StartedAt:"2020-09-08T07:39:38.000Z", StoppedAt:"", Progress:struct { Status string "json:\"status\""; TotalItemCount int "json:\"totalItemCount\""; CompletedItemCount int "json:\"completedItemCount\"" }{Status:"Validating Environment", TotalItemCount:4, CompletedItemCount:2}, CloudProvider:struct { Name string "json:\"name\""; Region string "json:\"region\""; AvailabilityZones []string "json:\"availabilityZones\"" }{Name:"aws", Region:"us-east-2", AvailabilityZones:[]string(nil)}, DiscoveryTokens:[]models.DiscoveryToken{}, Specs:struct { TotalMemory float64 "json:\"totalMemory\""; HeapMemory int "json:\"heapMemory\""; NativeMemory int "json:\"nativeMemory\""; Cpu int "json:\"cpu\""; InstanceType string "json:\"instanceType\""; InstancePerZone int "json:\"instancePerZone\"" }{TotalMemory:21, HeapMemory:4, NativeMemory:3, Cpu:1500, InstanceType:"m5.large", InstancePerZone:1}, Networking:struct { Type string "json:\"type\""; CidrBlock string "json:\"cidrBlock\""; Peering struct { IsEnabled bool "json:\"is_enabled\"" } "json:\"peering\""; PrivateLink struct { Url string "json:\"url\""; State string "json:\"state\"" } "json:\"privateLink\"" }{Type:"PUBLIC", CidrBlock:"10.18.0.0/16", Peering:struct { IsEnabled bool "json:\"is_enabled\"" }{IsEnabled:false}, PrivateLink:struct { Url string "json:\"url\""; State string "json:\"state\"" }{Url:"", State:"NOT_ACTIVE"}}, DataStructures:models.DataStructureResponse{MapConfigs:[]models.MapConfigResponse{}, JCacheConfigs:[]models.JCacheConfigResponse{}, ReplicatedMapConfigs:[]models.ReplicatedMapConfigResponse{}, QueueConfigs:[]models.QueueConfigResponse{}, SetConfigs:[]models.SetConfigResponse{}, ListConfigs:[]models.ListConfigResponse{}, TopicConfigs:[]models.TopicConfigResponse{}, MultiMapConfigs:[]models.MultiMapConfigResponse{}, RingBufferConfigs:[]models.RingBufferConfigResponse(nil), ReliableTopicConfigs:[]models.ReliableTopicConfigResponse{}}}}
}

func ExampleEnterpriseClusterService_get() {
	client, _, _ := New()
	cluster, _, _ := client.EnterpriseCluster.Get(context.Background(), &models.GetEnterpriseClusterInput{ClusterId: "53807"})
	fmt.Printf("Result: #%v ", cluster)
	//Output:Result: &models.Cluster{Id:"53807", CustomerId:10090, Name:"enterprise-cluster", Password:"e75140b884db4e488ecfb77fca92988c", Port:34023, HazelcastVersion:"3.12.6", IsAutoScalingEnabled:false, IsHotBackupEnabled:false, IsHotRestartEnabled:false, IsIpWhitelistEnabled:false, IsTlsEnabled:false, ProductType:struct { Name models.ProductTypeName "json:\"name\""; IsFree bool "json:\"isFree\"" }{Name:"Enterprise", IsFree:false}, State:"RUNNING", CreatedAt:"2020-09-08T07:39:38.000Z", StartedAt:"2020-09-08T07:50:59.000Z", StoppedAt:"", Progress:struct { Status string "json:\"status\""; TotalItemCount int "json:\"totalItemCount\""; CompletedItemCount int "json:\"completedItemCount\"" }{Status:"Installing Hazelcast", TotalItemCount:0, CompletedItemCount:4}, CloudProvider:struct { Name string "json:\"name\""; Region string "json:\"region\""; AvailabilityZones []string "json:\"availabilityZones\"" }{Name:"aws", Region:"us-east-2", AvailabilityZones:[]string{"us-east-2c", "us-east-2a", "us-east-2b"}}, DiscoveryTokens:[]models.DiscoveryToken{models.DiscoveryToken{Source:"public", Token:"RA6aaSPU9B0PouzK9VRYAISRYBFPCmqVl9E5lU737pyCat0NWQ"}}, Specs:struct { TotalMemory float64 "json:\"totalMemory\""; HeapMemory int "json:\"heapMemory\""; NativeMemory int "json:\"nativeMemory\""; Cpu int "json:\"cpu\""; InstanceType string "json:\"instanceType\""; InstancePerZone int "json:\"instancePerZone\"" }{TotalMemory:21, HeapMemory:4, NativeMemory:3, Cpu:1500, InstanceType:"m5.large", InstancePerZone:1}, Networking:struct { Type string "json:\"type\""; CidrBlock string "json:\"cidrBlock\""; Peering struct { IsEnabled bool "json:\"is_enabled\"" } "json:\"peering\""; PrivateLink struct { Url string "json:\"url\""; State string "json:\"state\"" } "json:\"privateLink\"" }{Type:"PUBLIC", CidrBlock:"10.18.0.0/16", Peering:struct { IsEnabled bool "json:\"is_enabled\"" }{IsEnabled:false}, PrivateLink:struct { Url string "json:\"url\""; State string "json:\"state\"" }{Url:"", State:"NOT_ACTIVE"}}, DataStructures:models.DataStructureResponse{MapConfigs:[]models.MapConfigResponse{}, JCacheConfigs:[]models.JCacheConfigResponse{}, ReplicatedMapConfigs:[]models.ReplicatedMapConfigResponse{}, QueueConfigs:[]models.QueueConfigResponse{}, SetConfigs:[]models.SetConfigResponse{}, ListConfigs:[]models.ListConfigResponse{}, TopicConfigs:[]models.TopicConfigResponse{}, MultiMapConfigs:[]models.MultiMapConfigResponse{}, RingBufferConfigs:[]models.RingBufferConfigResponse{}, ReliableTopicConfigs:[]models.ReliableTopicConfigResponse{}}}
}

func ExampleEnterpriseClusterService_delete() {
	client, _, _ := New()
	cluster, _, _ := client.EnterpriseCluster.Delete(context.Background(), &models.ClusterDeleteInput{ClusterId: "53807"})
	fmt.Printf("Result: #%v ", cluster)
	//Output:Result: &models.ClusterId{ClusterId:53805}}
}
