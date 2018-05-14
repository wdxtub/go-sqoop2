package gosqoop2

import (
	"testing"
	"fmt"
)

func TestParseVersion(t *testing.T) {
	jsonStr := JSONVersion
	version, err := ParseVersion(jsonStr)
	if err != nil {
		t.Error("版本解析失败。", err)
		return
	}
	fmt.Println("Version Info:")
	fmt.Println(version.APIVersions)
	fmt.Println(version.BuildVersion)
}

func TestParseConnectorsList(t *testing.T) {
	jsonStr := JSONAllConnector
	clist, err := ParseConnectorsList(jsonStr)
	if err != nil {
		t.Error("Connectors 解析失败。", err)
		return
	}
	fmt.Println("Connector List:")
	for _, v := range clist.Connectors {
		fmt.Println(v.ID, v.Name, v.Class, v.Version)
	}
}

func TestParseLinkList(t *testing.T) {
	jsonStr := JSONAllLink
	llist, err := ParseLinkList(jsonStr)
	if err != nil {
		t.Error("Links 解析失败。", err)
		return
	}
	fmt.Println("Link List:")
	for _, v := range llist.Links {
		fmt.Println(v.ID, v.Name, v.ConnectorName, v.CreationUser, v.UpdateUser)
	}
}

func TestParseJobList(t *testing.T) {
	jsonStr := JSONAllJob
	jlist, err := ParseJobList(jsonStr)
	if err != nil {
		t.Error("Jobs 解析失败。", err)
		return
	}
	fmt.Println("Job List:")
	for _, v := range jlist.Jobs {
		fmt.Println(v.ID, v.Name, v.CreationUser, v.UpdateUser)
	}
}

func TestCreateLink(t *testing.T) {
	jsonStr, err := CreateMySQLLink("yiplatform",
		"test_link",
		"generic-jdbc-connector",
		"com.mysql.jdbc.Driver",
		"jdbc:mysql://127.0.0.1:3306/yiplatform",
		"root",
		"zyqqv587^^")
	if err != nil {
		t.Error("Jobs 解析失败。", err)
		return
	}
	fmt.Println("POST DATA:")
	fmt.Println(jsonStr)
}
