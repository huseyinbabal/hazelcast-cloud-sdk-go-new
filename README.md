## :construction: This SDK is still in active development and will be published soon!
# Hazelcast Cloud SDK - Go

[![GoDoc](https://godoc.org/github.com/hazelcast/hazelcast-cloud-go-sdk?status.svg)](https://pkg.go.dev/github.com/huseyinbabal/hazelcast-cloud-sdk-go-new)

Hazelcast Cloud is a client library to consume Public API easily.

You can view Hazelcast Cloud API references from here: [https://docs.cloud.hazelcast.com/v1.0/docs/api-reference](https://docs.cloud.hazelcast.com/v1.0/docs/api-reference)

## Install
```sh
go get github.com/huseyinbabal/hazelcast-cloud-sdk-go-new@vX.Y.Z
```

where X.Y.Z is the [version](https://github.com/huseyinbabal/hazelcast-cloud-sdk-go-new/releases) you need.

or
```sh
go get github.com/hazelcast/hazelcast-cloud-go
```
for non Go modules usage or latest version.

## Usage

```go
import "github.com/hazelcast/hazelcast-cloud-go"
```

Create a new HazelcastCloud client, then use the exposed services to
access different parts of the Hazelcast Cloud Public API.

### Authentication

Currently, using API Keys and API Secrets is the only method of
authenticating with the API. You can manage your tokens
in [Hazelcast Cloud Developer Page](https://cloud.hazelcast.com/settings/developer).

Then, you can use your API Key and API Secret to create a new client as shown below:

```go
package main

import (
    "github.com/hazelcast/hazelcast-cloud-go-sdk"
)

func main() {
    client := hazelcastcloud.NewFromCredentials("API-KEY", "API-SECRET")
}
```

Also, you can use [`hazelcastcloud.New()`](https://github.com/huseyinbabal/hazelcast-cloud-sdk-go-new/blob/master/hazelcast_cloud.go#L113) to provider your API Key and API Secret from environment variables as `HZ_CLOUD_API_KEY` and `HZ_CLOUD_API_SECRET` 

## :rocket: Examples
- Create a new Starter Cluster:

```go
clusterName := "my-awesome-cluster"
createInput := models.CreateStarterClusterInput{
  Name:             clusterName,
  CloudProvider:    "aws",
  Region:           "us-west-2",
  ClusterType:      "FREE",
  HazelcastVersion: "4.0",
  TotalMemory:      0.2,
}

newCluster, _, createErr := client.StarterCluster.Create(context.Background(), &createInput)
if createErr != nil {
  fmt.Printf("An error occurred: %s\n\n", createErr)
  return
}
```

## 🏷️ Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/huseyinbabal/hazelcast-cloud-sdk-go-new/tags).

## ⭐️ Documentation

For details on all the functionality in this library, see the [GoDoc](http://godoc.org/github.com/huseyinbabal/hazelcast-cloud-sdk-go-new) documentation. Also, you can refer [API References](https://docs.cloud.hazelcast.com/docs/api-reference) for Graphql queries and mutations to understand payloads.


## 🤝 Contributing

Contributions, issues and feature requests are welcome!<br />Feel free to check [issues page](https://github.com/huseyinbabal/hazelcast-cloud-sdk-go-new/issues).


## 📝 License

Copyright © 2020 [Hazelcast](https://github.com/hazelcast).<br />
This project is [Apache License 2.0](https://github.com/hazelcast/hazelcast-cloud-cli) licensed.<br /><br />
<img alt="logo" width="300" src="https://cloud.hazelcast.com/static/images/hz-cloud-logo.svg" />
