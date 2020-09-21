package models

//Type of cluster types for starter cluster.
type StarterClusterType string

const (
	// This type represents Free Cluster.
	Free StarterClusterType = "FREE"
	// This type represents Small Cluster.
	Small StarterClusterType = "SMALL"
	// This type represents Medium Cluster.
	Medium StarterClusterType = "MEDIUM"
	// This type represents Large Cluster.
	Large StarterClusterType = "LARGE"
)

//Product type Hazelcast Cloud clusters
type ProductTypeName string

const (
	Starter    ProductTypeName = "STARTER"
	Enterprise ProductTypeName = "ENTERPRISE"
)

//Hazelcast version for starter versions of Hazelcast Cloud clusters
type StarterHazelcastVersion string

const (
	Version312 StarterHazelcastVersion = "VERSION_3_12"
	Version40  StarterHazelcastVersion = "VERSION_4_0"
)

//Eviction policy to be applied when the size of map grows larger than the value specified by the Max Size element described below. For more information, see [Eviction Policy](https://docs-dedicated.cloud.hazelcast.com/docs/map-configurations#eviction-policy)
type EvictionPolicy string

const (
	Lru    EvictionPolicy = "LRU"
	Lfu    EvictionPolicy = "LFU"
	None   EvictionPolicy = "NONE"
	Random EvictionPolicy = "RANDOM"
)

//Internal storage format. For more information, see [In Memory Format](https://docs-dedicated.cloud.hazelcast.com/docs/replicated-map#in-memory-format)
type InMemoryFormat string

const (
	Object InMemoryFormat = "OBJECT"
	Binary InMemoryFormat = "BINARY"
	Native InMemoryFormat = "NATIVE"
)

//Type of the value collection. For more information, see [Value Collection Type]( https://docs-dedicated.cloud.hazelcast.com/docs/multimap#value-collection-type)
type ValueCollectionType string

const (
	Set  ValueCollectionType = "SET"
	List ValueCollectionType = "LIST"
)

//Sets the initial load mode. For more information, see [MapLoader and MapStore](https://docs-dedicated.cloud.hazelcast.com/v1.1/docs/maploader-and-mapstore)
type InitialLoadMode string

const (
	Lazy  InitialLoadMode = "LAZY"
	Eager InitialLoadMode = "EAGER"
)

//Policy on how to define the map's maximum size. For more information, see [Maz Size Policy](https://docs-dedicated.cloud.hazelcast.com/docs/map-configurations#max-size-policy)
type MaxSizePolicy string

const (
	UsedNativeMemoryPercentage MaxSizePolicy = "USED_NATIVE_MEMORY_PERCENTAGE"
	FreeNativeMemoryPercentage MaxSizePolicy = "FREE_NATIVE_MEMORY_PERCENTAGE"
)

//A policy to deal with an overloaded topic; so topic where there is no place to store new messages. For more information, see [Overload Policy]( https://docs-dedicated.cloud.hazelcast.com/docs/reliable-topic#reliable-topic-overload-policy
type TopicOverloadPolicy string

const (
	DiscardOldest TopicOverloadPolicy = "DISCARD_OLDEST"
	DiscardNewest TopicOverloadPolicy = "DISCARD_NEWEST"
	Block         TopicOverloadPolicy = "BLOCK"
	Error         TopicOverloadPolicy = "ERROR"
)

//State of the cluster that show current status.
type State string

const (
	Running                  State = "RUNNING"
	Pending                  State = "PENDING"
	StopInProgress           State = "STOP_IN_PROGRESS"
	Stopped                  State = "STOPPED"
	DeleteInProgress         State = "DELETE_IN_PROGRESS"
	Deleted                  State = "DELETED"
	ScaleUpInProgress        State = "SCALE_UP_IN_PROGRESS"
	ScaleDownInProgress      State = "SCALE_DOWN_IN_PROGRESS"
	ScaleUp                  State = "SCALE_UP"
	ScaleDown                State = "SCALE_DOWN"
	ResumeInProgress         State = "RESUME_IN_PROGRESS"
	UpdatingClusterConfig    State = "UPDATING_CLUSTER_CONFIG"
	UpdatingInstancePerZone  State = "UPDATING_INSTANCE_PER_ZONE"
	UpdatingZones            State = "UPDATING_ZONES"
	UpdatingHazelcastVersion State = "UPDATING_HAZELCAST_VERSION"
	UpdatingInstanceType     State = "UPDATING_INSTANCE_TYPE"
	Failed                   State = "FAILED"
)

//The Hazelcast versions for Enterprise Hazelcast Version.
type EnterpriseHazelcastVersion struct {
	Version             string   `json:"version"`
	UpgradeableVersions []string `json:"upgradeableVersions"`
}

