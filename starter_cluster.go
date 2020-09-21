package hazelcastcloud

import (
	"context"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

//This StarterClusterService is used to interact with Starter Clusters.
type StarterClusterService interface {
	Get(ctx context.Context, request *models.GetStarterClusterInput) (*models.ClusterResponse, *Response, error)
	Create(ctx context.Context, request *models.CreateStarterClusterInput) (*models.ClusterResponse, *Response, error)
	List(ctx context.Context) (*[]models.ClusterResponse, *Response, error)
	Resume(ctx context.Context, request *models.ClusterResumeRequest) (*models.ClusterIdResponse, *Response, error)
	Stop(ctx context.Context, request *models.ClusterStopRequest) (*models.ClusterIdResponse, *Response, error)
	Delete(ctx context.Context, request *models.ClusterDeleteRequest) (*models.ClusterIdResponse, *Response, error)
}

type starterClusterServiceOp struct {
	client *Client
}

func NewStarterClusterService(client *Client) StarterClusterService {
	return &starterClusterServiceOp{client: client}
}

//This function gets detailed configuration of started cluster according to cluster id that provided on the request
func (c starterClusterServiceOp) Get(ctx context.Context, request *models.GetStarterClusterInput) (*models.ClusterResponse, *Response, error) {
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

//This function creates started cluster according to configuration provided in the request
func (c starterClusterServiceOp) Create(ctx context.Context, request *models.CreateStarterClusterInput) (*models.ClusterResponse, *Response, error) {
	var clusterResponse models.ClusterResponse
	var graphqlRequest = models.GraphqlRequest{
		Name:      "createStarterCluster",
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

//This function list all non-deleted Starter Cluster
func (c starterClusterServiceOp) List(ctx context.Context) (*[]models.ClusterResponse, *Response, error) {
	//noinspection ALL
	var clusterListResponse = []models.ClusterResponse{}
	graphqlRequest := models.GraphqlRequest{
		Name:      "clusters",
		Operation: models.Query,
		Input:     nil,
		Args: models.ClusterListRequest{
			ProductType: models.Starter,
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

//This function resume a stopped Starter Cluster
func (c starterClusterServiceOp) Resume(ctx context.Context, request *models.ClusterResumeRequest) (*models.ClusterIdResponse, *Response, error) {
	var clusterIdResponse = models.ClusterIdResponse{}
	graphqlRequest := models.GraphqlRequest{
		Name:      "resumeCluster",
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
		return nil, nil, err
	}

	return &clusterIdResponse, resp, err
}

//This function stops a running Starter Cluster
func (c starterClusterServiceOp) Stop(ctx context.Context, request *models.ClusterStopRequest) (*models.ClusterIdResponse, *Response, error) {
	var clusterIdResponse = models.ClusterIdResponse{}
	graphqlRequest := models.GraphqlRequest{
		Name:      "stopCluster",
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

//This function deletes a starter Starter Cluster
func (c starterClusterServiceOp) Delete(ctx context.Context, request *models.ClusterDeleteRequest) (*models.ClusterIdResponse, *Response, error) {
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
