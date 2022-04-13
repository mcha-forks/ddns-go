package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

// Config 配置
type Config struct {
	Ipv4 struct {
		Enable       bool
		NetInterface string
		Domains      []string
	}
	Ipv6 struct {
		Enable       bool
		NetInterface string
		Domains      []string
	}
	DNS DNSConfig
	Webhook
	TTL string
}

type DNSConfig struct {
	Name   string
	ID     string
	Secret string
}

func FromFile(path string) (conf Config, err error) {
	_, err = os.Stat(path)
	if err != nil {
		return
	}

	byt, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(byt, conf)
	return
}

func (conf *Config) GetIpv4Addr() (result string) {
	ipv4, _, err := GetNetInterface()
	if err != nil {
		return
	}

	for _, netInterface := range ipv4 {
		if netInterface.Name == conf.Ipv4.NetInterface && len(netInterface.Address) > 0 {
			return netInterface.Address[0]
		}
	}

	log.Println("[ipv4] failed reading address from interface: ", conf.Ipv4.NetInterface)
	return
}

// GetIpv6Addr 获得IPv6地址
func (conf *Config) GetIpv6Addr() (result string) {
	// 从网卡获取IP
	_, ipv6, err := GetNetInterface()
	if err != nil {
		return
	}

	for _, netInterface := range ipv6 {
		if netInterface.Name == conf.Ipv6.NetInterface && len(netInterface.Address) > 0 {
			return netInterface.Address[0]
		}
	}

	log.Println("[ipv6] failed reading address from interface: ", conf.Ipv6.NetInterface)
	return
}