//The input for Create Starter Cluster.
type CreateStarterClusterInput struct {
	//Name of the cluster.
	Name string `json:"name"`
	//Cloud provider of the cluster.
	CloudProvider string `json:"cloudProvider"`
	//Name of the region.
	Region string `json:"region"`
	//Cluster type of the cluster
	ClusterType StarterClusterType `json:"clusterType"`
	//Hazelcast IMDG version of the cluster.
	HazelcastVersion StarterHazelcastVersion `json:"hazelcastVersion"`
	//Total memory of the cluster.
	TotalMemory float64 `json:"totalMemory"`
	//Shows if auto scaling feature enabled or not.
	IsAutoScalingEnabled bool `json:"isAutoScalingEnabled"`
	//Shows if hot backup feature is enabled or not.
	IsHotBackupEnabled bool `json:"isHotBackupEnabled"`
	//Shows if hot restart feature is enabled or not.
	IsHotRestartEnabled bool `json:"isHotRestartEnabled"`
	//Shows if Ip whitelisting is enabled to restrict connection to the cluster or not.
	IsIPWhitelistEnabled bool `json:"isIpWhitelistEnabled"`
	//Shows if TLS(Transport Layer Security) is enabled while connecting to the cluster or not.
	IsTLSEnabled bool `json:"isTlsEnabled"`
	//Shows if Ip whitelisting is enabled to restrict connection to the cluster or not.
	IPWhitelist []string `json:"ipWhitelist"`
	//Data structure configuration of the cluster.
	DataStructure DataStructureRequest `json:"dataStructures"`
}

//The input for Get Cluster.
type GetStarterClusterInput struct {
	ClusterId string `json:"clusterId"`
}

//The input for Create Enterprise Cluster.
type CreateEnterpriseClusterInput struct {
	//Name of the cluster.
	Name string `json:"name"`
	//Cloud provider of the cluster.
	CloudProvider string `json:"cloudProvider"`
	//Name of the region.
	Region string `json:"region"`
	//Availability zones of the region.
	Zones []string `json:"zones"`
	//Instance type of the cluster
	InstanceType string `json:"instanceType"`
	//Number of Hazelcast cluster members per zone.
	InstancePerZone int `json:"instancePerZone"`
	//Hazelcast IMDG version of the cluster.
	HazelcastVersion string `json:"hazelcastVersion"`
	//Public access enabled flag of the cluster
	IsPublicAccessEnabled bool `json:"isPublicAccessEnabled"`
	//CIDR Range of the cluster network.
	CidrBlock string `json:"cidrBlock"`
	//Native memory of the cluster.
	NativeMemory int `json:"nativeMemory"`
	//Shows if auto scaling feature enabled or not.
	IsAutoScalingEnabled bool `json:"isAutoScalingEnabled"`
	//Shows if hot restart feature is enabled or not.
	IsHotRestartEnabled bool `json:"isHotRestartEnabled"`
	//Shows if hot backup feature is enabled or not.
	IsHotBackupEnabled bool `json:"isHotBackupEnabled"`
	//Shows if TLS(Transport Layer Security) is enabled while connecting to the cluster or not.
	IsTLSEnabled bool `json:"isTlsEnabled"`
	//Data structure configuration of the cluster.
	DataStructure DataStructureRequest `json:"dataStructures"`
}

//The input for Get Enterprise Cluster Input.
type GetEnterpriseClusterInput struct {
	ClusterId string `json:"clusterId"`
}

//This response id of the Cluster Response
type ClusterIdResponse struct {
	ClusterId int `json:"ClusterId"`
}

//This response for Cluster
type ClusterResponse struct {
	//Unique identifier of cluster.
	Id string `json:"id"`
	//Unique identifier of the cluster owner.
	CustomerId int `json:"customerId"`
	//Name of the cluster.
	Name string `json:"name"`
	//Password of the cluster.
	Password string `json:"password"`
	//Port of the Hazelcast IMDG on the cluster.
	Port int `json:"port"`
	//Hazelcast IMDG version of the cluster.
	HazelcastVersion string `json:"hazelcastVersion"`
	//Shows if auto scaling feature enabled or not.
	IsAutoScalingEnabled bool `json:"isAutoScalingEnabled"`
	//Shows if hot backup feature is enabled or not.
	IsHotBackupEnabled bool `json:"isHotBackupEnabled"`
	//Shows if hot restart feature is enabled or not.
	IsHotRestartEnabled bool `json:"isHotRestartEnabled"`
	//Shows if Ip whitelisting is enabled to restrict connection to the cluster or not.
	IsIpWhitelistEnabled bool `json:"isIpWhitelistEnabled"`
	//Shows if TLS(Transport Layer Security) is enabled while connecting to the cluster or not.
	IsTlsEnabled bool `json:"isTlsEnabled"`
	//Type of the product for the cluster.
	ProductType struct {
		//Name of the Product Type.
		Name ProductTypeName `json:"name"`
		//This field shows that if the product is free to use.
		IsFree bool `json:"isFree"`
	} `json:"productType"`
	//State of the cluster that show current status.
	State State `json:"state"`
	//Date that shows when cluster was created.
	CreatedAt string `json:"createdAt"`
	//Date that shows when cluster was started after the creation.
	StartedAt string `json:"startedAt"`
	//Date that shows when cluster was stopped.
	StoppedAt string `json:"stoppedAt"`
	//Progress of the cluster.
	Progress struct {
		//Status of progress.
		Status string `json:"status"`
		//Total item count of progress.
		TotalItemCount int `json:"totalItemCount"`
		//Completed item count of progress.
		CompletedItemCount int `json:"completedItemCount"`
	}
	//Cloud provider of the cluster.
	CloudProvider struct {
		//Name of the cloud provider.
		Name string `json:"name"`
		//Name of the region.
		Region string `json:"region"`
		//Availability zones of the region.
		AvailabilityZones []string `json:"availabilityZones"`
	} `json:"cloudProvider"`
	//Discovery tokens of the cluster to connect Hazelcast.
	DiscoveryTokens []DiscoveryToken `json:"discoveryTokens"`
	//Specs of the cluster.
	Specs struct {
		//Total memory of the cluster.
		TotalMemory float64 `json:"totalMemory"`
		//Heap memory of the cluster.
		HeapMemory int `json:"heapMemory"`
		//Native memory of the cluster.
		NativeMemory int `json:"nativeMemory"`
		//Number of cpu of the cluster.
		Cpu int `json:"cpu"`
		//Instance type of the cluster.
		InstanceType string `json:"instanceType"`
		//Number of Hazelcast cluster members per zone.
		InstancePerZone int `json:"instancePerZone"`
	} `json:"specs"`
	//Networking configuration of the cluster.
	Networking struct {
		//Type of the network. It determines that cluster is publicly accessible or not.
		Type string `json:"type"`
		//CIDR Range of the cluster network.
		CidrBlock string `json:"cidrBlock"`
		//Peering information of the cluster networking.
		Peering struct {
			//This field is true if at least one peering connection established, false otherwise.
			IsEnabled bool `json:"is_enabled"`
		} `json:"peering"`
		//Peering information of the cluster networking.
		PrivateLink struct {
			//State of the private link connection.
			Url string `json:"url"`
			//Url to cloud formation template that you can use for Private Link. For more information, see [AWS Private Link](https
			State string `json:"state"`
		} `json:"privateLink"`
	} `json:"networking"`
	//Data structure configuration of the cluster.
	DataStructures DataStructureResponse `json:"dataStructures"`
}

