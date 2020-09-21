package hazelcastcloud

import (
	"context"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

//This HazelcastVersionService is used to make operations related with enterprise clusters
type HazelcastVersionService interface {
	List(ctx context.Context) (*[]models.EnterpriseHazelcastVersion, *Response, error)
}

type hazelcastVersionServiceOp struct {
	client *Client
}

func NewHazelcastVersionService(client *Client) HazelcastVersionService {
	return &hazelcastVersionServiceOp{client}
}

//This function returns a list of available Hazelcast versions
func (c hazelcastVersionServiceOp) List(ctx context.Context) (*[]models.EnterpriseHazelcastVersion, *Response, error) {
	var hazelcastVersionList []models.EnterpriseHazelcastVersion
	graphqlRequest := models.GraphqlRequest{
		Name:      "hazelcastVersions",
		Operation: models.Query,
		Input:     nil,
		Args:      nil,
		Response:  hazelcastVersionList,
	}

	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &hazelcastVersionList)
	if err != nil {
		return nil, resp, err
	}

	return &hazelcastVersionList, resp, err
}
