package util

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

// GetHTTPResponse 处理HTTP结果，返回序列化的json
func GetHTTPResponse(resp *http.Response, url string, err error, result any) error {
	body, err := GetHTTPResponseOrg(resp, url, err)

	if err == nil {
		err = json.Unmarshal(body, &result)

		if err != nil {
			log.Println("failed requesting ", url, " with error: ", err)
		}
	}

	return err

}

// GetHTTPResponseOrg 处理HTTP结果，返回byte
func GetHTTPResponseOrg(resp *http.Response, url string, err error) ([]byte, error) {
	if err != nil {
		log.Println("failed requesting ", url, " with error: ", err)
		Ipv4Cache.ForceCompare = true
		Ipv6Cache.ForceCompare = true
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		Ipv4Cache.ForceCompare = true
		Ipv6Cache.ForceCompare = true
		log.Println("failed requesting ", url, " with error: ", err)
	}

	// 300及以上状态码都算异常
	if resp.StatusCode >= 300 {
		errMsg := fmt.Sprintf("failed requesting %s with status %d: %s", url, resp.StatusCode, string(body))
		log.Println(errMsg)
		err = fmt.Errorf(errMsg)
	}

	return body, err
}

// CreateHTTPClient CreateHTTPClient
func CreateHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			IdleConnTimeout:     10 * time.Second,
			TLSHandshakeTimeout: 10 * time.Second,
		},
	}
}
