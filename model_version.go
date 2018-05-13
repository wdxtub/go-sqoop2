package gosqoop2

import "encoding/json"

// Version Sqoop 的版本信息
type Version struct {
	SourceURL      string   `json:"source-url"`
	SourceRevision string   `json:"source-revision"`
	BuildVersion   string   `json:"build-version"`
	APIVersions    []string `json:"api-versions"`
	User           string   `json:"user"`
	BuildDate      string   `json:"build-date"`
}

// ParseVersion 根据 JSON 解析为对应的 Version 数据结构
func ParseVersion(jsonStr string) (*Version, error) {
	var version Version
	err := json.Unmarshal([]byte(jsonStr), &version)
	return &version, err
}
