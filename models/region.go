package models

//Type of region input
type RegionInput struct {
	CloudProvider string `json:"cloudProvider"`
}

//Type of region
type Region struct {
	Name                   string `json:"name"`
	IsEnabledForStarter    bool   `json:"isEnabledForStarter"`
	IsEnabledForEnterprise bool   `json:"isEnabledForEnterprise"`
}