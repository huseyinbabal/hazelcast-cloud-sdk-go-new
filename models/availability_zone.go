package models

//Type of availability zone input
type AvailabilityZoneInput struct {
	CloudProvider string `json:"cloudProvider"`
	Region        string `json:"region"`
	InstanceType  string `json:"instanceType"`
	InstanceCount int    `json:"instanceCount"`
}

//Type of availability zone
type AvailabilityZone struct {
	Name string `json:"name"`
}