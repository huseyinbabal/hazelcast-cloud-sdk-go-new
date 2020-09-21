package models

//Type of graphql request
type GraphqlRequest struct {
	Operation GraphqlOperation
	Name      string
	Input     interface{}
	Args      interface{}
	Response  interface{}
}
//Type of graphql operation
type GraphqlOperation string

const (
	Mutation GraphqlOperation = "mutation"
	Query    GraphqlOperation = "query"
)

//Type of graphql response
type GraphQLResponse interface {
	Data() interface{}
	Errors() []*GraphQLError
}

//Type of graphql error
type GraphQLError struct {
	Message   string `json:"message"`
	Locations []struct {
		Line   int `json:"line"`
		Column int `json:"column"`
	} `json:"locations"`
}