//Discovery token of the cluster to connect Hazelcast.
type DiscoveryToken struct {
	Source string `json:"source"`
	Token  string `json:"token"`
}

//Type of all data structure request
type DataStructureRequest struct {
	//List of map config for the cluster. For more information, see [Map Configuration](https://docs-dedicated.cloud.hazelcast.com/docs/map-configurations)
	MapConfigs []MapConfigRequest `json:"mapConfigs,omitempty"`
	//List of JCache config for the cluster. For more information, see [JCache](https://docs-dedicated.cloud.hazelcast.com/docs/jcache)
	JcacheConfigs []JCacheConfigRequest `json:"jcacheConfigs,omitempty"`
	//List of Replicated Map config for the cluster. For more information, see [Replicated Map](https://docs-dedicated.cloud.hazelcast.com/docs/replicated-map)
	ReplicatedMapConfigs []ReplicatedMapConfigRequest `json:"replicatedMapConfigs,omitempty"`
	//List of Queue config for the cluster. For more information, see [Queue](https://docs-dedicated.cloud.hazelcast.com/docs/queue)
	QueueConfigs []QueueConfigRequest `json:"queueConfigs,omitempty"`
	//List of Set config for the cluster. For more information, see [Set](https://docs-dedicated.cloud.hazelcast.com/docs/set)
	SetConfigs []SetConfigRequest `json:"setConfigs,omitempty"`
	//List of List config for the cluster. For more information, see [List Configuration](https://docs-dedicated.cloud.hazelcast.com/docs/list)
	ListConfigs []ListConfigRequest `json:"listConfigs,omitempty"`
	//List of Topic config for the cluster. For more information, see [Topic](https://docs-dedicated.cloud.hazelcast.com/docs/topic)
	TopicConfigs []TopicConfigRequest `json:"topicConfigs,omitempty"`
	//List of Multi Map config for the cluster. For more information, see [MultiMap](https://docs-dedicated.cloud.hazelcast.com/docs/multimap)
	MultiMapConfigs []MultiMapConfigRequest `json:"multiMapConfigs,omitempty"`
	//List of Ring Buffer config for the cluster. For more information, see [RingBuffer](https://docs-dedicated.cloud.hazelcast.com/docs/ringbuffer)
	RingBufferConfigs []RingBufferConfigRequest `json:"ringBufferConfigs,omitempty"`
	//List of Reliable Topic config for the cluster. For more information, see [ReliableTopic](https://docs-dedicated.cloud.hazelcast.com/docs/reliable-topic)
	ReliableTopicConfigs []ReliableTopicConfigRequest `json:"reliableTopicConfigs,omitempty"`
}

