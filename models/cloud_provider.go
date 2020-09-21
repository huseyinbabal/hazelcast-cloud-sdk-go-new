package models

//Type of cloud provider
type CloudProvider struct {
	Name                   string `json:"name"`
	IsEnabledForStarter    bool   `json:"isEnabledForStarter"`
	IsEnabledForEnterprise bool   `json:"isEnabledForEnterprise"`
}
