package gosqoop2

import "encoding/json"

// Connector 对应 Sqoop 中的 Connector
type Connector struct {
	ID                 int                      `json:id`
	LinkConfig         []map[string]interface{} `json:"LinkConfig"`
	JobConfig          map[string]interface{}   `json:"job-config"`
	Name               string                   `json:"name"`
	Class              string                   `json:"class"`
	AllConfigResources map[string]interface{}   `json:"all-config-resources"`
	Version            string                   `json:"version"`
}

type ConnectorList struct {
	Connectors []Connector `json:"connectors"`
}

// ParseConnectorsList 根据 JSON 解析为对应的 Connector 数据结构
func ParseConnectorsList(jsonStr string) (*ConnectorList, error) {
	var clist ConnectorList
	err := json.Unmarshal([]byte(jsonStr), &clist)
	return &clist, err
}
