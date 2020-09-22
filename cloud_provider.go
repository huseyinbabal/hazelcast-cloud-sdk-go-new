package hazelcastcloud

import (
	"context"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

//This CloudProviderService is used to make operations related with cloud providers
type CloudProviderService interface {
	List(ctx context.Context) (*[]models.CloudProvider, *Response, error)
}

type cloudProviderServiceOp struct {
	client *Client
}

func NewCloudProviderService(client *Client) CloudProviderService {
	return &cloudProviderServiceOp{client: client}
}

//This function returns a list of cloud providers according to use it while creating clusters for both starter and enterprise.
func (c cloudProviderServiceOp) List(ctx context.Context) (*[]models.CloudProvider, *Response, error) {
	var cloudProviderList []models.CloudProvider
	graphqlRequest := models.GraphqlRequest{
		Name:      "cloudProviders",
		Operation: models.Query,
		Input:     nil,
		Args:      nil,
		Response:  cloudProviderList,
	}
	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &cloudProviderList)
	if err != nil {
		return nil, resp, err
	}

	return &cloudProviderList, resp, err
}
