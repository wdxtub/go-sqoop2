package gosqoop2

import (
	"encoding/json"
	"time"
)

// Link 对应 Sqoop 中的 Link
type Link struct {
	ID               int         `json:"id"`
	Enabled          bool        `json:"enabled"`
	CreationUser     string      `json:"creation-user"`
	UpdateUser       string      `json:"update-user"`
	Name             string      `json:"name"`
	CreationDate     int64       `json:"creation-date"`
	UpdateDate       int64       `json:"update-date"`
	ConnectorName    string      `json:"connector-name"`
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
	Configs    []Config      `json:"configs"`
	Validators []interface{} `json:"validators"`
}


// ParseLinkList 解析 List 列表
func ParseLinkList(jsonStr string) (*LinkList, error) {
	var llist LinkList
	err := json.Unmarshal([]byte(jsonStr), &llist)
	return &llist, err
}

// CreateMySQLLink 创建一个 MySQLlink
func CreateMySQLLink(user, linkName, connectorName, driverStr, connectionStr, sqluser, sqlpasswd string) (string, error) {

	var links LinkList

	var link Link
	link.ID = -1
	link.Enabled = true
	link.UpdateDate = time.Now().UnixNano() / 1e6
	link.UpdateUser = user
	link.CreationDate = link.UpdateDate
	link.CreationUser = link.UpdateUser
	link.ConnectorName = connectorName

	link.Name = linkName
	var configValue ConfigValue
	configValue.Validators = []interface{}{}
	var config Config
	config.Validators = []interface{}{}
	config.Name = "linkConfig"
	config.ID = 1
	config.Type = "LINK"
	config.Inputs = CreateMySQLConfigInputs(driverStr, connectionStr,sqluser, sqlpasswd)
	configValue.Configs = append(configValue.Configs, config)

	var dialectConfig Config
	dialectConfig.Validators = []interface{}{}
	dialectConfig.Name = "dialect"
	dialectConfig.ID = 2
	dialectConfig.Type = "LINK"
	dialectConfig.Inputs = CreateMySQLDiaglectInputs()
	configValue.Configs = append(configValue.Configs, dialectConfig)

	link.LinkConfigValues = configValue
	links.Links = append(links.Links, link)
	jsonStr, err := json.Marshal(links)

	return string(jsonStr), err

}

//
func CreateMySQLDiaglectInputs() []map[string]interface{} {
	data := []map[string]interface{}{}
	dialect := map[string]interface{}{}
	dialect["size"] = 5
	dialect["editable"] = "ANY"
	dialect["validators"] = []interface{}{}
	dialect["name"] = "dialect.identifierEnclose"
	dialect["id"] = 7
	dialect["sensitive"] = false
	dialect["overrides"] = ""
	dialect["type"] = "STRING"
	dialect["value"] = "++"

	data = append(data, dialect)
	return data
}


// CreateMySQLConfigInputs 创建 MySQL 的配置
func CreateMySQLConfigInputs(driverStr, connectionStr, user, passwd string ) []map[string]interface{} {
	data := []map[string]interface{}{}
	driver := map[string]interface{}{}
	connectString := map[string]interface{}{}
	userName := map[string]interface{}{}
	password := map[string]interface{}{}
	fetchSize := map[string]interface{}{}
	jdbcProperties := map[string]interface{}{}

	// Driver
	driver["size"] = 128
	driver["editable"] = "ANY"
	driver["validators"] = []interface{}{}
	driver["name"] = "linkConfig.jdbcDriver"
	driver["id"] = 1
	driver["sensitive"] = false
	driver["overrides"] = ""
	driver["type"] = "STRING"
	driver["value"] = driverStr
	// ConnectString
	connectString["size"] = 2000
	connectString["editable"] = "ANY"
	connectString["validators"] = []interface{}{}
	connectString["name"] = "linkConfig.connectionString"
	connectString["id"] = 2
	connectString["sensitive"] = false
	connectString["overrides"] = ""
	connectString["type"] = "STRING"
	connectString["value"] = connectionStr
	// Username
	userName["size"] = 40
	userName["editable"] = "ANY"
	userName["validators"] = []interface{}{}
	userName["name"] = "linkConfig.username"
	userName["id"] = 3
	userName["sensitive"] = false
	userName["overrides"] = ""
	userName["type"] = "STRING"
	userName["value"] = user
	// Password
	password["size"] = 40
	password["editable"] = "ANY"
	password["validators"] = []interface{}{}
	password["name"] = "linkConfig.password"
	password["id"] = 4
	password["sensitive"] = false
	password["overrides"] = ""
	password["type"] = "STRING"
	password["value"] = passwd
	// FetchSize
	fetchSize["editable"] = "ANY"
	fetchSize["validators"] = []interface{}{}
	fetchSize["name"] = "linkConfig.fetchSize"
	fetchSize["id"] = 5
	fetchSize["sensitive"] = false
	fetchSize["overrides"] = ""
	fetchSize["type"] = "INTEGER"
	// Jdbc
	jdbcProperties["editable"] = "ANY"
	jdbcProperties["validators"] = []interface{}{}
	jdbcProperties["name"] = "linkConfig.jdbcProperties"
	jdbcProperties["id"] = 6
	jdbcProperties["sensitive"] = false
	jdbcProperties["overrides"] = ""
	jdbcProperties["type"] = "MAP"
	jdbcProperties["value"] = map[string]interface{}{"protocol": "tcp"}
	jdbcProperties["sensitive-pattern"] = ""


	data = append(data, driver)
	data = append(data, connectString)
	data = append(data, userName)
	data = append(data, password)
	//data = append(data, fetchSize)
	data = append(data, jdbcProperties)

	return data
}
