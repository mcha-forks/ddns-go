package dns

import (
	"dnsd/config"
	"log"
	"time"
)

type DNS interface {
	Init(conf *config.Config)
	AddUpdateDomainRecords() (domains config.Domains)
}

// RunTimer 定时运行
func RunTimer(delay time.Duration, interval time.Duration, conf config.Config) {
	log.Printf("waiting for %ds", int(delay.Seconds()))
	time.Sleep(delay)
	for {
		RunOnce(conf)
		time.Sleep(interval)
	}
}

func RunOnce(conf config.Config) {
	var dnsSelected DNS
	switch conf.DNS.Name {
	case "cloudflare":
		dnsSelected = &Cloudflare{}
	case "callback":
		dnsSelected = &Callback{}
	default:
		dnsSelected = &Cloudflare{}
	}
	dnsSelected.Init(&conf)

	domains := dnsSelected.AddUpdateDomainRecords()
	config.ExecWebhook(&domains, &conf)
}
