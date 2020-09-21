package to

import (
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuery_When_Args_And_Input_Are_Not_Nil(t *testing.T) {
	//given
	name := "queryName"
	query := models.Query
	response := models.ClusterIdResponse{}
	//when
	queryString := Query(name, query, nil, nil, response)

	//then
	assert.Equal(t, "query{response:queryName{clusterId\n}}", queryString)
}

func TestQuery_When_Args_Is_Not_Nil(t *testing.T) {
	//given
	name := "queryName"
	query := models.Query
	args := models.GetStarterClusterInput{ClusterId: "123456"}
	response := models.ClusterIdResponse{}

	//when
	queryString := Query(name, query, nil, args, response)

	//then
	assert.Equal(t, "query{response:queryName(clusterId:\"123456\"){clusterId\n}}", queryString)
}

func TestQuery_When_Input_Is_Not_Nil(t *testing.T) {
	//given
	name := "queryName"
	query := models.Query
	input := models.GetStarterClusterInput{ClusterId: "123456"}
	response := models.ClusterIdResponse{}

	//when
	queryString := Query(name, query, input, nil, response)

	//then
	assert.Equal(t, "query($input:GetStarterClusterInput){response:queryName(input:$input){clusterId\n}}", queryString)
}

func TestArguments(t *testing.T) {
	//given
	test := ExampleModel{
		Name:    "name",
		Size:    10,
		Enabled: false,
	}

	//when
	arguments := Arguments(test)

	//then
	assert.Equal(t, "name:\"name\",size:10,enabled:%!s(bool=false)", arguments)
}

func TestVariables_When_Input_Is_Nil(t *testing.T) {
	//given
	//when
	variables := Variables(nil)
	//then
	assert.Empty(t, variables)
}

func TestVariables_When_Input_Is_Not_Nil(t *testing.T) {
	//given
	test := ExampleModel{
		Name:    "name",
		Size:    10,
		Enabled: false,
	}

	//when
	variables := Variables(test)

	//then
	assert.Equal(t, variables["Name"], "name")
	assert.Equal(t, variables["Size"], float64(10))
	assert.Equal(t, variables["Enabled"], false)
}

func TestVariables_When_Input_Is_Not_Unmarshall(t *testing.T) {
	//given
	input := "test"

	//when
	testFunc := assert.PanicTestFunc(func() {
		Variables(input)
	})

	//then
	panics := assert.Panics(t, testFunc)
	assert.True(t, panics)
}

func TestGraphqlResponseSelector_For_Struct(t *testing.T) {
	//given
	response := models.ClusterResponse{}

	//when
	selector := GraphqlResponseSelector(response)

	//then
	assert.Equal(t,
		"id\ncustomerId\nname\npassword\nport\nhazelcastVersion\nisAutoScalingEnabled\nisHotBackupEnabled\nisHotRestartEnabled\nisIpWhitelistEnabled\nisTlsEnabled\nproductType {\nname\nisFree\n}\nstate\ncreatedAt\nstartedAt\nstoppedAt\nprogress {\nstatus\ntotalItemCount\ncompletedItemCount\n}\ncloudProvider {\nname\nregion\navailabilityZones}\ndiscoveryTokens {\nsource\ntoken\n}\nspecs {\ntotalMemory\nheapMemory\nnativeMemory\ncpu\ninstanceType\ninstancePerZone\n}\nnetworking {\ntype\ncidrBlock\npeering {\nisEnabled\n}\nprivateLink {\nurl\nstate\n}\n}\ndataStructures {\nmapConfigs {\nid\nname\nasyncBackupCount\nbackupCount\nevictionPolicy\nisReady\nmapIndices {\nid\nname\n}\nmapStore {\nclassName\nwriteBatchSize\nwriteCoalescing\ninitialLoadMode\nwriteDelaySeconds\n}\nmaxIdleSeconds\nmaxSize\nmaxSizePolicy\nttlSeconds\n}\njCacheConfigs {\nid\nname\nasyncBackupCount\nbackupCount\nevictionPolicy\nisReady\nkeyType\nmaxSize\nmaxSizePolicy\nttlSeconds\nvalueType\n}\nreplicatedMapConfigs {\nid\nname\nisReady\nasyncFillUp\ninMemoryFormat\n}\nqueueConfigs {\nid\nname\nasyncBackupCount\nbackupCount\nemptyQueueTtl\nisReady\nmaxSize\n}\nsetConfigs {\nid\nname\nasyncBackupCount\nbackupCount\nisReady\nmaxSize\n}\nlistConfigs {\nid\nname\nasyncBackupCount\nbackupCount\nisReady\nmaxSize\n}\ntopicConfigs {\nid\nname\nglobalOrdering\nisReady\n}\nmultiMapConfigs {\nid\nname\nasyncBackupCount\nbackupCount\nisReady\nvalueCollectionType\n}\nringBufferConfigs {\nid\nname\nasyncBackupCount\nbackupCount\nisReady\ncapacity\ninMemoryFormat\nttlSeconds\n}\nreliableTopicConfigs {\nid\nname\nisReady\nreadBatchSize\ntopicOverloadPolicy\n}\n}\n",
		selector)
}

func TestGraphqlResponseSelector_For_Slice(t *testing.T) {
	//given
	var response []models.ClusterResponse

	//when
	selector := GraphqlResponseSelector(response)

	//then
	assert.Equal(t,
		"id\ncustomerId\nname\npassword\nport\nhazelcastVersion\nisAutoScalingEnabled\nisHotBackupEnabled\nisHotRestartEnabled\nisIpWhitelistEnabled\nisTlsEnabled\nproductType {\nname\nisFree\n}\nstate\ncreatedAt\nstartedAt\nstoppedAt\nprogress {\nstatus\ntotalItemCount\ncompletedItemCount\n}\ncloudProvider {\nname\nregion\navailabilityZones}\ndiscoveryTokens {\nsource\ntoken\n}\nspecs {\ntotalMemory\nheapMemory\nnativeMemory\ncpu\ninstanceType\ninstancePerZone\n}\nnetworking {\ntype\ncidrBlock\npeering {\nisEnabled\n}\nprivateLink {\nurl\nstate\n}\n}\ndataStructures {\nmapConfigs {\nid\nname\nasyncBackupCount\nbackupCount\nevictionPolicy\nisReady\nmapIndices {\nid\nname\n}\nmapStore {\nclassName\nwriteBatchSize\nwriteCoalescing\ninitialLoadMode\nwriteDelaySeconds\n}\nmaxIdleSeconds\nmaxSize\nmaxSizePolicy\nttlSeconds\n}\njCacheConfigs {\nid\nname\nasyncBackupCount\nbackupCount\nevictionPolicy\nisReady\nkeyType\nmaxSize\nmaxSizePolicy\nttlSeconds\nvalueType\n}\nreplicatedMapConfigs {\nid\nname\nisReady\nasyncFillUp\ninMemoryFormat\n}\nqueueConfigs {\nid\nname\nasyncBackupCount\nbackupCount\nemptyQueueTtl\nisReady\nmaxSize\n}\nsetConfigs {\nid\nname\nasyncBackupCount\nbackupCount\nisReady\nmaxSize\n}\nlistConfigs {\nid\nname\nasyncBackupCount\nbackupCount\nisReady\nmaxSize\n}\ntopicConfigs {\nid\nname\nglobalOrdering\nisReady\n}\nmultiMapConfigs {\nid\nname\nasyncBackupCount\nbackupCount\nisReady\nvalueCollectionType\n}\nringBufferConfigs {\nid\nname\nasyncBackupCount\nbackupCount\nisReady\ncapacity\ninMemoryFormat\nttlSeconds\n}\nreliableTopicConfigs {\nid\nname\nisReady\nreadBatchSize\ntopicOverloadPolicy\n}\n}\n",
		selector)
}

type ExampleModel struct {
	Name    string
	Size    int
	Enabled bool
}
