package models

//Type of login input with apiKey and apiSecret
type LoginInput struct {
	ApiKey    string `json:"apiKey"`
	ApiSecret string `json:"apiSecret"`
}

//Type of login response that has token
type Login struct {
	Token string `json:"token"`
}