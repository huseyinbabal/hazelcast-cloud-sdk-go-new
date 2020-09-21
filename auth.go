package hazelcastcloud

import (
	"context"
	"github.com/hazelcast/hazelcast-cloud-sdk-go/models"
)

//This AuthService is used to make authorization operations
type AuthService interface {
	Login(ctx context.Context, input *models.LoginInput) (*models.Login, *Response, error)
}

type authServiceOp struct {
	client *Client
}

func NewAuthService(client *Client) AuthService {
	return &authServiceOp{client: client}
}

//This function logins you with apiKey and apiSecret in input and returns token in the response
func (s authServiceOp) Login(ctx context.Context, input *models.LoginInput) (*models.Login, *Response, error) {
	var login models.Login
	var graphqlRequest = models.GraphqlRequest{
		Name:      "login",
		Operation: models.Mutation,
		Args:      *input,
		Response:  login,
	}
	req, err := s.client.NewRequest(&graphqlRequest)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(ctx, req, &login)
	if err != nil {
		return nil, resp, err
	}

	return &login, resp, err
}
