package hazelcastcloud

import (
	"context"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

//This InstanceTypeService is used to make operations related with instance types
type InstanceTypeService interface {
	List(ctx context.Context, request *models.InstanceTypeRequest) (*[]models.InstanceType, *Response, error)
}

type instanceTypeServiceOp struct {
	client *Client
}

func NewInstanceTypeService(client *Client) InstanceTypeService {
	return &instanceTypeServiceOp{client}
}

//This function returns a list of available instance types
func (c instanceTypeServiceOp) List(ctx context.Context, request *models.InstanceTypeRequest) (*[]models.InstanceType, *Response, error) {
	//noinspection GoPreferNilSlice
	var instanceTypeListResponse = []models.InstanceType{}
	graphqlRequest := models.GraphqlRequest{
		Name:      "instanceTypes",
		Operation: models.Query,
		Input:     nil,
		Args:      *request,
		Response:  instanceTypeListResponse,
	}
	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &instanceTypeListResponse)
	if err != nil {
		return nil, resp, err
	}

	return &instanceTypeListResponse, resp, err
}