//Type of all data structure response
type DataStructureResponse struct {
	//List of map config for the cluster. For more information, see [Map Configuration](https://docs-dedicated.cloud.hazelcast.com/docs/map-configurations)
	MapConfigs []MapConfigResponse `json:"mapConfigs"`
	//List of JCache config for the cluster. For more information, see [JCache](https://docs-dedicated.cloud.hazelcast.com/docs/jcache)
	JCacheConfigs []JCacheConfigResponse `json:"jCacheConfigs"`
	//List of Replicated Map config for the cluster. For more information, see [Replicated Map](https://docs-dedicated.cloud.hazelcast.com/docs/replicated-map)
	ReplicatedMapConfigs []ReplicatedMapConfigResponse `json:"replicatedMapConfigs"`
	//List of Queue config for the cluster. For more information, see [Queue](https://docs-dedicated.cloud.hazelcast.com/docs/queue)
	QueueConfigs []QueueConfigResponse `json:"queueConfigs"`
	//List of Set config for the cluster. For more information, see [Set](https://docs-dedicated.cloud.hazelcast.com/docs/set)
	SetConfigs []SetConfigResponse `json:"setConfigs"`
	//List of List config for the cluster. For more information, see [List Configuration](https://docs-dedicated.cloud.hazelcast.com/docs/list)
	ListConfigs []ListConfigResponse `json:"listConfigs"`
	//List of Topic config for the cluster. For more information, see [Topic](https://docs-dedicated.cloud.hazelcast.com/docs/topic)
	TopicConfigs []TopicConfigResponse `json:"topicConfigs"`
	//List of Multi Map config for the cluster. For more information, see [MultiMap](https://docs-dedicated.cloud.hazelcast.com/docs/multimap)
	MultiMapConfigs []MultiMapConfigResponse `json:"multiMapConfigs"`
	//List of Ring Buffer config for the cluster. For more information, see [RingBuffer](https://docs-dedicated.cloud.hazelcast.com/docs/ringbuffer)
	RingBufferConfigs []RingBufferConfigResponse `json:"ringBufferConfigs"`
	//List of Reliable Topic config for the cluster. For more information, see [ReliableTopic](https://docs-dedicated.cloud.hazelcast.com/docs/reliable-topic)
	ReliableTopicConfigs []ReliableTopicConfigResponse `json:"reliableTopicConfigs"`
}

//Type of Map Data Structures Request
type MapConfigRequest struct {
	//Name of the map config. For more information, see [Map Name]( https://docs-dedicated.cloud.hazelcast.com/docs/map-configurations#map-name)
	Name string `json:"name"`
	//Number of asynchronous backups/replications. For more information, see [Asynchronous Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/map-configurations#asynchronous-backup-count)
	AsyncBackupCount int `json:"asyncBackupCount"`
	//Number of backups/replications. For more information, see [Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/map-configurations#backup-count)
	BackupCount int `json:"backupCount"`
	//Eviction policy to be applied when the size of map grows larger than the value specified by the Max Size element described below. For more information, see [Eviction Policy](https://docs-dedicated.cloud.hazelcast.com/docs/map-configurations#eviction-policy)
	EvictionPolicy EvictionPolicy `json:"evictionPolicy"`
	//Hazelcast iterates through all the owned entries and finds the matching ones. This can be made faster by indexing the most frequently queried fields.
	MapIndices *[]string `json:"indexes,omitempty"`
	//Unique Identifier of the map store.
	MapStore *MapStoreRequest `json:"mapStore,omitempty"`
	//Maximum time in seconds for each entry to stay idle in the map. For more information, see [Title](https://docs-dedicated.cloud.hazelcast.com/docs/map-configurations#max-idle-seconds)
	MaxIdleSeconds int `json:"maxIdleSeconds"`
	//The percentage value when the eviction will start depending on the selected Max Size Policy. For more information, see [Max Size](https://docs-dedicated.cloud.hazelcast.com/docs/map-configurations#max-size)
	MaxSize int `json:"maxSize"`
	//Policy on how to define the map's maximum size. For more information, see [Maz Size Policy](https://docs-dedicated.cloud.hazelcast.com/docs/map-configurations#max-size-policy)
	MaxSizePolicy MaxSizePolicy `json:"maxSizePolicy"`
	//It is the maximum time in seconds for each entry to stay in the map. For more information, see [TTL](https://docs-dedicated.cloud.hazelcast.com/docs/map-configurations#ttl)
	TtlSeconds int `json:"ttlSeconds"`
}

//Type of Map Store Data Structures Request
type MapStoreRequest struct {
	//Name of the class implementing MapLoader and/or MapStore. For more information, see [MapLoader and MapStore](https://docs-dedicated.cloud.hazelcast.com/v1.1/docs/maploader-and-mapstore)
	ClassName string `json:"className"`
	//Used to create batch chunks when writing map store. For more information, see [MapLoader and MapStore](https://docs-dedicated.cloud.hazelcast.com/v1.1/docs/maploader-and-mapstore)
	WriteBatchSize int `json:"writeBatchSize"`
	//In write-behind mode, Hazelcast coalesces updates on a specific key by default; it applies only the last update on it. For more information, see [MapLoader and MapStore](https://docs-dedicated.cloud.hazelcast.com/v1.1/docs/maploader-and-mapstore)
	WriteCoalescing bool `json:"writeCoalescing"`
	//Sets the initial load mode. For more information, see [MapLoader and MapStore](https://docs-dedicated.cloud.hazelcast.com/v1.1/docs/maploader-and-mapstore)
	InitialLoadMode InitialLoadMode `json:"initialLoadMode"`
	//Number of seconds to delay to call the MapStore.store(key, value). For more information, see [MapLoader and MapStore](https://docs-dedicated.cloud.hazelcast.com/v1.1/docs/maploader-and-mapstore)
	WriteDelaySeconds int                `json:"writeDelaySeconds"`
	Properties        []MapStoreProperty `json:"properties"`
}

