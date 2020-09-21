package hazelcastcloud

import (
	"context"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

//This InstanceTypeService is used to make operations related with instance types
type InstanceTypeService interface {
	List(ctx context.Context, input *models.InstanceTypeInput) (*[]models.InstanceType, *Response, error)
}

type instanceTypeServiceOp struct {
	client *Client
}

func NewInstanceTypeService(client *Client) InstanceTypeService {
	return &instanceTypeServiceOp{client}
}

//This function returns a list of available instance types
func (c instanceTypeServiceOp) List(ctx context.Context, input *models.InstanceTypeInput) (*[]models.InstanceType, *Response, error) {
	var instanceTypeList []models.InstanceType
	graphqlRequest := models.GraphqlRequest{
		Name:      "instanceTypes",
		Operation: models.Query,
		Input:     nil,
		Args:      *input,
		Response:  instanceTypeList,
	}
	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &instanceTypeList)
	if err != nil {
		return nil, resp, err
	}

	return &instanceTypeList, resp, err
}
