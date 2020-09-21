package hazelcastcloud

import (
	"context"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

//This AvailabilityZoneService is used to make operations related with availability zones
type AvailabilityZoneService interface {
	List(ctx context.Context, request *models.AvailabilityZoneRequest) (*[]models.AvailabilityZone, *Response, error)
}

type availabilityZoneServiceOp struct {
	client *Client
}

func NewAvailabilityZoneService(client *Client) AvailabilityZoneService {
	return &availabilityZoneServiceOp{client}
}

//This function returns a list of availability zones according to request parameters
func (c availabilityZoneServiceOp) List(ctx context.Context, request *models.AvailabilityZoneRequest) (*[]models.AvailabilityZone, *Response, error) {
	//noinspection GoPreferNilSlice
	var availabilityZoneListResponse = []models.AvailabilityZone{}
	graphqlRequest := models.GraphqlRequest{
		Name:      "availabilityZones",
		Operation: models.Query,
		Input:     nil,
		Args:      *request,
		Response:  availabilityZoneListResponse,
	}
	req, err := c.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := c.client.Do(ctx, req, &availabilityZoneListResponse)
	if err != nil {
		return nil, resp, err
	}

	return &availabilityZoneListResponse, resp, err
}