//Type of Map Store Property Data Structures Request
type MapStoreProperty struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//Type of Map Config Response
type MapConfigResponse struct {
	//Unique Identifier of the map config.
	Id string `json:"id"`
	//Name of the map config. For more information, see [Map Name]( https://docs-dedicated.cloud.hazelcast.com/docs/map-configurations#map-name)
	Name string `json:"name"`
	//Number of asynchronous backups/replications. For more information, see [Asynchronous Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/map-configurations#asynchronous-backup-count)
	AsyncBackupCount int `json:"asyncBackupCount"`
	//Number of backups/replications. For more information, see [Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/map-configurations#backup-count)
	BackupCount int `json:"backupCount"`
	//Eviction policy to be applied when the size of map grows larger than the value specified by the Max Size element described below. For more information, see [Eviction Policy](https://docs-dedicated.cloud.hazelcast.com/docs/map-configurations#eviction-policy)
	EvictionPolicy string `json:"evictionPolicy"`
	//Shows that if this config is ready to use.
	IsReady bool `json:"isReady"`
	//Hazelcast iterates through all the owned entries and finds the matching ones. This can be made faster by indexing the most frequently queried fields.
	MapIndices []struct {
		//Unique Identifier of the map index.
		Id string `json:"id"`
		//Name of the map index.
		Name string `json:"name"`
	} `json:"mapIndices"`
	//Unique Identifier of the map store.
	MapStore struct {
		//Name of the class implementing MapLoader and/or MapStore. For more information, see [MapLoader and MapStore](https://docs-dedicated.cloud.hazelcast.com/v1.1/docs/maploader-and-mapstore)
		ClassName string `json:"class_name"`
		//Used to create batch chunks when writing map store. For more information, see [MapLoader and MapStore](https://docs-dedicated.cloud.hazelcast.com/v1.1/docs/maploader-and-mapstore)
		WriteBatchSize int `json:"write_batch_size"`
		//In write-behind mode, Hazelcast coalesces updates on a specific key by default; it applies only the last update on it. For more information, see [MapLoader and MapStore](https://docs-dedicated.cloud.hazelcast.com/v1.1/docs/maploader-and-mapstore)
		WriteCoalescing bool `json:"write_coalescing"`
		//Sets the initial load mode. For more information, see [MapLoader and MapStore](https://docs-dedicated.cloud.hazelcast.com/v1.1/docs/maploader-and-mapstore)
		InitialLoadMode string `json:"initial_load_mode"`
		//Number of seconds to delay to call the MapStore.store(key, value). For more information, see [MapLoader and MapStore](https://docs-dedicated.cloud.hazelcast.com/v1.1/docs/maploader-and-mapstore)
		WriteDelaySeconds int `json:"write_delay_seconds"`
	} `json:"mapStore"`
	//Maximum time in seconds for each entry to stay idle in the map. For more information, see [Title](https://docs-dedicated.cloud.hazelcast.com/docs/map-configurations#max-idle-seconds)
	MaxIdleSeconds int `json:"maxIdleSeconds"`
	//The percentage value when the eviction will start depending on the selected Max Size Policy. For more information, see [Max Size](https://docs-dedicated.cloud.hazelcast.com/docs/map-configurations#max-size)
	MaxSize int `json:"maxSize"`
	//Policy on how to define the map's maximum size. For more information, see [Maz Size Policy](https://docs-dedicated.cloud.hazelcast.com/docs/map-configurations#max-size-policy)
	MaxSizePolicy string `json:"maxSizePolicy"`
	//It is the maximum time in seconds for each entry to stay in the map. For more information, see [TTL](https://docs-dedicated.cloud.hazelcast.com/docs/map-configurations#ttl)
	TtlSeconds int `json:"ttlSeconds"`
}

//Type of List Data Structure Request
type ListConfigRequest struct {
	//Name of the list config.
	Name string `json:"name"`
	//This value sets how many asynchronous backup will be. For more information, see [Asynchronous Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/list#asynchronous-backup-count)
	AsyncBackupCount int `json:"asyncBackupCount"`
	//This value sets that how many synchronous backup will be. For more information, see [Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/list#backup-count)
	BackupCount int `json:"backupCount"`
	//You can set maximum number of entries for the List you are creating with setting this value. For more information, see [Maximum Size](https://docs-dedicated.cloud.hazelcast.com/docs/list#maximum-size)
	MaxSize int `json:"maxSize"`
}

//Type of List Data Structure Response
type ListConfigResponse struct {
	//Unique Identifier of the list config.
	Id string `json:"id"`
	//Name of the list config.
	Name string `json:"name"`
	//This value sets how many asynchronous backup will be. For more information, see [Asynchronous Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/list#asynchronous-backup-count)
	AsyncBackupCount int `json:"asyncBackupCount"`
	//This value sets that how many synchronous backup will be. For more information, see [Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/list#backup-count)
	BackupCount int `json:"backupCount"`
	//Shows that if this config is ready to use.
	IsReady bool `json:"isReady"`
	//You can set maximum number of entries for the List you are creating with setting this value. For more information, see [Maximum Size](https://docs-dedicated.cloud.hazelcast.com/docs/list#maximum-size)
	MaxSize int `json:"maxSize"`
}

