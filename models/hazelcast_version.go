package models

//The Hazelcast versions for Enterprise Hazelcast Version.
type EnterpriseHazelcastVersion struct {
	Version             string   `json:"version"`
	UpgradeableVersions []string `json:"upgradeableVersions"`
}