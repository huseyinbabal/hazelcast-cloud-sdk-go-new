package models

//Type of login request with apiKey and apiSecret
type LoginInput struct {
	ApiKey    string `json:"apiKey"`
	ApiSecret string `json:"apiSecret"`
}

//Type of login response that has token
type LoginResponse struct {
	Token string `json:"token"`
}