//Type of Set Data Structure Request
type SetConfigRequest struct {
	//Name of the Set config.
	Name string `json:"name"`
	//This value sets how many asynchronous backup will be. For more information, see [Asynchronous Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/set#asynchronous-backup-count)
	AsyncBackupCount int `json:"asyncBackupCount"`
	//This value sets that how many synchronous backup will be. For more information, see [Backup Count]( https
	BackupCount int `json:"backupCount"`
	//You can set maximum number of entries for the set you are creating with setting this value. For more information, see [Maximum Size]( https://docs-dedicated.cloud.hazelcast.com/docs/set#maximum-size)
	MaxSize int `json:"maxSize"`
}

//Type of Set Data Structure Response
type SetConfigResponse struct {
	//Unique Identifier of the Set config.
	Id string `json:"id"`
	//Name of the Set config.
	Name string `json:"name"`
	//This value sets how many asynchronous backup will be. For more information, see [Asynchronous Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/set#asynchronous-backup-count)
	AsyncBackupCount int `json:"asyncBackupCount"`
	//This value sets that how many synchronous backup will be. For more information, see [Backup Count]( https
	BackupCount int `json:"backupCount"`
	//Shows that if this config is ready to use.
	IsReady bool `json:"isReady"`
	//You can set maximum number of entries for the set you are creating with setting this value. For more information, see [Maximum Size]( https://docs-dedicated.cloud.hazelcast.com/docs/set#maximum-size)
	MaxSize int `json:"maxSize"`
}

//Type of Queue Data Structure Request
type QueueConfigRequest struct {
	//Name of the Queue config.
	Name string `json:"name"`
	//Number of asynchronous backups. For more information, see [Asynchronous Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/queue#asynchronous-backup-count)
	AsyncBackupCount int `json:"asyncBackupCount"`
	//Number of synchronous backups. For more information, see [Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/queue#backup-count)
	BackupCount int `json:"backupCount"`
	//Used to purge unused or empty queues. For more information, see [TTL](https://docs-dedicated.cloud.hazelcast.com/docs/queue#time-to-live)
	EmptyQueueTtl int `json:"emptyQueueTtl"`
	//Maximum number of items in the queue. For more information, see [Maximum Size](https://docs-dedicated.cloud.hazelcast.com/docs/queue#maximum-size)
	MaxSize int `json:"maxSize"`
}

//Type of Queue Data Structure Response
type QueueConfigResponse struct {
	//Unique Identifier of the Queue config.
	Id string `json:"id"`
	//Name of the Queue config.
	Name string `json:"name"`
	//Number of asynchronous backups. For more information, see [Asynchronous Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/queue#asynchronous-backup-count)
	AsyncBackupCount int `json:"asyncBackupCount"`
	//Number of synchronous backups. For more information, see [Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/queue#backup-count)
	BackupCount int `json:"backupCount"`
	//Used to purge unused or empty queues. For more information, see [TTL](https://docs-dedicated.cloud.hazelcast.com/docs/queue#time-to-live)
	EmptyQueueTtl int `json:"emptyQueueTtl"`
	//Shows that if this config is ready to use.
	IsReady bool `json:"isReady"`
	//Maximum number of items in the queue. For more information, see [Maximum Size](https://docs-dedicated.cloud.hazelcast.com/docs/queue#maximum-size)
	MaxSize int `json:"maxSize"`
}

//Type of Replicated Map Data Structure Request
type ReplicatedMapConfigRequest struct {
	//Name of the Replicated Map config.
	Name string `json:"name"`
	//Specifies whether the Replicated Map is available for reads before the initial replication is completed. For more information, see [Asynchronous Fillup](https://docs-dedicated.cloud.hazelcast.com/docs/replicated-map#asynchronous-fillup)
	AsyncFillUp bool `json:"asyncFillUp"`
	//Internal storage format. For more information, see [In Memory Format](https://docs-dedicated.cloud.hazelcast.com/docs/replicated-map#in-memory-format)
	InMemoryFormat InMemoryFormat `json:"inMemoryFormat"`
}

//Type of Replicated Map Data Structure Response
type ReplicatedMapConfigResponse struct {
	//Unique Identifier of the Replicated Map config.
	Id string `json:"id"`
	//Name of the Replicated Map config.
	Name string `json:"name"`
	//Shows that if this config is ready to use.
	IsReady bool `json:"isReady"`
	//Specifies whether the Replicated Map is available for reads before the initial replication is completed. For more information, see [Asynchronous Fillup](https://docs-dedicated.cloud.hazelcast.com/docs/replicated-map#asynchronous-fillup)
	AsyncFillUp bool `json:"asyncFillUp"`
	//Internal storage format. For more information, see [In Memory Format](https://docs-dedicated.cloud.hazelcast.com/docs/replicated-map#in-memory-format)
	InMemoryFormat string `json:"inMemoryFormat"`
}

