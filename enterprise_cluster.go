package hazelcastcloud

import (
	"context"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

//This EnterpriseClusterService is used to make operations related with enterprise clusters
type EnterpriseClusterService interface {
	Get(ctx context.Context, input *models.GetEnterpriseClusterInput) (*models.Cluster, *Response, error)
	Create(ctx context.Context, input *models.CreateEnterpriseClusterInput) (*models.Cluster, *Response, error)
	List(ctx context.Context) (*[]models.Cluster, *Response, error)
	Delete(ctx context.Context, input *models.ClusterDeleteInput) (*models.ClusterId, *Response, error)
}

type enterpriseClusterServiceOp struct {
	client *Client
}

func NewEnterpriseClusterService(client *Client) EnterpriseClusterService {
	return &enterpriseClusterServiceOp{client: client}
}

//This function returns detailed configuration of the cluster
func (c enterpriseClusterServiceOp) Get(ctx context.Context, input *models.GetEnterpriseClusterInput) (*models.Cluster, *Response, error) {
	var cluster models.Cluster
	var graphqlRequest = models.GraphqlRequest{
		Name:      "cluster",
		Operation: models.Query,
		Input:     nil,
		Args:      *input,
		Response:  cluster,
	}
	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &cluster)
	if err != nil {
		return nil, resp, err
	}

	return &cluster, resp, err
}

//This function creates an Enterprise Cluster with a configuration provided in the input
func (c enterpriseClusterServiceOp) Create(ctx context.Context, input *models.CreateEnterpriseClusterInput) (*models.Cluster, *Response, error) {
	var cluster models.Cluster
	var graphqlRequest = models.GraphqlRequest{
		Name:      "createEnterpriseCluster",
		Operation: models.Mutation,
		Input:     input,
		Args:      nil,
		Response:  cluster,
	}
	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &cluster)
	if err != nil {
		return nil, resp, err
	}

	return &cluster, resp, err
}

//This function lists all non-deleted Enterprise Clusters
func (c enterpriseClusterServiceOp) List(ctx context.Context) (*[]models.Cluster, *Response, error) {
	var clusterList []models.Cluster
	graphqlRequest := models.GraphqlRequest{
		Name:      "clusters",
		Operation: models.Query,
		Input:     nil,
		Args: models.ClusterListInput{
			ProductType: models.Enterprise,
		},
		Response: clusterList,
	}
	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &clusterList)
	if err != nil {
		return nil, resp, err
	}

	return &clusterList, resp, err
}

//This function deletes an Enterprise Cluster
func (c enterpriseClusterServiceOp) Delete(ctx context.Context, input *models.ClusterDeleteInput) (*models.ClusterId, *Response, error) {
	var clusterId models.ClusterId
	graphqlRequest := models.GraphqlRequest{
		Name:      "deleteCluster",
		Operation: models.Mutation,
		Input:     nil,
		Args:      *input,
		Response:  clusterId,
	}
	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &clusterId)
	if err != nil {
		return nil, resp, err
	}

	return &clusterId, resp, err
}
