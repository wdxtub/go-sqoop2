package gosqoop2

import (
	"net/http"
	"net"
	"time"
	"github.com/astaxie/beego"
	"io/ioutil"
	"encoding/json"
	"strings"
)

// SendGet 发送 Get 请求
func SendGet(endpoint string) ([]byte, error) {
	c := http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(2 * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*1)
				if err != nil {
					beego.Debug(err)
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}

	// 给客户端发送请求
	resp, err := c.Get(endpoint)
	if err != nil {
		beego.Debug(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err)
		return nil, err
	}
	return body, nil
}

// SendPost 发送 Post 请求
// 传入完整地址和要发送的参数
func SendPost(endpoint string, paras map[string]interface{}) ([]byte, error) {
	paraJSON, err := json.Marshal(paras)
	if err != nil {
		beego.Debug(err)
		return nil, err
	}

	// 给客户端发送请求
	resp, err := http.Post(endpoint, "application/json",
		strings.NewReader(string(paraJSON)))
	if err != nil {
		beego.Debug(err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Debug(err)
		return nil, err
	}
	return body, nil
}