//Type of JCache Data Structure Request
type JCacheConfigRequest struct {
	//Name of the JCache config.
	Name string `json:"name"`
	//Number of asynchronous backups/replications. For more information, see [Asynchronous Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/jcache#asynchronous-backup-count)
	AsyncBackupCount int `json:"asyncBackupCount"`
	//Number of synchronous backups/replications. For more information, see [Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/jcache#backup-count)
	BackupCount int `json:"backupCount"`
	//Eviction policy that compares values to find the best matching eviction candidate. For more information, see [Eviction Policy](https://docs-dedicated.cloud.hazelcast.com/docs/jcache#eviction-policy)
	EvictionPolicy EvictionPolicy `json:"evictionPolicy"`
	//Fully qualified class name of the cache key type. For more information, see [Key Type](https://docs-dedicated.cloud.hazelcast.com/docs/jcache#key-type)
	KeyType string `json:"keyType"`
	//Maximum free/used percentage or maximum size in bytes depending on the selected maxSizePolicy. For more information, see [Maximum Size]( https://docs-dedicated.cloud.hazelcast.com/docs/jcache#maximum-size)
	MaxSize int `json:"maxSize"`
	//Maximum size policy. If maximum size is reached, the cache is evicted based on the eviction policy. For more information, see [Max Size Policy](https://docs-dedicated.cloud.hazelcast.com/docs/jcache#max-size-policy)
	MaxSizePolicy MaxSizePolicy `json:"maxSizePolicy"`
	//Time-to-live for cache entries, in seconds. For more information, see [TTL](https://docs-dedicated.cloud.hazelcast.com/docs/jcache#time-to-live)
	TtlSeconds int `json:"ttlSeconds"`
	//Fully qualified class name of the cache value type. For more information, see [Value Type](https://docs-dedicated.cloud.hazelcast.com/docs/jcache#value-type)
	ValueType string `json:"valueType"`
}

//Type of JCache Data Structure Response
type JCacheConfigResponse struct {
	//Unique Identifier of the JCache config.
	Id string `json:"id"`
	//Name of the JCache config.
	Name string `json:"name"`
	//Number of asynchronous backups/replications. For more information, see [Asynchronous Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/jcache#asynchronous-backup-count)
	AsyncBackupCount int `json:"asyncBackupCount"`
	//Number of synchronous backups/replications. For more information, see [Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/jcache#backup-count)
	BackupCount int `json:"backupCount"`
	//Eviction policy that compares values to find the best matching eviction candidate. For more information, see [Eviction Policy](https://docs-dedicated.cloud.hazelcast.com/docs/jcache#eviction-policy)
	EvictionPolicy string `json:"evictionPolicy"`
	//Shows that if this config is ready to use.
	IsReady bool `json:"isReady"`
	//Fully qualified class name of the cache key type. For more information, see [Key Type](https://docs-dedicated.cloud.hazelcast.com/docs/jcache#key-type)
	KeyType string `json:"keyType"`
	//Maximum free/used percentage or maximum size in bytes depending on the selected maxSizePolicy. For more information, see [Maximum Size]( https://docs-dedicated.cloud.hazelcast.com/docs/jcache#maximum-size)
	MaxSize int `json:"maxSize"`
	//Maximum size policy. If maximum size is reached, the cache is evicted based on the eviction policy. For more information, see [Max Size Policy](https://docs-dedicated.cloud.hazelcast.com/docs/jcache#max-size-policy)
	MaxSizePolicy string `json:"maxSizePolicy"`
	//Time-to-live for cache entries, in seconds. For more information, see [TTL](https://docs-dedicated.cloud.hazelcast.com/docs/jcache#time-to-live)
	TtlSeconds int `json:"ttlSeconds"`
	//Fully qualified class name of the cache value type. For more information, see [Value Type](https://docs-dedicated.cloud.hazelcast.com/docs/jcache#value-type)
	ValueType string `json:"valueType"`
}

//Type of Topic Data Structure Request
type TopicConfigRequest struct {
	//Name of the Topic config.
	Name string `json:"name"`
	//Global ordering enables, all receivers will receive all messages from all sources with the same order. For more information, see [Global Ordering]( https://docs-dedicated.cloud.hazelcast.com/docs/topic#global-ordering)
	GlobalOrdering bool `json:"globalOrdering"`
}

//Type of Topic Data Structure Response
type TopicConfigResponse struct {
	//Unique Identifier of the Topic config.
	Id string `json:"id"`
	//Name of the Topic config.
	Name string `json:"name"`
	//Global ordering enables, all receivers will receive all messages from all sources with the same order. For more information, see [Global Ordering]( https://docs-dedicated.cloud.hazelcast.com/docs/topic#global-ordering)
	GlobalOrdering bool `json:"globalOrdering"`
	//Shows that if this config is ready to use.
	IsReady bool `json:"isReady"`
}

//Type of MultiMap Data Structure Request
type MultiMapConfigRequest struct {
	//Name of the MultiMap config.
	Name string `json:"name"`
	//The number of asynchronous backups. For more information, see [Asynchronous Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/multimap#asynchronous-backups-count)
	AsyncBackupCount int `json:"asyncBackupCount"`
	//The number of synchronous backups. For more information, see [Backup Count]( https://docs-dedicated.cloud.hazelcast.com/docs/multimap#backup-count)
	BackupCount int `json:"backupCount"`
	//Type of the value collection. For more information, see [Value Collection Type]( https://docs-dedicated.cloud.hazelcast.com/docs/multimap#value-collection-type)
	ValueCollectionType ValueCollectionType `json:"valueCollectionType"`
}

