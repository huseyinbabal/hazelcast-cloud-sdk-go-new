package models

//Type of region request
type RegionRequest struct {
	CloudProvider string `json:"cloudProvider"`
}

//Type of region
type Region struct {
	Name                   string `json:"name"`
	IsEnabledForStarter    bool   `json:"isEnabledForStarter"`
	IsEnabledForEnterprise bool   `json:"isEnabledForEnterprise"`
}
