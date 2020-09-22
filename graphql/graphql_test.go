package graphql

import (
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuery_When_Args_And_Input_Are_Not_Nil(t *testing.T) {
	//given
	name := "queryName"
	query := models.Query
	response := models.ClusterId{}
	//when
	queryString := Query(name, query, nil, nil, response)

	//then
	assert.Equal(t, "query{response:queryName{clusterId}}", queryString)
}

func TestQuery_When_Args_Is_Not_Nil(t *testing.T) {
	//given
	name := "queryName"
	query := models.Query
	args := models.GetStarterClusterInput{ClusterId: "123456"}
	response := models.ClusterId{}

	//when
	queryString := Query(name, query, nil, args, response)

	//then
	assert.Equal(t, "query{response:queryName(clusterId:\"123456\"){clusterId}}", queryString)
}

func TestQuery_When_Input_Is_Not_Nil(t *testing.T) {
	//given
	name := "queryName"
	query := models.Query
	input := models.GetStarterClusterInput{ClusterId: "123456"}
	response := models.ClusterId{}

	//when
	queryString := Query(name, query, input, nil, response)

	//then
	assert.Equal(t, "query($input:GetStarterClusterInput){response:queryName(input:$input){clusterId}}", queryString)
}

func TestArguments(t *testing.T) {
	//given
	test := ExampleModel{
		Name:    "name",
		Size:    10,
		Enabled: false,
	}

	//when
	arguments := argumentsBuilder(test)

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
	response := models.Cluster{}

	//when
	selector := responseBuilder(response)

	//then
	assert.Equal(t,
		"{id,customerId,name,password,port,hazelcastVersion,isAutoScalingEnabled,isHotBackupEnabled,isHotRestartEnabled,isIpWhitelistEnabled,isTlsEnabled,productType{name,isFree},state,createdAt,startedAt,stoppedAt,progress{status,totalItemCount,completedItemCount},cloudProvider{name,region,availabilityZones},discoveryTokens{source,token},specs{totalMemory,heapMemory,nativeMemory,cpu,instanceType,instancePerZone},networking{type,cidrBlock,peering{isEnabled},privateLink{url,state}},dataStructures{mapConfigs{id,name,asyncBackupCount,backupCount,evictionPolicy,isReady,mapIndices{id,name},mapStore{className,writeBatchSize,writeCoalescing,initialLoadMode,writeDelaySeconds},maxIdleSeconds,maxSize,maxSizePolicy,ttlSeconds},jCacheConfigs{id,name,asyncBackupCount,backupCount,evictionPolicy,isReady,keyType,maxSize,maxSizePolicy,ttlSeconds,valueType},replicatedMapConfigs{id,name,isReady,asyncFillUp,inMemoryFormat},queueConfigs{id,name,asyncBackupCount,backupCount,emptyQueueTtl,isReady,maxSize},setConfigs{id,name,asyncBackupCount,backupCount,isReady,maxSize},listConfigs{id,name,asyncBackupCount,backupCount,isReady,maxSize},topicConfigs{id,name,globalOrdering,isReady},multiMapConfigs{id,name,asyncBackupCount,backupCount,isReady,valueCollectionType},ringBufferConfigs{id,name,asyncBackupCount,backupCount,isReady,capacity,inMemoryFormat,ttlSeconds},reliableTopicConfigs{id,name,isReady,readBatchSize,topicOverloadPolicy}}}",
		selector)
}

func TestGraphqlResponseSelector_For_Slice(t *testing.T) {
	//given
	var response []models.Cluster

	//when
	selector := responseBuilder(response)

	//then
	assert.Equal(t,
		"{id,customerId,name,password,port,hazelcastVersion,isAutoScalingEnabled,isHotBackupEnabled,isHotRestartEnabled,isIpWhitelistEnabled,isTlsEnabled,productType{name,isFree},state,createdAt,startedAt,stoppedAt,progress{status,totalItemCount,completedItemCount},cloudProvider{name,region,availabilityZones},discoveryTokens{source,token},specs{totalMemory,heapMemory,nativeMemory,cpu,instanceType,instancePerZone},networking{type,cidrBlock,peering{isEnabled},privateLink{url,state}},dataStructures{mapConfigs{id,name,asyncBackupCount,backupCount,evictionPolicy,isReady,mapIndices{id,name},mapStore{className,writeBatchSize,writeCoalescing,initialLoadMode,writeDelaySeconds},maxIdleSeconds,maxSize,maxSizePolicy,ttlSeconds},jCacheConfigs{id,name,asyncBackupCount,backupCount,evictionPolicy,isReady,keyType,maxSize,maxSizePolicy,ttlSeconds,valueType},replicatedMapConfigs{id,name,isReady,asyncFillUp,inMemoryFormat},queueConfigs{id,name,asyncBackupCount,backupCount,emptyQueueTtl,isReady,maxSize},setConfigs{id,name,asyncBackupCount,backupCount,isReady,maxSize},listConfigs{id,name,asyncBackupCount,backupCount,isReady,maxSize},topicConfigs{id,name,globalOrdering,isReady},multiMapConfigs{id,name,asyncBackupCount,backupCount,isReady,valueCollectionType},ringBufferConfigs{id,name,asyncBackupCount,backupCount,isReady,capacity,inMemoryFormat,ttlSeconds},reliableTopicConfigs{id,name,isReady,readBatchSize,topicOverloadPolicy}}}",
		selector)
}

type ExampleModel struct {
	Name    string
	Size    int
	Enabled bool
}
