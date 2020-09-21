package hazelcastcloud

import (
	"context"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

//This RegionService is used to make operations related with regions
type RegionService interface {
	List(ctx context.Context, input *models.RegionInput) (*[]models.Region, *Response, error)
}

type regionServiceOp struct {
	client *Client
}

func NewRegionService(client *Client) RegionService {
	return &regionServiceOp{client}
}

//This function returns a list of available regions
func (c regionServiceOp) List(ctx context.Context, input *models.RegionInput) (*[]models.Region, *Response, error) {
	var regionList []models.Region
	graphqlRequest := models.GraphqlRequest{
		Name:      "regions",
		Operation: models.Query,
		Input:     nil,
		Args:      *input,
		Response:  regionList,
	}
	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &regionList)
	if err != nil {
		return nil, resp, err
	}

	return &regionList, resp, err
}
