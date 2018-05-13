package gosqoop2

import (
	"encoding/json"
)

// Link 对应 Sqoop 中的 Link
type Link struct {
	ID               int             `json:"id"`
	Enabled          bool            `json:"enabled"`
	CreationUser     string          `json:"creation-user"`
	UpdateUser       string          `json:"update-user"`
	Name             string          `json:"name"`
	CreationDate     int64           `json:"creation-date"`
	UpdateDate       int64           `json:"update-date"`
	ConnectorName    string          `json:"connector-name"`
	LinkConfigValues ConfigValue `json:"link-config-values"`
}

type LinkList struct {
	Links []Link `json:"links"`
}

//type LinkInput struct {
//	ID         int           `json:"id"`
//	Size       int           `json:"size"`
//	Editable   string        `json:"editable"`
//	Validators []interface{} `json:"validators"`
//	Name       string        `json:"name"`
//	Sensitive  bool          `json:"sensitive"`
//	Overrides  string        `json:"overrides"`
//	Type       string        `json:"type"`
//	Value      string        `json:"value"`
//}

type Config struct {
	ID         int                      `json:"id"`
	Type       string                   `json:"type"`
	Validators []interface{}            `json:"validators"`
	Inputs     []map[string]interface{} `json:"inputs"`
	Name       string                   `json:"name"`
}

type ConfigValue struct {
	Configs    []Config  `json:"configs"`
	Validators []interface{} `json:"validators"`
}

// ParseLinkList 解析 List 列表
func ParseLinkList(jsonStr string) (*LinkList, error) {
	var llist LinkList
	err := json.Unmarshal([]byte(jsonStr), &llist)
	return &llist, err
}
