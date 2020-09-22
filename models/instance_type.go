package models

//Type of Instance types input
type InstanceTypeInput struct {
	CloudProvider string `json:"cloudProvider"`
}

//Type of Instance type
type InstanceType struct {
	Name        string `json:"name"`
	TotalMemory int    `json:"totalMemory"`
}