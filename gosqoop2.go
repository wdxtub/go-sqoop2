package gosqoop2

import "fmt"

// SqoopClient Sqoop2 的客户端
type SqoopClient struct {
	SqoopHost string
	SqoopPort int
	SqoopApp  string
}

// String 获取 Sqoop 的配置
func (c *SqoopClient) String() string {
	return fmt.Sprintf("Sqoop2<%s:%d/%s>", c.SqoopHost, c.SqoopPort, c.SqoopApp)
}

// SqoopAPIRoot 返回 API 的 root 地址
func (c *SqoopClient) SqoopAPIRoot() string {
	return fmt.Sprintf("%s:%d/%s", c.SqoopHost, c.SqoopPort, c.SqoopApp)
}

// CheckConnection 检查连接状态
// 每次调用之前都需要检查状态
func (c *SqoopClient) CheckConnection() bool {
	url := fmt.Sprintf("%s/version", c.SqoopAPIRoot())
	_, err := SendGet(url)
	if err != nil {
		return true
	}
	return false
}

//func (c *SqoopClient) GetAllConnector() []*Connector {
//
//}

