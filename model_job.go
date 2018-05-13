package gosqoop2

import "encoding/json"

// Job 对应 Sqoop 里的 Job
type Job struct {
	ID                 int         `json:"id"`
	Name               string      `json:"name"`
	CreationUser       string      `json:"creation-user"`
	UpdateUser         string      `json:"update-user"`
	CreationDate       int64       `json:"creation-date"`
	UpdateDate         int64       `json:"update-date"`
	ToConfigValues     ConfigValue `json:"to-config-values"`
	FromConfigValues   ConfigValue `json:"from-config-values"`
	DriverConfigValues ConfigValue `json:"driver-config-values"`
	ToConnectorName    string      `json:"to-connector-name"`
	FromConnectorName  string      `json:"from-connector-name"`
	ToLinkName         string      `json:"to-link-name"`
	FromLinkName       string      `json:"from-link-name"`
	Enabled            bool        `json:"enabled"`
}

type JobList struct {
	Jobs []Job `json:"jobs"`
}

// ParseJobList 解析 List 列表
func ParseJobList(jsonStr string) (*JobList, error) {
	var jlist JobList
	err := json.Unmarshal([]byte(jsonStr), &jlist)
	return &jlist, err
}