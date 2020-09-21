package hazelcastcloud

import (
	"context"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

//This EnterpriseClusterService is used to make operations related with enterprise clusters
type EnterpriseClusterService interface {
	Get(ctx context.Context, request *models.GetEnterpriseClusterInput) (*models.ClusterResponse, *Response, error)
	Create(ctx context.Context, request *models.CreateEnterpriseClusterInput) (*models.ClusterResponse, *Response, error)
	List(ctx context.Context) (*[]models.ClusterResponse, *Response, error)
	Delete(ctx context.Context, request *models.ClusterDeleteRequest) (*models.ClusterIdResponse, *Response, error)
}

type enterpriseClusterServiceOp struct {
	client *Client
}

func NewEnterpriseClusterService(client *Client) EnterpriseClusterService {
	return &enterpriseClusterServiceOp{client: client}
}

//This function returns detailed configuration of the cluster
func (c enterpriseClusterServiceOp) Get(ctx context.Context, request *models.GetEnterpriseClusterInput) (*models.ClusterResponse, *Response, error) {
	var clusterResponse models.ClusterResponse
	var graphqlRequest = models.GraphqlRequest{
		Name:      "cluster",
		Operation: models.Query,
		Input:     nil,
		Args:      *request,
		Response:  clusterResponse,
	}
	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &clusterResponse)
	if err != nil {
		return nil, resp, err
	}

	return &clusterResponse, resp, err
}

//This function creates an Enterprise Cluster with a configuration provided in the request
func (c enterpriseClusterServiceOp) Create(ctx context.Context, request *models.CreateEnterpriseClusterInput) (*models.ClusterResponse, *Response, error) {
	//noinspection GoPreferNilSlice
	var clusterResponse = models.ClusterResponse{}
	var graphqlRequest = models.GraphqlRequest{
		Name:      "createEnterpriseCluster",
		Operation: models.Mutation,
		Input:     request,
		Args:      nil,
		Response:  clusterResponse,
	}
	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &clusterResponse)
	if err != nil {
		return nil, resp, err
	}

	return &clusterResponse, resp, err
}

//This function lists all non-deleted Enterprise Clusters
func (c enterpriseClusterServiceOp) List(ctx context.Context) (*[]models.ClusterResponse, *Response, error) {
	//noinspection GoPreferNilSlice
	var clusterListResponse = []models.ClusterResponse{}
	graphqlRequest := models.GraphqlRequest{
		Name:      "clusters",
		Operation: models.Query,
		Input:     nil,
		Args: models.ClusterListRequest{
			ProductType: models.Enterprise,
		},
		Response: clusterListResponse,
	}
	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &clusterListResponse)
	if err != nil {
		return nil, resp, err
	}

	return &clusterListResponse, resp, err
}

//This function deletes an Enterprise Cluster
func (c enterpriseClusterServiceOp) Delete(ctx context.Context, request *models.ClusterDeleteRequest) (*models.ClusterIdResponse, *Response, error) {
	//noinspection GoPreferNilSlice
	var clusterIdResponse = models.ClusterIdResponse{}
	graphqlRequest := models.GraphqlRequest{
		Name:      "deleteCluster",
		Operation: models.Mutation,
		Input:     nil,
		Args:      *request,
		Response:  clusterIdResponse,
	}
	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &clusterIdResponse)
	if err != nil {
		return nil, resp, err
	}

	return &clusterIdResponse, resp, err
}
