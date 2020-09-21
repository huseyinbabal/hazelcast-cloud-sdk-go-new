package hazelcastcloud

import (
	"context"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

//This AuthService is used to make authorization operations
type AuthService interface {
	Login(ctx context.Context, request *models.LoginInput) (*models.LoginResponse, *Response, error)
}

type authServiceOp struct {
	client *Client
}

func NewAuthService(client *Client) AuthService {
	return &authServiceOp{client: client}
}

//This function logins you with apiKey and apiSecret in request and returns token in the response
func (s authServiceOp) Login(ctx context.Context, request *models.LoginInput) (*models.LoginResponse, *Response, error) {
	//noinspection GoPreferNilSlice
	var loginResponse = models.LoginResponse{}
	var graphqlRequest = models.GraphqlRequest{
		Name:      "login",
		Operation: models.Mutation,
		Args:      *request,
		Response:  loginResponse,
	}
	req, err := s.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(ctx, req, &loginResponse)
	if err != nil {
		return nil, resp, err
	}

	return &loginResponse, resp, err
}