//Type of MultiMap Data Structure Response
type MultiMapConfigResponse struct {
	//Unique Identifier of the MultiMap config.
	Id string `json:"id"`
	//Name of the MultiMap config.
	Name string `json:"name"`
	//The number of asynchronous backups. For more information, see [Asynchronous Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/multimap#asynchronous-backups-count)
	AsyncBackupCount int `json:"asyncBackupCount"`
	//The number of synchronous backups. For more information, see [Backup Count]( https://docs-dedicated.cloud.hazelcast.com/docs/multimap#backup-count)
	BackupCount int `json:"backupCount"`
	//Shows that if this config is ready to use.
	IsReady bool `json:"isReady"`
	//Type of the value collection. For more information, see [Value Collection Type]( https://docs-dedicated.cloud.hazelcast.com/docs/multimap#value-collection-type)
	ValueCollectionType string `json:"valueCollectionType"`
}

//Type of Ring Buffer Data Structure Request
type RingBufferConfigRequest struct {
	//Name of the RingBuffer config.
	Name string `json:"name"`
	//The number of asynchronous backups. For more information, see [Asynchronous Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/ringbuffer#asynchronous-backup-count)
	AsyncBackupCount int `json:"asyncBackupCount"`
	//The number of synchronous backups. For more information, see [Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/ringbuffer#backup-count)
	BackupCount int `json:"backupCount"`
	//Capacity of the RingBuffer. For more information, see [Capacity](https://docs-dedicated.cloud.hazelcast.com/docs/ringbuffer#capacity)
	Capacity int `json:"capacity"`
	//In-memory format that controls the format of the RingBuffer’s stored items. For more information, see [In Memory Format]( https://docs-dedicated.cloud.hazelcast.com/docs/ringbuffer#in-memory-format)
	InMemoryFormat InMemoryFormat `json:"inMemoryFormat"`
	//The value that controls how long the items remain in the RingBuffer before they are expired. For more information, see [TTL]( https://docs-dedicated.cloud.hazelcast.com/docs/ringbuffer#time-to-live)
	TtlSeconds int `json:"ttlSeconds"`
}

//Type of Ring Buffer Data Structure Response
type RingBufferConfigResponse struct {
	//Unique Identifier of the RingBuffer config.
	Id string `json:"id"`
	//Name of the RingBuffer config.
	Name string `json:"name"`
	//The number of asynchronous backups. For more information, see [Asynchronous Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/ringbuffer#asynchronous-backup-count)
	AsyncBackupCount int `json:"asyncBackupCount"`
	//The number of synchronous backups. For more information, see [Backup Count](https://docs-dedicated.cloud.hazelcast.com/docs/ringbuffer#backup-count)
	BackupCount int `json:"backupCount"`
	//Shows that if this config is ready to use.
	IsReady bool `json:"isReady"`
	//Capacity of the RingBuffer. For more information, see [Capacity](https://docs-dedicated.cloud.hazelcast.com/docs/ringbuffer#capacity)
	Capacity int `json:"capacity"`
	//In-memory format that controls the format of the RingBuffer’s stored items. For more information, see [In Memory Format]( https://docs-dedicated.cloud.hazelcast.com/docs/ringbuffer#in-memory-format)
	InMemoryFormat string `json:"inMemoryFormat"`
	//The value that controls how long the items remain in the RingBuffer before they are expired. For more information, see [TTL]( https://docs-dedicated.cloud.hazelcast.com/docs/ringbuffer#time-to-live)
	TtlSeconds int `json:"ttlSeconds"`
}

//Type of Reliable Topic Data Structure Request
type ReliableTopicConfigRequest struct {
	//Name of the ReliableTopic config.
	Name string `json:"name"`
	//The minimum number of messages that Reliable Topic tries to read in batches. For more information, see [Batch Size](https://docs-dedicated.cloud.hazelcast.com/docs/reliable-topic#reliable-topic-batch-size)
	ReadBatchSize int `json:"readBatchSize"`
	//A policy to deal with an overloaded topic; so topic where there is no place to store new messages. For more information, see [Overload Policy]( https://docs-dedicated.cloud.hazelcast.com/docs/reliable-topic#reliable-topic-overload-policy )
	TopicOverloadPolicy TopicOverloadPolicy `json:"topicOverloadPolicy"`
}

//Type of Reliable Topic Data Structure Response
type ReliableTopicConfigResponse struct {
	//Unique Identifier of the ReliableTopic config.
	Id string `json:"id"`
	//Name of the ReliableTopic config.
	Name string `json:"name"`
	//Shows that if this config is ready to use.
	IsReady bool `json:"isReady"`
	//The minimum number of messages that Reliable Topic tries to read in batches. For more information, see [Batch Size](https://docs-dedicated.cloud.hazelcast.com/docs/reliable-topic#reliable-topic-batch-size)
	ReadBatchSize int `json:"readBatchSize"`
	//A policy to deal with an overloaded topic; so topic where there is no place to store new messages. For more information, see [Overload Policy]( https://docs-dedicated.cloud.hazelcast.com/docs/reliable-topic#reliable-topic-overload-policy )
	TopicOverloadPolicy string `json:"topicOverloadPolicy"`
}

//Type of Cluster Resume Request
type ClusterResumeRequest struct {
	ClusterId string `json:"ClusterId"`
}

//Type of Cluster Delete Request
type ClusterDeleteRequest struct {
	ClusterId string `json:"ClusterId"`
}

//Type of Cluster Stop Request
type ClusterStopRequest struct {
	ClusterId string `json:"ClusterId"`
}

//Type of Cluster List Request
type ClusterListRequest struct {
	ProductType ProductTypeName `json:"productType"`
}
